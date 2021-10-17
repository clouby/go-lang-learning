package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (s T) M() {
	fmt.Println(s.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
