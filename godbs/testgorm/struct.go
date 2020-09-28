package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {

	var user User
	user.Name = "adamy"
	user.Age = 123

	fmt.Println(user)
	fmt.Println(user.Name)

	puser := &User{}
	puser.Name = "palalkjdf"
	puser.Age = 1234

	fmt.Println(puser)
	fmt.Println(puser.Name)
}
