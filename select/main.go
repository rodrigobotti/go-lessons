package main

import (
	"fmt"
	"time"
)

var (
	irc     = make(chan string)
	sms     = make(chan string)
	ping    = publisher("ping")
	pong    = publisher("pong")
	sup     = publisher("'sup ?")
	notmuch = publisher("not much, u?")
)

func main() {
	go ping(irc)
	go pong(irc)

	go sup(sms)
	go notmuch(sms)

	go printer()

	var line string
	fmt.Scanln(&line)
}

func publisher(message string) func(chan string) {
	return func(channel chan string) {
		for {
			channel <- message
		}
	}
}

func printer() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		case msg = <-sms:
			fmt.Println("<static>...", msg)
			time.Sleep(time.Millisecond * 500)
		}
	}
}
