package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}

func pow(i, n, lim float64) float64 {
	if v := math.Pow(i, n); v < lim {
		return v
	}
	return lim
}
