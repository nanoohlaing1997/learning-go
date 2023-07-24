package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan string, 1)
	go worker(done)
	fmt.Println(<-done)
}

func worker(done chan string) {
	fmt.Println("working.....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- "worker done"
}
