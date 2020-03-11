package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

type Chat struct {
	Clients map[client]bool

	Entering    chan client
	Leaving     chan client
	Messages    chan string
	ServerClose chan struct{}
}

func main() {
	chat := &Chat{
		Clients:     make(map[client]bool),
		Entering:    make(chan client),
		Leaving:     make(chan client),
		Messages:    make(chan string),
		ServerClose: make(chan struct{}),
	}
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	chatExit := make(chan struct{})
	go chat.broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go chat.handleConn(conn)
		defer func() {
			chatExit <- struct{}{}
			select {
			case <-chatExit:
				chat.ServerClose <- struct{}{}
			}
		}()
	}
}

func (chat *Chat) broadcaster() {
	for {
		select {
		case msg := <-chat.Messages:
			for cli := range chat.Clients {
				cli <- msg
			}

		case cli := <-chat.Entering:
			chat.Clients[cli] = true

		case cli := <-chat.Leaving:
			delete(chat.Clients, cli)
			close(cli)
		case <-chat.ServerClose:
			for cli := range chat.Clients {
				delete(chat.Clients, cli)
				close(cli)
			}
		}
	}
}

func (chat *Chat) handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	chat.Messages <- who + " has arrived"
	chat.Entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		chat.Messages <- who + ": " + input.Text()
	}
	chat.Leaving <- ch
	chat.Messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
