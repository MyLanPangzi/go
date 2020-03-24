package main

import "fmt"

type I interface {
	M()
}
type Hello interface {
	hello()
}

type T struct {
	S string
}

func (t T) hello() {
	fmt.Println("hello", t.S)
}

func (t T) M() {
	fmt.Println(t.S)
}
func main() {
	var i I = T{"hello"}
	i.M()

	var hello Hello = T{"world"}
	hello.hello()
}
