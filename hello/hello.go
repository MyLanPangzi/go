package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"hiscat.com/hello/morestrings"
)

func main() {
	fmt.Print(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
