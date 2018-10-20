/*
已知一棵完全二叉树， 求其节点的个数
【要求】 时间复杂度低于O(N)， N为这棵树的节点个数
【思路】可利用满二叉树的节点数计算公式，若满二叉树的高度为h，则节点数为2^h-1
		对于完全二叉树来说，如果一个节点的右子树的最左边界达到了最后一层，则其左子树必为满二叉树，再利用递归计算右子树的节点个数
		否则其右子树必为满二叉树，再利用递归计算其左子树的节点个数
*/
package main

import (
	"fmt"
	// "math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func mostLeftHeight(head *Node, level int) int {
	if head == nil {
		return 0
	}
	p := head
	for p.left != nil {
		level++
		p = p.left
	}
	return level
}

func nodesNum(head *Node, level int, height int) int {
	if level == height {
		return 1
	}
	if mostLeftHeight(head.right, level+1) == height {
		//不用写成count = count + ... ，递归函数会自己加上之前的数，递归函数只用写成一层的做法
		//位移操作符右边只能是无符号整数
		return 1<<uint(height-level) + nodesNum(head.right, level+1, height)
	}
	return 1<<uint(height-level-1) + nodesNum(head.left, level+1, height)

}
func powInt(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * powInt(x, y-1)
}
func main() {
	head := &Node{3, nil, nil}
	head.left = &Node{4, nil, nil}
	head.right = &Node{5, nil, nil}
	head.left.left = &Node{0, nil, nil}
	head.left.right = &Node{2, nil, nil}
	// head.right.left = &Node{4, nil, nil}
	// head.right.right = &Node{6, nil, nil}
	height := mostLeftHeight(head, 1)
	// fmt.Println(height)
	fmt.Println(nodesNum(head, 1, height))
	// fmt.Println(powInt(2, 6))

}
