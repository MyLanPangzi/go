package main

import (
	"fmt"
	"math"
)

func main() {
	v := Vertex7{3, 4}
	fmt.Println(v, v.Abs())
	v.Scale(10)
	fmt.Println(v, v.Abs())
}

type Vertex7 struct {
	X, Y float64
}

func (v *Vertex7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex7) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
