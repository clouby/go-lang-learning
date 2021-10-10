package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	r := []bool{true, false, true, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{i: 45, b: false},
		{i: 1e3, b: true},
	}
	fmt.Println(s)
}
