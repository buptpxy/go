/*
介绍二叉树的序列化和反序列化
【思路】要用分隔符把每个节点隔开，空节点也要打印出来
*/

package main

import (
	"fmt"
	"queue"
	"strconv"
	"strings"
	"tree"
)

func serialByPre(head *tree.Node) string {
	if head == nil {
		return "# "
	}
	s := strconv.Itoa(head.Value) + " "
	s = s + serialByPre(head.Left)
	s = s + serialByPre(head.Right)
	return s
}
func stringToQueue(s string) *queue.Queue {
	s1 := strings.Fields(s)
	q := queue.NewQueue()
	for _, v := range s1 {
		q.Push(v)
	}
	return q
}

func reconstructByPre(q *queue.Queue) *tree.Node {
	v := q.Pop().(string)
	if v == "#" {
		return nil
	}
	i, _ := strconv.Atoi(v)
	head := &tree.Node{i, nil, nil}
	head.Left = reconstructByPre(q)
	head.Right = reconstructByPre(q)
	return head
}
func serialByLevel(head *tree.Node) string { //注意"#"打印的位置
	if head == nil {
		return "# "
	}
	q := queue.NewQueue()
	n := head
	q.Push(n)
	s := ""
	for !q.IsEmpty() {
		n = q.Pop().(*tree.Node)
		if n == nil {
			s = s + "# "
		} else {
			s = s + strconv.Itoa(n.Value) + " "
			q.Push(n.Left)
			q.Push(n.Right)
		}
	}
	return s
}
func reconstructByLevel(s string) *tree.Node {
	q := queue.NewQueue()
	ss := strings.Fields(s)
	index := 0
	head := generateNodeByString(ss[index])
	index++
	n := head
	if head != nil {
		q.Push(n)
	}
	for !q.IsEmpty() {
		n = q.Pop().(*tree.Node)
		n.Left = generateNodeByString(ss[index])
		index++
		n.Right = generateNodeByString(ss[index])
		index++
		if n.Left != nil {
			q.Push(n.Left)
		}
		if n.Right != nil {
			q.Push(n.Right)
		}
	}
	return head

}
func generateNodeByString(s string) *tree.Node {
	if s == "#" {
		return nil
	}
	i, _ := strconv.Atoi(s)
	return &tree.Node{i, nil, nil}
}
func main() {
	head := &tree.Node{0, nil, nil}
	head.Left = &tree.Node{1, nil, nil}
	head.Right = &tree.Node{2, nil, nil}
	head.Left.Left = &tree.Node{3, nil, nil}
	head.Left.Right = &tree.Node{4, nil, nil}
	head.Right.Left = &tree.Node{5, nil, nil}
	head.Right.Right = &tree.Node{6, nil, nil}
	s1 := serialByPre(head)
	fmt.Println("serialize by preOrder: ", s1)
	q := stringToQueue(s1)
	head1 := reconstructByPre(q)
	head1.Print()
	s2 := serialByLevel(head)
	fmt.Println("serialize by level: ", s2)
	head2 := reconstructByLevel(s2)
	head2.Print()
}
