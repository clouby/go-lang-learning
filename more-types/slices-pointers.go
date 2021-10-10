package main

import "fmt"

func main() {
	names := [4]string{
		"Carlos",
		"Carla",
		"Baku",
		"Kora",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	b[0] = "666"
	fmt.Println(a, b)
	fmt.Println(names)
}
