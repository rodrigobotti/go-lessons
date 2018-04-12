package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go ping(channel)
	go pong(channel)
	go printer(channel)

	var line string
	fmt.Scanln(&line)
}

func ping(channel chan string) {
	for {
		channel <- "ping"
	}
}

func pong(channel chan string) {
	for {
		channel <- "pong"
	}
}

func printer(channel chan string) {
	for {
		msg := <-channel
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
