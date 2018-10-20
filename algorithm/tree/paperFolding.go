/*
折纸问题
【 题目】 请把一段纸条竖着放在桌子上， 然后从纸条的下边向上方对折1次， 压出折痕后展开。
		此时折痕是凹下去的， 即折痕突起的方向指向纸条的背面。
		如果从纸条的下边向上方连续对折2次，压出折痕后展开，此时有三条折痕，从上到下依次是下折痕、下折痕和上折痕。
		给定一个输入参数N， 代表纸条都从下边向上方连续对折N次，请从上到下打印所有折痕的方向。
	例如： N=1时， 打印： down
		  N=2时， 打印： down down up
【思路一】折纸的过程可看成一颗满二叉树的生成，折几次树的高度就为几，根节点值为down，所有节点的左孩子为down，右孩子为up。
然后把树的中序遍历打印出来

*/
package main

import (
	"fmt"
)

type Node struct {
	value string
	left  *Node
	right *Node
}

func generateTree(head *Node, value string, height int) *Node {
	if height == 0 {
		return head
	}
	//由于新生成的节点一开始都为空，所以必须给节点初始化
	if head == nil {
		head = &Node{value, nil, nil}
	}
	head.left = generateTree(head.left, "down", height-1)
	fmt.Printf("%s ", head.value)
	head.right = generateTree(head.right, "up", height-1)
	return head
}

func getSpace(num int) string {
	var s string
	for i := 0; i < num; i++ {
		s = s + " "
	}
	return s
}
func printTree(head *Node, height int, to string, lenT int) {
	if head == nil {
		return
	}
	val := head.value
	lenM := len(val)
	lenL := (lenT - lenM) / 2
	lenR := lenT - lenM - lenL
	printTree(head.right, height+1, "v", 10)
	val = getSpace(height*lenT) + getSpace(lenL) + to + val + to + getSpace(lenR)
	fmt.Println(val)
	printTree(head.left, height+1, "^", 10)
}

func main() {
	head := &Node{"down", nil, nil}
	generateTree(head, "down", 2)
	// printTree(head, 0, "", 10)
	fmt.Println()
}
