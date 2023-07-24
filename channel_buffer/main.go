package main

import "fmt"

func main() {
	message := make(chan string, 1)

	message <- "Buffer"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
