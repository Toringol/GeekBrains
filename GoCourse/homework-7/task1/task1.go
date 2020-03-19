package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	time.Sleep(time.Second * 10)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
