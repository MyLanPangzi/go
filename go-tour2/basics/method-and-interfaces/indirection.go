package main

import "fmt"

func main() {
	v := Vertex5{3, 4}
	fmt.Println(v)
	ScaleFunc(&v, 10)
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
	p := &v
	p.Scale(10)
	fmt.Println(p)
	ScaleFunc(p, 10)
	fmt.Println(p)
}

type Vertex5 struct {
	X, Y float64
}

func ScaleFunc(v *Vertex5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func (v *Vertex5) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
