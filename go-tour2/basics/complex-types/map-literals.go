package main

import "fmt"

type MVertex2 struct {
	Lat, Long float64
}

func main() {
	m := map[string]MVertex2{
		"hello": MVertex2{1, 2},
		"world": MVertex2{3, 4},
	}
	fmt.Println(m)
}
