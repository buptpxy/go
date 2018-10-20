package list

import (
	"testing"
)

func TestInsert(t *testing.T) {
	var l List
	l.Init() //实际上被转换为了(&l).Init()
	l.Insert(0, "a")
	if l.Get(0) != "a" {
		t.Log("after insert l.Get(0) should be a")
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	l := New()
	l.Insert(0, "a")
	l.Insert(0, "b")

	if l.Remove(0) != "b" {
		t.Log("l.Remove(0) should be b")
		t.Fail()
	}
}
