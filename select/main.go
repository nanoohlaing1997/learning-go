package main

import (
	"fmt"
	"time"
)

func main() {
	number := make(chan int)
	message := make(chan string)

	go channelNumber(number)
	go channelMessage(message)

	select {
	case secondChannel := <-message:
		fmt.Println("Second channel :", secondChannel)
	case firstChannel := <-number:
		fmt.Println("First channel :", firstChannel)
	default:
		fmt.Println("Channels are not ready")
	}
}

func channelNumber(number chan int) {
	time.Sleep(time.Second * 1)
	number <- 15
}

func channelMessage(message chan string) {
	time.Sleep(time.Second * 2)
	message <- "Hello Noh"
}
