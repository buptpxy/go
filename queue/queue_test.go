package queue

import (
	"testing"
)

func TestStack(t *testing.T) {
	// var arr = []int{3, 5}
	c := NewQueue()
	c.Push(3, 5)
	if c.Pop() != 3 {
		t.Log("Pop doesn't give 3")
		t.Fail()
	}
	if c.Top() != 5 {
		t.Log("Top doesn't give 5")
		t.Fail()
	}
}
