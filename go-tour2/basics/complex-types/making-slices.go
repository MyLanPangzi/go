package main

import "fmt"

func main() {
	a := make([]int, 5)
	ps(a)
	b := make([]int, 0, 5)
	ps(b)
	c := b[:2]
	ps(c)
	//fmt.Println(c[2])
	//c[0]=2
	//ps(c)
	d := c[2:5]
	ps(d)
}
func ps(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
