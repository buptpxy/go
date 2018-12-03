/*给定一棵二叉树的头节点head， 判断这棵树是不是平衡二叉树
【思路】判断每个节点为根的数是否是平衡二叉树，只要有一个节点不是平衡二叉树，这整棵树就不是
1、可能性：1）左子树是或不是平衡二叉树 2）右子树是或不是平衡二叉树 3）左右子树都是，当前树是；左右子树只要有一个不是，当前树肯定不是
2、每个节点要收集的信息：是否平衡、树高h
3、basecase:head==nil时，是平衡，h=0
*/
package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
}

type res struct {
	IsBS   bool
	height int
}

func MoreThanOneGap(i, j int) bool {
	if (i > j && i-j > 1) || (i < j && j-i > 1) {
		return true
	}
	return false
}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//O(2n-1) all nodes have been searched twice except root
func IsBalancedTree(head *Node) res {
	fmt.Println("rec")
	if head == nil {
		return res{true, 0}
	}
	leftRes := IsBalancedTree(head.left)
	rightRes := IsBalancedTree(head.right)
	if !leftRes.IsBS {
		return res{false, 0}
	}
	if !rightRes.IsBS {
		return res{false, 0}
	}
	h := 0
	if MoreThanOneGap(leftRes.height, rightRes.height) {
		return res{false, 0}
	}
	h = max(leftRes.height, rightRes.height) + 1 //对递归条件的接收不要放在if条件语句中，尽量放在外面
	return res{true, h}
}

func main() {
	head1 := &Node{nil, nil}
	head1.left = &Node{nil, nil}
	head1.right = &Node{nil, nil}
	head1.left.left = &Node{nil, nil}
	head1.left.right = &Node{nil, nil}
	head1.right.left = &Node{nil, nil}
	head1.right.right = &Node{nil, nil}
	head1.right.right.right = &Node{nil, nil}

	head2 := &Node{nil, nil}
	head2.left = &Node{nil, nil}
	head2.right = &Node{nil, nil}
	head2.right.left = &Node{nil, nil}
	head2.right.right = &Node{nil, nil}
	head2.right.left.left = &Node{nil, nil}
	head2.right.left.left.left = &Node{nil, nil}
	head2.right.right.right = &Node{nil, nil}
	head2.right.right.right.right = &Node{nil, nil}

	res := IsBalancedTree(head2)
	fmt.Println(res)

}
