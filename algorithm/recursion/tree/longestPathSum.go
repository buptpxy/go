/*
二叉树节点间距离的概念： 二叉树一个节点到另一个节点间最短线路上的节点数量， 叫做两个节点间的距离。
给定一棵二叉树的头节点head， 请返回这棵二叉树上的最大距离。
【思路】
1、从每个节点收集可能的路径长度信息
2、三种情况：1）最大距离存在左子树上 2）最大距离存在右子树上 3）要穿过当前节点，最大距离为左子树最大距离加右子树最大距离加1
3、每个节点要收集的信息：最大距离d，树高h
4、先假设左右节点都可返回需要的信息，然后把三种情况得到的信息处理成当前节点要返回的信息
5、baseCase: head==nil 时d=0,h=0
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

type res struct {
	distance int
	height   int
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//O(2n-1)
func longestPathSum(head *Node) res {
	fmt.Println("rec")
	if head == nil {
		return res{0, 0}
	}
	leftRes := longestPathSum(head.left)
	rightRes := longestPathSum(head.right)
	maxD := leftRes.height + rightRes.height + 1
	maxH := 0
	if leftRes.distance > rightRes.distance {
		maxH = leftRes.height + 1
	} else {
		maxH = rightRes.height + 1
	}
	maxD = max(maxD, max(leftRes.distance, rightRes.distance))
	return res{maxD, maxH}
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
	head1 := &Node{10, nil, nil}
	head1.left = &Node{11, nil, nil}
	head1.right = &Node{9, nil, nil}
	head1.left.left = &Node{13, nil, nil}
	head1.left.right = &Node{17, nil, nil}
	head1.right.left = &Node{15, nil, nil}
	head1.right.right = &Node{16, nil, nil}
	head1.right.right.right = &Node{19, nil, nil}
	head1.Print()
	r := longestPathSum(head1)
	fmt.Println(r)

	head2 := &Node{10, nil, nil}
	head2.left = &Node{11, nil, nil}
	head2.right = &Node{9, nil, nil}
	head2.right.left = &Node{13, nil, nil}
	head2.right.right = &Node{17, nil, nil}
	head2.right.left.left = &Node{15, nil, nil}
	head2.right.left.left.left = &Node{18, nil, nil}
	head2.right.right.right = &Node{16, nil, nil}
	head2.right.right.right.right = &Node{19, nil, nil}
	head2.Print()
	r2 := longestPathSum(head2)
	fmt.Println(r2)

}
