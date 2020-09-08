package main

import "fmt"

type I2 interface {
	M()
}

func desc2(i I2) {
	fmt.Printf("%v %T", i, i)
}
func main() {
	var i I2
	desc2(i)
	//i.M()
}
