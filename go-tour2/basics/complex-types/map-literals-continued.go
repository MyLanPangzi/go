package main

import "fmt"

type MVertex3 struct {
	Lat, Long float64
}

func main() {
	m := map[string]MVertex3{
		"hello": {1, 2},
		"world": {3, 4},
	}
	fmt.Println(m)
}
