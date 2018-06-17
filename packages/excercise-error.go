package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	const limit = 1.0 / float64(1<<31)
	fmt.Println(limit)
	z := x/2
	for (z*z - x > limit || z*z - x < -limit) {
		z -= (z*z - x) / (2*z)

		fmt.Println(z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))	
	fmt.Println(Sqrt(-2))
}