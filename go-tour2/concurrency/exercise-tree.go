package main

import (
	"fmt"
	"math/rand"
)

func main() {
	tree := New(10)
	ch := make(chan int)
	go Walk(tree, ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(1), New(2)))
	fmt.Println(Same(New(2), New(2)))
}

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	walkImpl(t, ch)
	close(ch)
}
func walkImpl(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
	return false
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
