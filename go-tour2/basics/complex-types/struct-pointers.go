package main

import "fmt"

type V2 struct {
	X int
	Y int
}

func main() {
	v := V2{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
