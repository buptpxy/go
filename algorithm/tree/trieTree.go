/*
介绍前缀树
何为前缀树？ 如何生成前缀树？
例子：
一个字符串类型的数组arr1， 另一个字符串类型的数组arr2。
arr2中有哪些字符， 是arr1中出现的？ 请打印
arr2中有哪些字符， 是作为arr1中某个字符串前缀出现的？ 请打印
arr2中有哪些字符， 是作为arr1中某个字符串前缀出现的？ 请打印arr2中出现次数最大的前缀。
*/
package main

import (
	"fmt"
)

type Node struct {
	path  int //这个字符被加入的次数
	end   int //以这个字符串结尾的次数
	nexts []*Node
}

func insert(head *Node, value string) {
	vals := []rune(value)
	index := 0
	for _, v := range vals {
		index = int(v - 'a')
		if head.nexts[index] == nil {
			head.nexts[index] = &Node{0, 0, make([]*Node, 26)}
		}
		head = head.nexts[index]
		head.path++
	}
	head.end++
}
func search(head *Node, word string) int {
	words := []rune(word)
	index := 0
	for _, v := range words {
		index = int(v - 'a')
		if head.nexts[index] == nil {
			return 0
		}
		head = head.nexts[index]
	}
	return head.end
}
func searchPrefix(head *Node, word string) int {
	words := []rune(word)
	index := 0
	for _, v := range words {
		index = int(v - 'a')
		if head.nexts[index] == nil {
			return 0
		}
		head = head.nexts[index]
	}
	return head.path
}

func delete(head *Node, word string) {
	if search(head, word) == 0 {
		return
	}
	words := []rune(word)
	index := 0
	for _, v := range words {
		index = int(v - 'a')
		head.nexts[index].path--
		if head.nexts[index].path == 0 {
			head.nexts[index] = nil
			return
		} else {
			head = head.nexts[index]
		}
	}
	head.end--
}

//不论是值类型还是引用类型，在函数里都会重新创建一个副本
func test(p *Node) {
	p.nexts[0] = &Node{1, 0, make([]*Node, 26)}
	// p = p.nexts[0] //这种赋值只是把p.nexts[0]的地址给了副本p的地址，在函数外部p指向的内容不会变成p.nexts[0]指向的内容
	*p = *(p.nexts[0]) //把p.nexts[0]的内容赋值给副本p的内容，在函数外部p指向的内容才会变成p.nexts[0]指向的内容
}
func main() {
	head := &Node{0, 0, make([]*Node, 26)}
	// test(head)
	insert(head, "dog")
	insert(head, "dodg")
	insert(head, "snack")
	insert(head, "snack")
	insert(head, "mouse")
	delete(head, "dog")
	fmt.Println(head)
	// fmt.Println(search(head, "snack"))
	// fmt.Println(searchPrefix(head, "snac"))
	fmt.Println(head.nexts['d'-'a'])
	fmt.Println(head.nexts['d'-'a'].nexts['o'-'a'])
	// fmt.Println(head.nexts['d'-'a'].nexts['o'-'a'].nexts['g'-'a'])
}
