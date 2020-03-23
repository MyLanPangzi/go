package main

import "fmt"

func main() {
	sum := 10
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
