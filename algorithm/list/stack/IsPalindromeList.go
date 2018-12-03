/*
判断一个链表是否为回文结构，头结点也存储数据
【 题目】 给定一个链表的头节点head， 请判断该链表是否为回文结构。
	例如：1->2->1， 返回true。 1->2->2->1， 返回true。
		15->6->15， 返回true。 1->2->3， 返回false。
	进阶： 如果链表长度为N， 时间复杂度达到O(N)， 额外空间复杂度达到O(1)。
	思路：
		【非进阶】：
			1. 使用一个快指针和一个慢指针，快指针一次走两步，慢指针一次走一步。直到快指针来到链尾部时，慢指针来到链中间。
			2. 慢指针继续遍历，并把后半段链表的元素存入栈中。
			3. 快指针回到链头，继续开始遍历。然后依次弹出栈顶元素与快指针所指节点元素比较。
		【进阶】：
			1. 使用一个快指针和一个慢指针，快指针一次走两步，慢指针一次走一步。直到快指针来到链尾部时，慢指针来到链中间。
			2. 把后半段链表反转
			3. 两个指针同时从头和尾开始遍历，并比较元素是否相同
			4. 把后半段链表反转回来
*/
package stack

import (
	"fmt"
	"newList"
	"stack"
)

//need extra O(n) space
func IsPalindrome1(head *newList.LinkNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	p := head
	s := stack.NewStack()
	for p != nil {
		s.Push(p.Data) //写成(*p).Data也对
		p = p.Next
	}
	p = head
	for !s.IsEmpty() {
		if s.Pop() != p.Data {
			return false
		} else {
			p = p.Next
		}
	}
	return true
}

//need extra O(n/2) space
func IsPalindrome2(head *newList.LinkNode) bool {
	if head == nil || head.Next == nil { //注意考虑head == nil || head.Next == nil的情况
		return true
	}
	fast := head //不能写成fast := head.Next.Next，因为head.Next.Next可能不存在
	slow := head.Next
	for fast.Next != nil && fast.Next.Next != nil { //fast.Next != nil 必须加上，不然fast.Next.Next可能不存在
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = head
	s := stack.NewStack()
	for slow != nil {
		s.Push(slow.Data)
		slow = slow.Next
	}
	for !s.IsEmpty() {
		if s.Pop() != fast.Data {
			return false
		} else {
			fast = fast.Next
		}
	}
	return true
}

////need extra O(1) space
func IsPalindrome3(head *newList.LinkNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast := head                                    //不能写成fast := head.Next.Next，因为head.Next.Next可能不存在
	slow := head                                    //此处的slow从head开始
	for fast.Next != nil && fast.Next.Next != nil { //fast.Next != nil 必须加上，不然fast.Next.Next可能不存在
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1 := slow
	// fmt.Println(p1.Data)
	if slow.Next != nil {
		p1 = slow.Next
		for p1 != nil && p1.Next != nil { //注意限定p1 != nil
			p2 := p1.Next
			p1.Next = slow
			slow = p1
			p1 = p2
		}
		p1.Next = slow
	}
	fast = head //fast为头节点
	slow.Next = nil
	for fast.Next != nil && p1.Next != nil {

		if fast.Data != p1.Data {
			return false
		} else {
			// fmt.Println(p1.Data)
			fast = fast.Next
			p1 = p1.Next
			// fmt.Println(p1.Data)
		}
	}
	if (p1.Next == fast || fast.Next == p1) && fast.Data != p1.Data {
		return false
	}
	return true
}
func main() {
	head := newList.LinkNode{0, nil}
	head.Append(1)
	head.Append(0)
	head.Append(1)
	head.Append(0)
	fmt.Printf("%v \n", IsPalindrome1(&head))
	fmt.Printf("%v \n", IsPalindrome2(&head))
	fmt.Printf("%v \n", IsPalindrome3(&head))
}
