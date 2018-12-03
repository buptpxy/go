/*
Morris遍历
利用Morris遍历实现二叉树的先序， 中序， 后续遍历， 时间复杂度O(N)， 额外空间复杂度O(1)。
步骤：
1. 一个指向当前节点的指针为cur，如果cur.left为空，则cur=cur.right
2. 如果cur.left不为空，分两种情况 1)cur.left的最右节点mostRight.right为空，则让mostRight.right=cur,cur=cur.left
								2)cur.left的最右节点mostRight.right=cur，则mostRight.right=nil,cur=cur.right
3. 如果cur.left和cur.right都为空，则遍历结束
4. 树中有左子树的节点都被遍历了2次，没有左子树的节点都被遍历了1次
5. 在节点第一次被遍历时就打印，结果就是先序遍历
6. 在节点第二次被遍历时打印，就是中序遍历
7. 在节点第一次被遍历时，打印它的左子树的最右节点，最后再逆序打印整棵树的右边界，结果就是后序遍历。逆序打印整棵树右边界的方法是先把指针逆置，打印完再逆回来。
*/
package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func morrisPre(head *Node) {
	if head == nil {
		return
	}
	cur := head
	mostRight := head
	for cur != nil {
		if cur.left != nil {
			mostRight = cur.left
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				fmt.Printf("%d ", cur.value) //第一次遍历到cur，先序遍历
				cur = cur.left
			} else {
				mostRight.right = nil
				cur = cur.right
			}
		} else {
			//当一开始左子树为空时，此节点只遍历一次
			fmt.Printf("%d ", cur.value) //遍历到cur，先序遍历
			cur = cur.right
		}
	}
	fmt.Println()
}

func morrisIn(head *Node) {
	if head == nil {
		return
	}
	cur := head
	mostRight := head
	for cur != nil {
		mostRight = cur.left
		if cur.left != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				//求最右边界
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
				fmt.Printf("%d ", cur.value) //第二次遍历到cur，中序遍历
				cur = cur.right
			}
		} else {
			//当一开始左子树为空时，此节点只遍历一次
			fmt.Printf("%d ", cur.value) //遍历到cur，中序遍历
			cur = cur.right
		}
	}
	fmt.Println()
}
func morrisIn1(head *Node) {
	if head == nil {
		return
	}
	cur := head
	mostRight := head
	for cur != nil {
		mostRight = cur.left
		if cur.left != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				//求最右边界
				cur = cur.left
				continue
			} else {
				mostRight.right = nil //此处的cur=cur.right和cur.left=nil的情况中的cur=cur.right一起合并到了最后
			}
		}
		fmt.Printf("%d ", cur.value) //第二次遍历到cur，中序遍历
		cur = cur.right
	}
	fmt.Println()
}
func reverseEdge(from *Node) *Node {
	if from == nil || from.right == nil {
		return from
	}
	var pre *Node
	for from != nil {
		next := from.right
		from.right = pre
		pre = from
		from = next
	}
	return pre
}
func printEdge(head *Node) {
	tail := reverseEdge(head)
	cur := tail
	for cur != nil {
		fmt.Printf("%d ", cur.value)
		cur = cur.right
	}
	head = reverseEdge(tail)
}

func morrisPos(head *Node) {
	if head == nil {
		return
	}
	cur := head
	for cur != nil {
		mostRight := cur.left
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
			} else {
				mostRight.right = nil
				printEdge(cur.left) //打印当前节点左子树的右边界
				cur = cur.right
			}
		} else {
			cur = cur.right
		}
	}
	printEdge(head) //打印整棵树的右边界
	fmt.Println()
}

func main() {
	head := &Node{0, nil, nil}
	head.left = &Node{1, nil, nil}
	head.right = &Node{2, nil, nil}
	head.left.left = &Node{3, nil, nil}
	head.left.right = &Node{4, nil, nil}
	head.right.left = &Node{5, nil, nil}
	head.right.right = &Node{6, nil, nil}
	morrisPre(head)
	morrisIn(head)
	morrisIn1(head)
	morrisPos(head)
}
