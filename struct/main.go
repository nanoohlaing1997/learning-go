package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	// no inheritance in go lang
	noh := User{Name: "nan oo hlaing"}
	fmt.Println("My Name is ", noh)

	noh.Age = 26
	fmt.Println("My age is ", noh.Age)

	tdk := User{}
	tdk.Name = "Thida"
	tdk.Email = "thida@dev.com"
	tdk.Age = 32
	tdk.Status = true
	fmt.Println("Detail of thida are ", tdk)
	fmt.Printf("Detail of thida are %+v\n", tdk)
}
