package main

import "fmt"

type User struct {
	name   string
	age    int
	email  string
	galaxy int
}

func main() {
	u := User{"Carlos", 27, "celb25@gmail.com", 1e5}
	fmt.Println(u)
	p := &u
	p.galaxy = 1e9
	fmt.Println(u)
}
