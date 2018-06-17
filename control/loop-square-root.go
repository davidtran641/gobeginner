package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	const limit = 1.0 / float64(1<<31)
	fmt.Println(limit)
	z := x/2
	for (z*z - x > limit || z*z - x < -limit) {
		z -= (z*z - x) / (2*z)

		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1000000))	
}