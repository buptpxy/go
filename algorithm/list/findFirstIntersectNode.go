/*
两个单链表相交的一系列问题
【 题目】 在本题中， 单链表可能有环， 也可能无环。
 给定两个单链表的头节点 head1和head2， 这两个链表可能相交， 也可能不相交。
 请实现一个函数， 如果两个链表相交， 请返回相交的第一个节点；
 如果不相交， 返回null 即可。
 【要求】 如果链表1的长度为N， 链表2的长度为M， 时间复杂度请达到 O(N+M)， 额外空间复杂度请达到O(1)。
【思路】1、判断链表是否有环：设定一个快指针一个慢指针，若快指针直接走完则无环。
		若两指针相遇，则快指针回到头结点重新出发，两指针再次相遇的节点就是进入环的第一个节点。
	2、如果这两个链表相交，则有三种可能状态：
			第一种：都无环，在某点相交后一起到达终点，则尾指针必定相同。
				   计算两条链长度的差值n，长的先走n步，返回两者第一次相遇的节点。
			第二种：都有环，且入环的节点相同。
				   计算从头结点到入环节点的长度差值n，长的先走n步，返回两者第一次相遇的节点。
			第三种：都有环，且不在一个节点进入环。
				   返回两条链中任意一个入环节点。
*/

package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func (head *Node) Append(value int) {
	p := head
	for p.next != nil {
		p = p.next
	}
	var newNode Node
	newNode.value = value
	p.next = &newNode
	newNode.next = nil
}

func IsLoop(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}
	fast := head.next.next //注意初始条件，只能是这一种情况
	slow := head.next
	for fast != slow {
		if fast == nil || fast.next == nil || fast.next.next == nil {
			return nil
		}
		fast = fast.next.next
		slow = slow.next
	}
	fast = head
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}
	return slow
}
func bothNoLoop(head1 *Node, head2 *Node) *Node {
	n := 0
	p1, p2 := head1, head2
	for p1.next != nil {
		n++
		p1 = p1.next
	}
	for p2.next != nil {
		n--
		p2 = p2.next
	}
	if p1 != p2 { //最后一个节点不相等则一定不相交
		return nil
	}
	p1, p2 = head1, head2
	if n > 0 {
		for i := 0; i < n; i++ {
			p1 = p1.next
		}
	} else {
		for i := 0; i < 0-n; i++ {
			p2 = p2.next
		}
	}
	for p1 != p2 && p1 != nil && p2 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p1
}

//入环节点相同的情况下一定相交
func bothSameLoop(head1 *Node, head2 *Node, loop *Node) *Node {
	n := 0
	p1, p2 := head1, head2
	for p1.next != loop {
		n++
		p1 = p1.next
	}
	for p2.next != loop {
		n--
		p2 = p2.next
	}
	p1, p2 = head1, head2
	if n > 0 {
		for i := 0; i < n; i++ {
			p1 = p1.next
		}
	} else {
		for i := 0; i < 0-n; i++ {
			p2 = p2.next
		}
	}
	for p1 != p2 && p1 != nil && p2 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p1
}

//loop1把自己的环走完如果还没碰到loop2则不相交
func bothNoSameLoop(head1, head2, loop1, loop2 *Node) *Node {
	p1 := loop1.next  //p1 == loop1时和loop2一定不相交，故可从loop1.next开始
	for p1 != loop1 { //不能在p1.next == loop1时退出，因为有可能p1.next == loop2
		if p1 == loop2 {
			return loop2
		}
		p1 = p1.next
	}
	return nil
}
func findIntersectNode(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1, loop2 := IsLoop(head1), IsLoop(head2)
	if loop1 == nil && loop2 == nil {
		return bothNoLoop(head1, head2)
	}
	if loop2 != nil && loop2 != nil {
		if loop1 == loop2 {
			return bothSameLoop(head1, head2, loop1)
		}
		return bothNoSameLoop(head1, head2, loop1, loop2)
	}
	return nil
}

func main() {
	head1 := &Node{0, nil}
	head1.Append(1)
	head1.Append(2)
	head1.Append(3)
	head1.next.next.next.next = head1.next
	head2 := &Node{0, nil}
	head2.next = head1.next.next.next
	node := findIntersectNode(head1, head2)
	fmt.Println(node)

}
