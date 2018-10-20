package main

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node //Next，List都是指针
	List  *List //List是指向Root的指针，即头指针，也代表这个链表
}
type List struct {
	Root Node //root不是指针，是链表的头结点
	Len  int
}

//初始化方式一 ：var l List; l.Init()
func (l *List) Init() {
	l.Root.Next = &l.Root //头结点的Next暂时指向自己
	l.Len = 0
}

//初始化方式二：l:= New()
func New() *List {
	var root Node
	root.Next = &root
	return &List{root, 0}
}

//在链表的第pos个位置前插入元素，首元结点为0位置
func (l *List) Insert(pos int, v interface{}) {
	var n Node
	n.Value = v
	if pos == 0 {
		n.Next = l.Root.Next
		l.Root.Next = &n
		n.List = l
		l.Len++
	} else if pos > 0 && pos < l.Len {
		pre := &(l.Root)
		for i := 0; i < pos; i++ {
			pre = pre.Next
		}
		n.Next = pre.Next
		pre.Next = &n
		n.List = l
		l.Len++
	} else {
		fmt.Println("Please give a pos between 0 to l.len()-1!")
		return
	}
}

//获得链表在第pos个位置上的元素
func (l *List) Get(pos int) interface{} {
	if pos >= 0 && pos < l.Len {
		pre := &(l.Root)
		for i := 0; i <= pos; i++ {
			pre = pre.Next
		}
		return (*pre).Value
	} else {
		fmt.Println("Please give a pos between 0 to l.len()-1!")
		return nil
	}
}

//删除链表在第pos个位置的元素
func (l *List) Remove(pos int) interface{} {
	if pos >= 0 && pos < l.Len {
		pre := &(l.Root)
		var n *Node
		if pos == 0 {
			n = pre.Next
		} else {
			for i := 0; i < pos; i++ {
				pre = pre.Next
			}
			n = pre.Next
		}

		pre.Next = n.Next
		n.Next = nil
		n.List = nil
		l.Len--
		return (*n).Value
	} else {
		fmt.Println("Please give a pos between 0 to l.len()-1!")
		return nil
	}
}

//打印链表里的所有元素
func (l *List) Print() {
	p := l.Root.Next
	fmt.Printf("root->")
	for i := 0; i < l.Len; i++ {
		fmt.Printf("%v->", (*p).Value)
		p = p.Next
	}
	fmt.Printf("nil\n")
}
func main() {
	var l List
	l.Init()
	l.Insert(0, "a")
	l.Insert(0, "b")
	fmt.Println(l.Get(0))
	//fmt.Println(l.Remove(0))
	l.Print()
}
