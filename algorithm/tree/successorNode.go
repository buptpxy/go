/*
在二叉树中找到一个节点的后继节点
【 题目】 现在有一种新的二叉树节点类型如下：
public class Node {
	public int value; public Node left;
	public Node right; public Node parent;
	public Node(int data) { this.value = data; }
}
该结构比普通二叉树节点结构多了一个指向父节点的parent指针。
假设有一棵Node类型的节点组成的二叉树,树中每个节点的parent指针都正确地指向自己的父节点，头节点的parent指向nil。
只给一个在二叉树中的某个节点node，请实现返回node的后继节点的函数。
在二叉树的中序遍历的序列中， node的下一个节点叫作node的后继节点。
【思路】
	一个节点如果有右孩子，则它的后继节点为右孩子的最左边界；
	如果没有右孩子，则一直往上追溯，直到某个父节点p为p的父节点的左孩子，则后继节点为p的父节点
*/
package main

import (
	"fmt"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func findSuccessor(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.right == nil {
		h := head
		p := head.parent
		for p != nil {
			if p.left == h {
				return p
			}
			h = p
			p = p.parent
		}
		return nil
	}
	n := head.right
	for n.left != nil {
		n = n.left
	}
	return n
}

func main() {
	head := &Node{0, nil, nil, nil}
	head.left = &Node{1, nil, nil, head}
	head.right = &Node{2, nil, nil, head}
	head.left.left = &Node{3, nil, nil, head.left}
	head.left.right = &Node{4, nil, nil, head.left}
	head.right.left = &Node{5, nil, nil, head.right}
	head.right.right = &Node{6, nil, nil, head.right}
	fmt.Println(findSuccessor(head.right.right))
}
