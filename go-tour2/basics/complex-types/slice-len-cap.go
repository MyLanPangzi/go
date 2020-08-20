package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	printSlices(s)

	s = s[:0]
	printSlices(s)

	s = s[:4]
	printSlices(s)

	s = s[2:]
	printSlices(s)
	//s = s[:7]
	//printSlices(s)
	//s = s[:6]
	s = s[0:5]
	printSlices(s)
}
func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
