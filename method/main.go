package main

import "fmt"

type User struct {
	Name   string
	Age    int16
	Status bool
}

func main() {
	fmt.Println("welcome to method testing")

	user := User{"NOH", 25, true}
	fmt.Printf("User details are : %+v\n", user)

	age := user.getAge()
	fmt.Println("Age of user is ", age)

	setStatus := user.setStatus(false)
	fmt.Println("User set status is ", setStatus)

	fmt.Println("User current status is ", user.Status)
}

func (u *User) getAge() int {
	return int(u.Age)
}

func (u *User) setStatus(status bool) bool {
	fmt.Printf("testing pointer %+v\n", &u)
	u.Status = status
	return u.Status
}
