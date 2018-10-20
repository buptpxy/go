package main

/*
Q20. Linked List
	1. Make use of the package container/list to create a (doubly) linked list. Push thevalues 1, 2 and 4 to the list and then print it.
	2. Create your own linked list implementation. And perform the same actions as in question 1
*/
/*

 */
import (
	"container/list" //实现了一个循环双链表
	"fmt"
)

func main() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertAfter(2, e1)
	l.InsertBefore(3, e4)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \n", e.Value)
	}
}
