/*
要求： 给定两个二叉树T1和T2， 返回T1的某个子树结构是否与T2的结构相等。
思路： 分别写出两棵树的先续遍历（带空子树），然后判断T2是否是T1的子串。
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

func serialByPre(head *Node) string {
	if head == nil {
		return "#"
	}
	s := strconv.Itoa(head.value) + ""
	s = s + serialByPre(head.left)
	s = s + serialByPre(head.right)
	return s
}

func next(str string) []int {
	r := []rune(str)
	next := make([]int, len(r))
	next[0] = -1
	next[1] = 0
	for i := 2; i < len(r); i++ {
		j := i - 1
		for j > 0 {
			if r[i-1] == r[next[j]] {
				next[i] = next[i-1] + 1
				break
			}
			j = next[j]
		}
	}
	return next
}

func kmp(str1 string, str2 string) int {
	nex := next(str2)
	r1 := []rune(str1)
	r2 := []rune(str2)
	i := 0
	j := 0
	for i < len(r1) && j < len(r2) {
		if r1[i] == r2[j] {
			i++
			j++
		} else if j != 0 {
			j = nex[j]
		} else {
			i++
		}
	}
	if j == len(r2) {
		return i - j
	}
	return -1
}

func isSubTree(head1 *Node, head2 *Node) bool {
	str1 := serialByPre(head1)
	str2 := serialByPre(head2)
	if kmp(str1, str2) != -1 {
		return true
	}
	return false

}

func main() {
	head1 := &Node{1, nil, nil}
	head1.left = &Node{2, nil, nil}
	head1.right = &Node{3, nil, nil}
	head1.left.left = &Node{4, nil, nil}
	head1.left.right = &Node{5, nil, nil}
	head1.right.left = &Node{6, nil, nil}
	head1.right.right = &Node{7, nil, nil}
	head1.left.left.right = &Node{8, nil, nil}
	head1.left.right.left = &Node{9, nil, nil}

	head2 := &Node{2, nil, nil}
	head2.left = &Node{4, nil, nil}
	head2.right = &Node{5, nil, nil}
	head2.left.right = &Node{8, nil, nil}
	head2.right.left = &Node{9, nil, nil}
	fmt.Println(isSubTree(head1, head2))
}
