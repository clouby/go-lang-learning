package main

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

func main() {
	user := User{"Carlos", 27, "celb25@gmail.com"}

	user.age = 28

	fmt.Println(user)
}
