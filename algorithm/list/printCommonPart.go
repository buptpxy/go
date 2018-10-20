/*
打印两个有序链表的公共部分
【 题目】 给定两个有序链表的头指针head1和head2， 打印两个链表的公共部分
*/
package main

import (
	"fmt"
	"newList"
)

func printCommonPart(head1 *newList.LinkNode, head2 *newList.LinkNode) {
	p1 := head1
	p2 := head2
	for p1 != nil && p2 != nil {
		if p1.Data.(int) < p2.Data.(int) {
			p1 = p1.Next
		} else if p1.Data.(int) > p2.Data.(int) {
			p2 = p2.Next
		} else {
			fmt.Println(p1.Data)
			p1 = p1.Next
			p2 = p2.Next
		}
	}
}

func main() {
	head1 := newList.LinkNode{0, nil}
	head2 := newList.LinkNode{0, nil}
	head1.Append(1)
	head1.Append(2)
	head1.Append(3)
	head1.Print()
	head2.Append(2)
	head2.Append(3)
	head2.Append(4)
	head2.Print()
	printCommonPart(&head1, &head2)
}
