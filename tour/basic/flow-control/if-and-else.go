package main

import (
	"fmt"
	"math"
)

func pow(i, n, lim float64) float64 {
	if v := math.Pow(i, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))

}