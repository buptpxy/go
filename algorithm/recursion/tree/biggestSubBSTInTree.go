/*
【树型dp的时间复杂度都为O（n）                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         】
求一颗树里面最大的搜索二叉树子树，搜索二叉树：左子树的最大值一定小于当前节点，右子树的最小值一定大于当前节点。
【思路】递归求每个节点的最大搜索二叉树子树
1、找出可能情况：1）最大搜索二叉树存在于左子树中 2）最大搜索二叉树存在于右子树中 3）最大搜索二叉树为当前树本身：左子树中最大搜索二叉树的头为当前节点左孩子，右...为右孩子，左子树的最大值小于当前节点值，右子树的最小值大于当前节点值
2、每个节点需要收集的信息：最大搜索二叉子树的大小size，最大搜索二叉子树的头head，以当前节点为根的树的最小值min，以当前节点为根的树的最大值max。
3、baseCase:当前节点n==nil，则返回的值应为：size=0,head=nil,min=INT_Max,max=INT_MIN
*/
package main

import (
	"fmt"
	"strconv"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

type Node struct {
	value int
	left  *Node
	right *Node
}

type res struct {
	size int
	head *Node
	min  int
	max  int
}

//O(2n-1)
func biggestBST(n *Node) res {
	fmt.Println("rec")
	if n == nil {
		return res{0, nil, INT_MAX, INT_MIN}
	}
	leftRes := biggestBST(n.left)
	rightRes := biggestBST(n.right)
	maxSize := 0
	nSize := 0
	minN := INT_MAX
	maxN := INT_MIN
	maxHead := n
	if leftRes.head == n.left && rightRes.head == n.right && leftRes.max < n.value && rightRes.min > n.value {
		nSize = leftRes.size + 1 + rightRes.size
	}
	if leftRes.size > rightRes.size {
		maxSize = max(nSize, leftRes.size)
		maxHead = leftRes.head
	} else {
		maxSize = max(nSize, rightRes.size)
		maxHead = rightRes.head
	}
	if maxSize == nSize {
		maxHead = n
	}
	minN = min(n.value, min(leftRes.min, rightRes.min))
	maxN = max(n.value, max(leftRes.max, rightRes.max))
	return res{maxSize, maxHead, minN, maxN}
}

//this is a bad method
func biggestBST1(n *Node) res {
	fmt.Println("rec")
	if n == nil {
		return res{0, nil, INT_MAX, INT_MIN}
	}
	leftRes := biggestBST1(n.left)
	rightRes := biggestBST1(n.right)
	if leftRes.head == n.left && rightRes.head == n.right && leftRes.max < n.value && rightRes.min > n.value {
		return res{leftRes.size + rightRes.size + 1, n, leftRes.min, rightRes.max}
	} else if leftRes.size > rightRes.size {
		return res{leftRes.size, leftRes.head, min(n.value, min(leftRes.min, rightRes.min)), max(n.value, max(leftRes.max, rightRes.max))} //这里返回的值在上层函数中并没有用到，一返回就被丢掉了
	} else {
		return res{rightRes.size, rightRes.head, min(n.value, min(leftRes.min, rightRes.min)), max(n.value, max(leftRes.max, rightRes.max))}
	}
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
	head := &Node{1, nil, nil}
	head.left = &Node{5, nil, nil}
	head.right = &Node{12, nil, nil}
	head.left.left = &Node{4, nil, nil}
	head.left.right = &Node{6, nil, nil}
	head.right.left = &Node{11, nil, nil}
	head.right.right = &Node{13, nil, nil}
	//n1:=biggestBST1(head).head
	//n1.Print()
	n := biggestBST(head).head
	n.Print()
	//r:=biggestBST(head)
	//fmt.Println(r)
}
