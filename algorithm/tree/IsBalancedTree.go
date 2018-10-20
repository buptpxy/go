/*
【要求】判断一棵二叉树是否是平衡二叉树。平衡二叉树：左右子树的高度差不超过1
【思路】如果此树为平衡二叉树则它的左右子树都为平衡二叉树，故可用递归。
		记录以每个节点为根的树的高度值，比较左孩子与右孩子的差值是否超过1。
*/
package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func IsBalanceTree(head *Node, level int, flag []bool) int {
	if head == nil {
		return level
	}
	hL := IsBalanceTree(head.left, level+1, flag)
	if flag[0] == false {
		return 0
	}
	hR := IsBalanceTree(head.right, level+1, flag)
	if flag[0] == false {
		return 0
	}
	if math.Abs(float64(hL-hR)) > 1 {
		flag[0] = false
	}
	return int(math.Max(float64(hL), float64(hR)))
}

//递归要素：终止条件，传递值
func getHeight(head *Node, level int) int {
	if head == nil {
		return level
	}
	hL := getHeight(head.left, level+1)
	hR := getHeight(head.right, level+1)
	return int(math.Max(float64(hL), float64(hR)))
}
func main() {
	head := &Node{0, nil, nil}
	head.left = &Node{1, nil, nil}
	// head.right = &Node{2, nil, nil}
	head.left.left = &Node{3, nil, nil}
	head.left.left.left = &Node{4, nil, nil}
	// head.left.right = &Node{4, nil, nil}
	// head.right.left = &Node{5, nil, nil}
	// head.right.right = &Node{6, nil, nil}
	flag := []bool{true}
	height := IsBalanceTree(head, 0, flag)
	// height := getHeight(head, 0)
	fmt.Println(height)
	fmt.Println(flag[0])
}
