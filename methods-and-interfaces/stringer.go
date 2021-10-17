package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Carlos Lopez", 27}
	z := Person{"Carla Arrazola", 25}

	fmt.Println(a, z)
}
