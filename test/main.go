package main

import "fmt"

const jwtToken string = "testing token"

func main() {
	fmt.Printf("Variable is of Type: %T \n", jwtToken)

	fmt.Println(jwtToken)
}
