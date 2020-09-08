package main

import (
	"fmt"
	"math"
)

func main() {
	var a Abser
	v := Vertex8{3, 4}
	f := MyFloat1(-4)
	a = &v
	fmt.Println(a.Abs())
	a = f
	fmt.Println(a.Abs())
	//a=v
}

type Abser interface {
	Abs() float64
}
type Vertex8 struct {
	X, Y float64
}

func (v *Vertex8) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat1 float64

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
