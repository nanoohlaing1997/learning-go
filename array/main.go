package main

import "fmt"

func main() {
	var nameArray = [3][2]string{
		{"nan", "oo"},
		{"tdk", "kn"},
		{"kph", "dg"},
	}
	fmt.Println("Name lists are ", nameArray[0][0])

	var testArray [4]string
	testArray[0] = "testing"
	fmt.Println("Test array is ", testArray)

	var testing []string
	testing[0] = "value 1"
	testing[1] = "value 2"
	fmt.Println("Testing is ", testArray)
}
