package main

import (
	"fmt"
)

func main() {
	// normal function
	testOne := test(1)
	fmt.Println(testOne)

	testTwo := test(0)
	fmt.Println(testTwo)

	// variadic function
	resOne := vFunction(1, 2, 3)
	fmt.Println(resOne)

	// variadic function
	resTwo, resMessage, _ := vKeyValFunction("one", "three")
	fmt.Println(resTwo)
	fmt.Println(resMessage)

}

// Normal Function
func test(t int16) string {
	if t == 0 {
		return "true"
	}
	return "false"
}

// Variadic Function
func vFunction(values ...int) int {
	result := 0
	for i := range values {
		result += values[i]
	}
	return result
}

// Variadic Function
func vKeyValFunction(values ...string) (int, string, bool) {
	result := 0
	for _, val := range values {
		if val == "one" {
			result = result + 1
		}
		if val == "two" {
			result = result + 2
		}
		if val == "three" {
			result = result + 3
		}
	}
	return result, "Testing variadic function with key value", true
}
