package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 5, 6, 77, 8}
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
}
