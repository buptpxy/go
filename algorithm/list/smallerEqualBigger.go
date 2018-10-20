/*
将单向链表按某值划分成左边小、 中间相等、 右边大的形式
【 题目】 给定一个单向链表的头节点head， 节点的值类型是整型， 再给定一个整数pivot。
	实现一个调整链表的函数， 将链表调整为左部分都是值小于 pivot的节点， 中间部分都是值等于pivot的节点， 右部分都是值大于 pivot的节点。
	除这个要求外， 对调整后的节点顺序没有更多的要求。
	例如： 链表9->0->4->5->1， pivot=3。 调整后链表可以是1->0->4->9->5， 也可以是0->1->9->5->4。
	总之， 满足左部分都是小于3的节点， 中间部分都是等于3的节点（ 本例中这个部分为空） ， 右部分都是大于3的节点即可。 对某部分内部的节点顺序不做要求。
	额外空间复杂度不作要求
	【思路】把节点都放入一个数组中，把数组partition后，再把数组里面的节点串起来

【进阶】 在原问题的要求之上再增加如下两个要求。
	在左、 中、 右三个部分的内部也做顺序要求， 要求每部分里的节点从左到右的顺序与原链表中节点的先后次序一致。
	例如： 链表9->0->4->5->1， pivot=3。调整后的链表是0->1->9->4->5。
	在满足原问题要求的同时，左部分节点从左到右为0、1。在原链表中也 是先出现0，后出现1；
	中间部分在本例中为空，不再讨论；
	右部分节点从左到右为9、4、5。在原链表中也是先出现9，然后出现4，最后出现5。

	如果链表长度为N， 时间复杂度请达到O(N)， 额外空间复杂度请达到O(1)。

	【思路】找到第一个小于、等于、大于给定数的节点，把剩下的节点分别连接到这三个节点后面
*/
package main

import (
	// "fmt"
	"newList"
)

func arrPartition(arr []newList.LinkNode, pivot int) {
	s, b, i := -1, len(arr), 0
	for i < b {
		if arr[i].Data.(int) < pivot {
			arr[i], arr[s+1] = arr[s+1], arr[i]
			s++
			i++
		} else if arr[i].Data.(int) > pivot {
			arr[i], arr[b-1] = arr[b-1], arr[i]
			b--
		} else {
			i++
		}
	}
}

//need extra O(n) space
func listPartition1(head *newList.LinkNode, pivot int) newList.LinkNode {
	if head == nil {
		return *head
	}
	nodeArr := make([]newList.LinkNode, head.GetLength())
	for i := 0; i < len(nodeArr); i++ {
		nodeArr[i] = *head
		head = head.Next
	}
	arrPartition(nodeArr, pivot)
	// fmt.Println(nodeArr)
	for i := 1; i < len(nodeArr); i++ {
		nodeArr[i-1].Next = &nodeArr[i]
	}
	nodeArr[len(nodeArr)-1].Next = nil
	return nodeArr[0]
}

//need extra O(1) space
func listPartition2(head *newList.LinkNode, pivot int) newList.LinkNode {
	var node *newList.LinkNode
	smallH, smallT, bigH, bigT, equalH, equalT := node, node, node, node, node, node
	if head == nil {
		return *head
	}
	p := head
	for p != nil {
		// fmt.Println("p: ", p)
		if p.Data.(int) < pivot {
			if smallH == nil {
				smallH = p
				smallT = p
				// fmt.Println("smallT: ", smallT)
			} else {

				smallT.Next = p
				smallT = p
				// fmt.Println("smallT: ", smallT)
			}
		} else if p.Data.(int) > pivot {
			if bigH == nil {

				bigH = p
				bigT = p
				// fmt.Println("bigT: ", bigT)
			} else {

				bigT.Next = p
				bigT = p
				// fmt.Println("bigT: ", bigT)
			}
		} else {
			if equalH == nil {

				equalH = p
				equalT = p
				// fmt.Println("equalT: ", equalT)
			} else {

				equalT.Next = p
				equalT = p
				// fmt.Println("equalT: ", equalT)
			}
		}
		p = p.Next
	}
	if smallT != nil {
		smallT.Next = nil
	}
	if equalT != nil {
		equalT.Next = nil
	}
	if bigT != nil {
		bigT.Next = nil
	}

	// smallH.Print()
	// equalH.Print()
	// bigH.Print()
	if smallH == nil {
		if equalH == nil {
			if bigH != nil {
				head = bigH
			}
		} else {
			if bigH != nil {
				equalT.Next = bigH
				head = equalH
			} else {
				head = equalH
			}
		}
	} else {
		if equalH == nil {
			if bigH != nil {
				smallT.Next = bigH
			}
		} else {
			if bigH != nil {
				smallT.Next = equalH
				equalT.Next = bigH
			} else {
				smallT.Next = equalH
			}
		}
		head = smallH
	}
	return *head
}
func main() {
	head1 := newList.LinkNode{9, nil}
	head1.Append(0)
	head1.Append(4)
	head1.Append(3)
	head1.Append(5)
	head1.Append(1)
	head1.Print()
	head1 = listPartition1(&head1, 3)
	head1.Print()

	head2 := newList.LinkNode{9, nil}
	head2.Append(0)
	head2.Append(4)
	head2.Append(3)
	head2.Append(5)
	head2.Append(1)
	//head := listPartition2(&head2, 3)
	//head.Print()

	head2 = listPartition2(&head2, 3) //不能直接用head2
	head2.Print()

}
