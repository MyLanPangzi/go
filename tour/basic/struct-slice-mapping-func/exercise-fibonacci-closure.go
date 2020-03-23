package main

import "fmt"

func fibonacci() func() int {
	cur, next := 0, 1
	return func() int {
		tmp := cur
		cur, next = next, cur+next
		return tmp
	}

}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
