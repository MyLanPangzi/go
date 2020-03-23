package main

import (
	"fmt"
)

func main() {
	fmt.Println(split(17))

}

func split(i int) (x, y int) {
	x = i * 4 / 9
	y = i - x
	return
}
