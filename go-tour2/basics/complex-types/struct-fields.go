package main

import "fmt"

type V struct {
	X int
	Y int
}

func main() {
	v := V{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
