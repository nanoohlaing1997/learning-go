package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 00:00:00 Tuesday"))

	createDate := time.Date(2022, time.January, 10, 23, 23, 0, 0, time.Now().Local().Location())
	fmt.Println(createDate)

	fmt.Println(createDate.Format("2006-01-02 Monday 01:02:04"))
}
