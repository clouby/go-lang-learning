package main

import "fmt"

func main() {
	var s []int
	fmt.Println(len(s), cap(s), s)
	if s == nil {
		fmt.Println("nil!")
	}
}
