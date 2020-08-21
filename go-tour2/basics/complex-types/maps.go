package main

import "fmt"

type MVertex struct {
	Lat, Long float64
}

func main() {
	var m map[string]MVertex
	fmt.Println(m)
	m = make(map[string]MVertex)
	m["hello"] = MVertex{1, 2}
	fmt.Println(m)
	fmt.Println(m["hello"])
}
