package main

import "fmt"

func main() {
	pow := make([]int, 4)
	for i := range pow {
		pow[i] = 1 << i
	}
	for _, e := range pow {
		fmt.Println(e)
	}
}
