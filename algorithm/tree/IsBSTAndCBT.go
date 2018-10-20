/*
	判断一棵树是否是搜索二叉树，左子树都比自身小，右子树都比自身大，即中序遍历为升序
	判断一棵树是否是完全二叉树，
	完全二叉树：节点必须从左往右的排列，
	若只有右子树没有左子树则一定不是完全二叉树，若有左孩子没有右孩子，则层序遍历中，此节点后的节点只能是叶子节点
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

func IsBST(head *Node) bool {
	if head == nil {
		return false
	}
	fmt.Printf("InOreder: ")
	s := stack.NewStack()
	n := head
	for n != nil {
		s.Push(n)
		n = n.left
	}
	pre := s.Top().(*Node).value
	now := 0
	for !s.IsEmpty() {
		n = s.Pop().(*Node)
		now = n.value
		if pre > now {
			return false
		}
		pre = now
		// fmt.Printf("%d ", n.value)
		if n.right != nil {
			n = n.right
			for n != nil {
				s.Push(n)
				n = n.left
			}
		}
	}
	// fmt.Printf("\n")
	return true
}

func IsCBT(head *Node) bool {
	leaf := false
	if head == nil {
		return false
	}

	fmt.Printf("LevelOrder: ")
	q := queue.NewQueue()
	n := head
	q.Push(n)
	for !q.IsEmpty() {
		n = q.Pop().(*Node)
		// fmt.Printf("%d ", n.value)
		if leaf && (n.left != nil || n.right != nil) {
			return false
		}
		if n.left == nil && n.right != nil {
			return false
		}
		if n.left != nil && n.right == nil {
			leaf = true
		}
		if n.left != nil {
			q.Push(n.left)
		}
		if n.right != nil {
			q.Push(n.right)
		}
	}
	return true
}

func main() {
	head := &Node{3, nil, nil}
	head.left = &Node{4, nil, nil}
	head.right = &Node{5, nil, nil}
	head.left.left = &Node{0, nil, nil}
	head.left.right = &Node{2, nil, nil}
	head.right.left = &Node{4, nil, nil}
	head.right.right = &Node{6, nil, nil}
	// fmt.Println(IsBST(head))
	fmt.Println(IsCBT(head))
}
