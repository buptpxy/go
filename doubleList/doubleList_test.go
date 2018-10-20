package doubleList

import (
	"testing"
)

func TestInsert(t *testing.T) {
	l := New()
	e4 := l.PushBack(4)
	if l.Back() != e4 {
		t.Log("l.Back() should be e4")
		t.Fail()
	}
	e1 := l.PushFront(1)
	if l.Front() != e1 {
		t.Log("l.Front() should be e1")
		t.Fail()
	}
	e2 := l.InsertAfter(2, e1)
	if e1.Next() != e2 {
		t.Log("e1.Next() should be e2")
		t.Fail()
	}
	e3 := l.InsertBefore(3, e4)
	if e4.Prev() != e3 {
		t.Log("e4.Prev() should be e3")
		t.Fail()
	}
}
