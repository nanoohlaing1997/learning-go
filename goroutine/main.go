package main

import (
	"fmt"
	"time"
)

// The order of execution of goroutine is random
// With Goroutines, concurrency is achieved in Go programming.
// It helps two or more independent functions to run together.
func display(message string) {
	fmt.Println(message)
	// time.Sleep(time.Second * 1)
}

func main() {
	go display("Process 1")
	go display("Process 3")
	go display("Process 2")
	time.Sleep(time.Second * 1)
}
