package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	// var arr = []int{3, 5}
	c := NewStack()
	c.Push(3, 5)
	if c.Pop() != 5 {
		t.Log("Pop doesn't give 5")
		t.Fail()
	}
	if c.Top() != 3 {
		t.Log("Top doesn't give 3")
		t.Fail()
	}
}
