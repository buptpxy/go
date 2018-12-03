/*
构造数组的MaxTree
	一个数组的MaxTree定义如下：
	MaxTree是一颗二叉树，数组的每一个值对应一个二叉树节点。
	包括MaxTree树在内的任意一颗子树上，值最大的节点都是树的头
	给定一个没有重复元素的数组arr，写出生成这个数组的MaxTree的函数，要求如果数组长度为N，则时间复杂度为O(N)，空间复杂度为O(N)
思路：
	可以利用栈底最大的单调栈，得出数组的每个元素左右最近的比它大的值，
	左右大值都为-1的是整个树的头，有一边为-1的根节点为不为-1的那边
	其他节点的父节点为它左右大值中较大的那个数
*/
package main

import (
	"fmt"
	"github.com/pengpeng1314/go/stack"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func getNearMax(arr []int) [][]int {
	n := len(arr)
	if n == 0 {
		return nil
	}
	s := stack.NewStack()
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, 3)
		res[i][0] = arr[i]
		for !s.IsEmpty() && arr[i] > arr[s.Top().(int)] {
			index := s.Pop().(int)
			res[index][2] = arr[i]
			if !s.IsEmpty() {
				res[index][1] = arr[s.Top().(int)]
			} else {
				res[index][1] = -1
			}
		}
		s.Push(i)
	}
	for !s.IsEmpty() {
		index := s.Pop().(int)
		res[index][2] = -1
		if !s.IsEmpty() {
			res[index][1] = arr[s.Top().(int)]
		} else {
			res[index][1] = -1
		}
	}
	return res
}

func maxTree(arr []int) *Node {
	head := &Node{0, nil, nil}
	n := len(arr)
	if n == 0 {
		return nil
	}
	s := stack.NewStack()
	lMap := make(map[*Node]*Node)
	rMap := make(map[*Node]*Node)
	nodeArray := make([]*Node, n)
	for i := 0; i < n; i++ {
		curNode := &Node{arr[i], nil, nil}
		nodeArray[i] = curNode
	}
	for i := 0; i < n; i++ {
		for !s.IsEmpty() && arr[i] > s.Top().(*Node).value {
			m := s.Pop().(*Node)
			rMap[m] = nodeArray[i]
			if !s.IsEmpty() {
				lMap[m] = s.Top().(*Node)
				getParent(lMap[m], m, rMap[m])
			} else {
				getParent(nil, m, rMap[m])
			}
		}
		s.Push(nodeArray[i])
	}
	for !s.IsEmpty() {
		m := s.Pop().(*Node)
		if !s.IsEmpty() {
			lMap[m] = s.Top().(*Node)
			getParent(lMap[m], m, nil)
		} else {
			getParent(nil, m, nil)
		}
	}
	//需要一个nodeArray数组来接收形成树后的节点们，否则每次形成的节点都被丢掉了
	for i := 0; i < n; i++ {
		curNode := nodeArray[i]
		if lMap[curNode] == nil && rMap[curNode] == nil {
			head = curNode
		}
	}
	return head
}

func getParent(l, m, r *Node) {
	if l == nil && r == nil {
		return
	} else if l == nil {
		if r.left == nil {
			r.left = m
		} else {
			r.right = m
		}
	} else if r == nil {
		if l.left == nil {
			l.left = m
		} else {
			l.right = m
		}
	} else {
		if l.value > r.value {
			if l.left == nil {
				l.left = m
			} else {
				l.right = m
			}
		} else {
			if r.left == nil {
				r.left = m
			} else {
				r.right = m
			}
		}
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
	arr := []int{3, 4, 2, 5, 6, 0, 1, 7}
	res := getNearMax(arr)
	fmt.Println(res)
	head := maxTree(arr)
	head.Print()
}
