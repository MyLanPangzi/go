package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1)
	ps2(s)
	s = append(s, 2)
	ps2(s)
	s = append(s, 2, 3, 4)
	ps2(s)
}
func ps2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
