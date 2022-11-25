package main

import (
	"fmt"
	"sort"
)

func main() {
	var carList = []string{"SUZUKI"}
	carList = append(carList, "BMW")
	carList = append(carList, "KIA")
	carList = append(carList, "TOTOTA")
	carList = append(carList, "HONDA")
	fmt.Println("Car list is ", carList)

	carList = append(carList[1:])
	fmt.Println("Car List in rang ", carList)

	carList = append(carList[1:3])
	fmt.Println("Car List in rang ", carList)

	carList = append(carList[:1])
	fmt.Println("Car List in rang ", carList)

	salary := make([]int, 3)
	salary[0] = 1000
	salary[1] = 2000
	salary[2] = 850
	fmt.Println("Salary is ", salary)

	salary = append(salary, 1500, 2500)
	fmt.Println("Salary is ", salary)

	sort.Ints(salary)
	fmt.Println("Sorted salary is ", salary)

	var isSorted bool
	isSorted = sort.IntsAreSorted(salary)
	fmt.Println("Is salary is sorted ", isSorted)

	// How to delete values from slice
	var programmingLang = []string{"PHP", "GO", "JS", "JAVA"}
	var index int = 2
	programmingLang = append(programmingLang[:index], programmingLang[index+1:]...)
	fmt.Println("Deleted value from slice ", programmingLang)
}
