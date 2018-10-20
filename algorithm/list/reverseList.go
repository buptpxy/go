/*
反转单向和双向链表
【 题目】 分别实现反转单向链表和反转双向链表的函数。
【 要求】 如果链表长度为N， 时间复杂度要求为O(N)， 额外空间复杂度要求为O(1)
*/
package main

import (
	// "doubleList"
	"fmt"
	"list"
)

func ReverseList(list list.List) {
	p := list.Root.Next
	q := p.Next
	list.Root.Next = nil
	p.Next = nil
	for q != nil {
		r := q.Next
		q.Next = p
		p = q
		q = r
	}
	list.Root.Next = p
}

// func ReverseDoubleList(dbList doubleList.List) {
// 	p := dbListnext
// 	q := p.Next
// 	dbList.Next = nil
// 	p.Next = nil
// 	for q != nil {
// 		r := q.Next
// 		q.Next = p
// 		p.prev = q
// 		p = q
// 		q = r
// 	}
// 	dbList.Next = p
// 	p.prev = &(dbList)
// }
// func PrintDoubleList(dl doubleList.List) {
// 	p := dl.Next
// 	fmt.Printf("root<=>")
// 	for i := 0; i < dl.Len; i++ {
// 		fmt.Printf("%v<=>", (*p).Value)
// 		p = p.Next
// 	}
// 	fmt.Printf("nil\n")
// }
func main() {
	var l list.List
	l.Init()
	l.Insert(0, "a")
	l.Insert(0, "b")
	l.Insert(0, "c")
	fmt.Println("before reverse:")
	l.Print()
	ReverseList(l)
	fmt.Println("after reverse:")
	l.Print()

	// var dl doubleList.List
	// dl.New()
	// dl.InsertBefore("a", dl.Next)
	// dl.InsertBefore("b", dl.Next)
	// dl.InsertBefore("c", dl.Next)
	// fmt.Println("before reverse:")
	// PrintDoubleList(dl)
	// ReverseDoubleList(dbList)
	// fmt.Println("after reverse:")
	// PrintDoubleList(dl)

}
