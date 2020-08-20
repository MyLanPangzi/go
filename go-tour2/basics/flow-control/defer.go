package main

import "fmt"

func main() {
	defer fmt.Println(f(), f())
	fmt.Println("hello")
}
func f() int {
	fmt.Println("world")
	return 1
}
