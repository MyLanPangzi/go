package main

import "io"

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := range b {
		out := b[i]
		switch {
		case out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm':
			out += 13
		case out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z':
			out -= 13
		}
		b[i] = out
	}
	return n, err
}

func main() {

}
