package main

import "fmt"

func main() {
	num := 11
	if num < 10 {
		fmt.Println("num is less than 10")
	} else if num == 10 {
		fmt.Println("num is 10")
	} else {
		fmt.Println("num is greater than 10")
	}

	var salary int64
	if salary = 90000; salary == 100000 {
		fmt.Println("Salary is 1 thein")
	} else if salary < 100000 {
		fmt.Println("Salary is less than 1 thein")
	} else {
		fmt.Println("Salary is greater than 1 thein")
	}
}
