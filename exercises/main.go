package main

import (
	"fmt"
	"stack"
)

func main() {
	// var arr = []int{3, 5}
	c := stack.NewStack()
	c.Push(3, 5)
	c.Push("a", "b")
	fmt.Printf("%s\n", c.Pop())
}
