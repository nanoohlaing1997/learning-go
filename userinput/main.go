package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comman ok // error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank for reading, ", input)
	fmt.Printf("type of this rating is %T \n", input)
}
