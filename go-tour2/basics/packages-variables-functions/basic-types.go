package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe          = false
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("type:%T, v:%v\n", ToBe, ToBe)
	fmt.Printf("type:%T, v:%v\n", MaxInt, MaxInt)
	fmt.Printf("type:%T, v:%v\n", z, z)
}
