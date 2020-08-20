package main

import "fmt"

type V3 struct {
	X, Y int
}

var (
	v1 = V3{1, 2}
	v2 = V3{X: 1}
	v3 = V3{}
	p  = &V3{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
