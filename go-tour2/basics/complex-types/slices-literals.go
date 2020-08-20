package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	b := []bool{true, false}
	fmt.Println(b)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println(s)
}
