package main

import "fmt"

func main() {
	p := Person{"xiaoming"}
	var g Greeting
	g = p
	fmt.Println(g.Hi())
}

type Greeting interface {
	Hi() string
}
type Person struct {
	Name string
}

func (p Person) Hi() string {
	return p.Name + " : " + "hello"
}
