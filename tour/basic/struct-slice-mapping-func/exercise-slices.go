package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for i := range r {
		r[i] = make([]uint8, dy)
	}
	return r

}
func main() {
	pic.Show(Pic)
}
