package main

import "fmt"

func main() {
	// defer work at the end of the function
	// defer is last in first out
	defer fmt.Println("Defer One")

	defer fmt.Println("Defer Two")

	fmt.Println("Welcome to defer")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
