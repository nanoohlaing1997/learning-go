package main

import "fmt"

func zeroptr(iptr *int) {
	*iptr = 132312321
}

func main() {
	num := 50

	var pointer = &num

	fmt.Println("Pointer is ", pointer)
	fmt.Println("Value is ", *pointer)

	*pointer = *pointer + 2
	fmt.Println("Additional Value is ", num)

	i := 1
	zeroptr(&i)
	fmt.Println("Next Test ", i)

}
