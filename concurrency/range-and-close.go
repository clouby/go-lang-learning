package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	fmt.Println(cap(c), len(c))

	for i := range c {
		fmt.Println(i)
	}
}
