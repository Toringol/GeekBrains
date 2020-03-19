package main

import (
	"fmt"
	"net/http"
	"time"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("https://golang.org/")
	}()
	go func() {
		responses <- request("https://mail.ru/")
	}()
	go func() {
		responses <- request("https://yandex.ru/")
	}()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {

	start := time.Now()
	resp, err := http.Get(hostname)
	if err != nil {
		return "error get"
	}
	end := time.Now()

	return resp.Request.Host + " " + end.Sub(start).String()
}

func main() {
	fmt.Println(mirroredQuery())
}
