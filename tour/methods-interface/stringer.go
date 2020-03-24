package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.Name, p.Age)

}
func main() {
	hello := &Person{"Hello", 32}
	world := &Person{"World", 23}
	fmt.Println(hello, world)

}
