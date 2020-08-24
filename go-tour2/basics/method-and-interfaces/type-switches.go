package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d * 2 = %v\n", v, v)
	case string:
		fmt.Printf("%s len is %d\n", v, len(v))
	default:
		fmt.Printf("i don't know about type %T\n", v)
	}
}
func main() {
	do(10)
	do("string")
	do(true)
}
