package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	answer := "Answer"
	m[answer] = 42
	fmt.Println("The value", m[answer])

	delete(m, answer)
	v, ok := m[answer]
	fmt.Println("The value", v, "Present?", ok)
}
