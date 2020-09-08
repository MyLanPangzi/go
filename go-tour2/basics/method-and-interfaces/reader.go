package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "Hello, Reader"
	r := strings.NewReader(s)
	fmt.Println(len(s))
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%d err=%v,b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%v\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
