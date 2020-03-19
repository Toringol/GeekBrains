package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	doneChan := make(chan struct{})

	go exitCommand(doneChan)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, doneChan)
	}
}

func handleConn(c net.Conn, doneChan chan struct{}) {
	defer c.Close()
	for {
		select {
		case <-doneChan:
			return
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func exitCommand(doneChan chan struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			close(doneChan)
			break
		}
	}
}
