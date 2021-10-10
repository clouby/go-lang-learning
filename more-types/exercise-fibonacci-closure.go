package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, dprev, seq, n := 0, 1, 0, 0

	return func() int {
		if n > 1 {
			seq = prev + dprev
			prev = dprev
			dprev = seq
		} else {
			seq = n
			n = n + 1
		}

		return seq
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
