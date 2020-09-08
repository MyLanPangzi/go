package main

import (
	"fmt"
)

func main() {
	var r Reporter
	s := Student{1}
	t := Teacher{2}
	r = s
	desc(r)
	r = t
	desc(r)
	desc(s)
	desc(t)
}

type Reporter interface {
	Report() int
}
type Student struct {
	Id int
}
type Teacher struct {
	Id int
}

func (s Student) Report() int {
	return s.Id
}
func (t Teacher) Report() int {
	return t.Id
}

func desc(r Reporter) {
	fmt.Printf("%d %v %T\n", r.Report(), r, r)
}
