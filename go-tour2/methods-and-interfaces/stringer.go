package main

import "fmt"

type StringerPerson struct {
	Name string
	Age  int
}

func (p *StringerPerson) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	hello := StringerPerson{"Hello", 18}
	world := StringerPerson{"world", 20}
	fmt.Println(hello, world)
}
