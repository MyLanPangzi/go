package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

func Sqrt(x float64) float64 {
	z := 1.0
	for z -= (z*z - x) / (2 * z); z*z > x; {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
