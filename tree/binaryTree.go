package tree

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
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
	val := strconv.Itoa(head.Value)
	lenM := len(val)
	lenL := (lenT - lenM) / 2
	lenR := lenT - lenM - lenL
	printInOrder(head.Right, height+1, "v", 10)
	val = getSpace(height*lenT) + getSpace(lenL) + to + val + to + getSpace(lenR)
	fmt.Println(val)
	printInOrder(head.Left, height+1, "^", 10)
}

func (head *Node) Print() {
	fmt.Println("BinaryTree:")
	printInOrder(head, 0, "", 10)
}
