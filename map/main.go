package main

import (
	"fmt"
)

func main() {
	phoneOs := make(map[string]string)
	phoneOs["iphone"] = "ios"
	phoneOs["oppo"] = "android"
	fmt.Println("Lists of phone OS : ", phoneOs)

	phoneOs["oppo"] = "Android"
	fmt.Println("Oppo phone OS is ", phoneOs["oppo"])

	phoneOs["sumsaung"] = "Android"
	delete(phoneOs, "oppo")
	fmt.Println("List of deleted phone OS : ", phoneOs)

	languages := map[string]string{
		"JS": "JavaScript",
		"PY": "Python",
	}
	fmt.Println("Lists of programming language are ", languages)

	// Loop in golang
	for key, value := range languages {
		fmt.Printf("For key %v, Value is %v\n", key, value)
	}
}
