package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(7)

	switch randomNumber {
	case 1:
		fmt.Println("Random Number is 1")
	case 2:
		fmt.Println("Randon Number is 2")
	case 3:
		fmt.Println("Randon Number is 3")
	case 4:
		fmt.Println("Randon Number is 4")
	case 5:
		fmt.Println("Randon Number is 5")
	case 6:
		fmt.Println("Randon Number is 6")
	default:
		fmt.Println("What is this!!!!!")
	}
}
