package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	m["hello"] = 19
	fmt.Println(m["hello"])
	delete(m, "hello")
	fmt.Println(m["hello"])
	_, ok := m["hello"]
	if !ok {
		fmt.Println("not exists")
	}
}
