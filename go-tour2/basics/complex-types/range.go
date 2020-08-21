package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8}
	for i, e := range pow {
		fmt.Printf("2^%d=%d\n", i, e)
	}
}
