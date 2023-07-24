package main

import "fmt"

// func main() {
// 	number := make(chan int)
// 	message := make(chan string)

// 	go ChannelData(number, message)

// 	fmt.Println("Channel Data .", <-number)
// 	fmt.Println("Channel Data .", <-message)
// }

// func ChannelData(number chan int, message chan string) {
// 	number <- 15
// 	message <- "Hello Noh"
// }

// func main() {
// 	ch := make(chan string)

// 	go sentData(ch)

// 	fmt.Println(<-ch)
// }

// func sentData(ch chan string) {
// 	ch <- "Hello NOH"
// 	fmt.Println("No receiver! Send Operation Blocked")
// }

func main() {
	ch := make(chan string)
	ch <- "hello"
	go sentMessage(ch)
}

func sentMessage(ch chan string) {
	<-ch
	fmt.Println(<-ch)
}
