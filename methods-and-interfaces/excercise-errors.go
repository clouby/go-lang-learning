package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	}
	return fmt.Sprint(float64(e))
}

func Sqrt(x ErrNegativeSqrt) ErrNegativeSqrt {
	var z ErrNegativeSqrt = 1.0

	for i := 0; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
