package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx, dy)
}

func main() {
	pic.Show(Pic)
}
