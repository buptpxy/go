/*
如何直观的打印一颗二叉树
【思路】使用类似中序遍历的顺序，把树横着打印，先打印右节点，再打印中节点，再打印左节点
*/
package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func getSpace(num int) string {
	var s string
	for i := 0; i < num; i++ {
		s = s + " "
	}
	return s
}

func printInOrder(head *Node, height int, to string, lenT int) {
	if head == nil {
		return
	}
	val := strconv.Itoa(head.value)
	lenM := len(val)
	// fmt.Println(lenM)
	lenL := (lenT - lenM) / 2
	lenR := lenT - lenM - lenL
	printInOrder(head.right, height+1, "v", 10)
	val = getSpace(height*lenT) + getSpace(lenL) + to + val + to + getSpace(lenR)
	fmt.Println(val)
	printInOrder(head.left, height+1, "^", 10)
}

func (head *Node) Print() {
	fmt.Println("BinaryTree:")
	printInOrder(head, 0, "", 10)
}

func main() {
	head := &Node{10, nil, nil}
	head.left = &Node{11, nil, nil}
	head.right = &Node{12, nil, nil}
	head.left.left = &Node{13, nil, nil}
	head.left.right = &Node{14, nil, nil}
	head.right.left = &Node{15, nil, nil}
	head.right.right = &Node{16, nil, nil}
	head.Print()
}
