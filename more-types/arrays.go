package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Carlos"
	a[1] = "Carla"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	v := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(v)
}
