/*
【要求】实现二叉树的先序、 中序、 后序遍历， 包括递归方式和非递归方式
【思路】递归方式：每个节点都会被遍历三次，在第一次遍历的时候就打印则为先序遍历，第二次遍历打印为中序、第三次遍历打印为后序
非递归方式：
	先序遍历：打印顺序：中左右；
			1. 把当前节点压入栈中，
			2. 若栈不为空，栈顶元素出栈并打印，
			3. 如果出栈元素的右孩子不为空则右孩子入栈，如果左孩子不为空则左孩子入栈
			4. 回到2
	中序遍历：打印顺序：左中右；
			1. 把当前节点和当前节点的左边界都压入栈中，
			2. 若栈不为空，栈顶元素出栈并打印
			3. 若出栈元素的右孩子不为空则右孩子和右孩子的左边界全部入栈
			4. 回到2
	后序遍历：打印顺序：左右中；
			1. 把当前节点压入栈中，
			2. 若栈不为空，栈顶元素出栈并进入一个新的栈中
			3. 如果出栈元素的左孩子不为空则左孩子入栈，如果右孩子不为空则右孩子入栈
			4. 回到2
			5. 新的栈中元素依次出栈并打印
*/
package main

import (
	"fmt"
	"queue"
	"stack"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func recTraversal(head *Node) {
	if head == nil {
		return
	}
	// fmt.Println(head.value) //先序
	recTraversal(head.left)
	//fmt.Println(head.value) //中序
	recTraversal(head.right)
	fmt.Println(head.value) //后序
}
func preTraversal(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("PreOrder: ")
	s := stack.NewStack()
	n := head
	s.Push(n)
	for !s.IsEmpty() {
		n = s.Pop().(*Node)
		fmt.Printf("%d ", n.value)
		if n.right != nil {
			s.Push(n.right)
		}
		if n.left != nil {
			s.Push(n.left)
		}
	}
	fmt.Printf("\n")
}
func inTraversal(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("InOreder: ")
	s := stack.NewStack()
	n := head
	for n != nil {
		s.Push(n)
		n = n.left
	}
	for !s.IsEmpty() {
		n = s.Pop().(*Node)
		fmt.Printf("%d ", n.value)
		if n.right != nil {
			n = n.right
			for n != nil {
				s.Push(n)
				n = n.left
			}
		}
	}
	fmt.Printf("\n")
}
func PostTraversal(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("PostOrder: ")
	s1 := stack.NewStack()
	s2 := stack.NewStack()
	n := head
	s1.Push(n)
	for !s1.IsEmpty() {
		n = s1.Pop().(*Node)
		s2.Push(n)
		if n.left != nil {
			s1.Push(n.left)
		}
		if n.right != nil {
			s1.Push(n.right)
		}
	}
	for !s2.IsEmpty() {
		n = s2.Pop().(*Node)
		fmt.Printf("%d ", n.value)
	}
	fmt.Printf("\n")
}
func LevelTraversal(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("LevelOrder: ")
	q := queue.NewQueue()
	n := head
	q.Push(n)
	for !q.IsEmpty() {
		n = q.Pop().(*Node)
		fmt.Printf("%d ", n.value)
		if n.left != nil {
			q.Push(n.left)
		}
		if n.right != nil {
			q.Push(n.right)
		}
	}
	fmt.Printf("\n")
}
func main() {
	head := &Node{0, nil, nil}
	head.left = &Node{1, nil, nil}
	head.right = &Node{2, nil, nil}
	head.left.left = &Node{3, nil, nil}
	head.left.right = &Node{4, nil, nil}
	head.right.left = &Node{5, nil, nil}
	head.right.right = &Node{6, nil, nil}
	// recTraversal(head)
	preTraversal(head)
	inTraversal(head)
	PostTraversal(head)
	LevelTraversal(head)
}
