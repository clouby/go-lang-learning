package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good night!")
	}
}
