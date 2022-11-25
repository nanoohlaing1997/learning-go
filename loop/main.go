package main

import "fmt"

func main() {
	country := []string{
		"Myanmar",
		"Thailand",
		"England",
		"USA",
	}

	for i := 0; i < len(country); i++ {
		fmt.Println("Country is", country[i])
	}

	// like foreach
	for i := range country {
		fmt.Println("Second country is", country[i])
	}

	// like foreach with key value
	for key, value := range country {
		fmt.Printf("Third Country index is %v and thrid country value is %v\n", key, value)
	}

	value := 0
	// like while
	for value < 10 {
		if value == 0 {
			goto loc
		}

		if value == 4 {
			fmt.Println("Loop reach middle")
			value++
			continue
		}

		if value == 9 {
			fmt.Println("Loop will stop.")
			break
		}

		fmt.Println("value is", value)
		value++
	}

loc:
	fmt.Println("Jumpting and break")
}
