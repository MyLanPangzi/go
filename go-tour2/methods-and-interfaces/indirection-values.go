package main

import (
	"fmt"
	"math"
)

func main() {
	v := Vertex6{3, 4}
	fmt.Println(v.Abs())
	fmt.Println((&v).Abs())
	//fmt.Println(AbsFunc(&v))
	fmt.Println(AbsFunc(v))
}

type Vertex6 struct {
	X, Y float64
}

func (v Vertex6) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func AbsFunc(v Vertex6) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
