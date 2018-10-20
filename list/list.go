package list

/*
单链表也可以没有头结点,但一定有头指针。无论链表是否为空，头指针均不为空。

头指针具有标识作用，故常用头指针冠以链表的名字。
头指针是指链表指向第一个结点的指针，若链表有头结点，则头指针就是指向链表头结点的指针。

头结点是为了操作的统一与方便而设立的，放在第一个元素结点之前，其数据域一般无意义（当然有些情况下也可存放链表的长度、用做监视哨等等）。
有了头结点后，对在第一个元素结点前插入结点和删除第一个结点，其操作与对其它结点的操作统一了。

首元结点也就是第一个元素的结点，它是头结点后边的第一个结点。
*/
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

//单链表一般使用头插法。在链表的第pos个位置前插入元素，首元结点为0位置
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
