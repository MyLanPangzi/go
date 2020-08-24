package main

import "fmt"

func f(i interface{}) {
	fmt.Printf("%v %T\n", i, i)
}
func main() {
	var i interface{}
	f(i)
	i = 42
	f(i)
	i = "hello"
	f(i)
}
