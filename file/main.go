package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to work with file in golang")
	content := "Testing about file and writing something"
	filename := "./mytestfile.txt"

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("String length in file is ", length)
	file.Close()

	readFile("")
}

func readFile(filename string) {
	fileByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("File byte are ", fileByte)
	fmt.Println("File string are ", string(fileByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
