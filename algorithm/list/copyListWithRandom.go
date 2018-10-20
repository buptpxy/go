/*
复制含有随机指针节点的链表
【 题目】 一种特殊的链表节点类描述如下：
public class Node {
	public int value;
	public Node next;
	public	Node rand;
	public Node(int data){
		this.value = data;
	}
}
Node类中的value是节点值，next指针和正常单链表中next指针的意义一样，都指向下一个节点，
rand指针是Node类中新增的指针，这个指针可能指向链表中的任意一个节点，也可能指向null。
给定一个由Node节点类型组成的无环单链表的头节点head， 请实现一个函数完成这个链表中所有结构的复制， 并返回复制的新链表的头节点。(注意是新链表，只是值一样，不能是和原来链表一模一样的指针)
【思路】：可以把节点都放进一个map中，建立此节点和复制节点的映射，复制节点也会很容易的根据节点索引知道next和rand指针指向哪个节点

【进阶】：不使用额外的数据结构，只用有限几个变量，且在时间复杂度为 O(N)内完成原问题要实现的函数。
【思路】：把复制的新节点连接到被复制的节点后，这样新节点的next是此时当前自己的next，rand是旧节点的rand的next，然后再把链中的新节点挑选出来组成新的链

*/
package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
	rand  *Node
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
	newNode.rand = nil
}
func (head *Node) Print() {
	p := head
	for p != nil {
		if p.rand != nil {
			fmt.Printf("(%v->%v)->", p.value, p.rand.value)
		} else {
			fmt.Printf("(%v->%v)->", p.value, nil)
		}

		p = p.next

	}
	fmt.Printf("nil\n")
}
func (head *Node) GetLength() int {
	if head == nil {
		return 0
	}
	p := head
	length := 0
	for p != nil {
		length++
		p = p.next
	}
	return length
}
func copyList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	p := head
	for p != nil {
		nodeMap[p] = &Node{p.value, nil, nil}
		p = p.next
	}
	for oldN, newN := range nodeMap {
		newN.next = nodeMap[oldN.next]
		newN.rand = nodeMap[oldN.rand]
	}
	return nodeMap[head]
}

func copyList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	length := head.GetLength()
	for i := 0; i < length; i++ {
		newNode := Node{p.value, p.next, nil}
		p.next = &newNode
		p = p.next.next
	}
	// head.Print()
	p = head
	q := head.next
	for p != nil && q != nil {
		if p.rand != nil {
			q.rand = p.rand.next
		}
		p = p.next
		q = q.next
	}
	// head.Print()
	p = head
	q = head.next
	newHead := q
	for q != nil && q.next != nil {
		p.next = q.next
		q.next = q.next.next
		p = p.next //循环时不要忘掉变量的自增！！！
		q = q.next
	}
	p.next = nil

	return newHead
}

func main() {
	head := Node{0, nil, nil}
	head.Append(1)
	head.Append(2)
	head.Append(3)
	head.rand = head.next.next
	head.next.rand = head.next.next.next
	head1 := copyList1(&head)
	head2 := copyList2(&head)
	head.Print()
	head1.Print()
	head2.Print()
}
