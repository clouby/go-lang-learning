package main

import (
	"fmt"
	"math"
)

type MyFLoat float64

func (f MyFLoat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFLoat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs(), a)
}
