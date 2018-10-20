/*
如何仅用队列结构实现栈结构？
思路：
	使用两个队列，一个data队列和一个help队列，要进栈的元素直接进data,
	要出栈时，除了队尾的元素，其他元素全部出data进help,然后弹出队尾元素
	把help和data互换
*/
package main

import (
	"array/queue"
	"fmt"
)

var data queue.ArrayQueue
var help queue.ArrayQueue

func push(num int) {
	data.Push(num)
}
func pop() int {
	if data.Size == 0 && help.Size == 0 {
		fmt.Println("stack is empty!")
		return 0
	}
	if help.Size == 0 {
		for data.Size > 1 {
			help.Push(data.Pop())
		}
	}
	top := data.Pop()
	swap()
	fmt.Println(top)
	return top
}

func peek() int {
	if data.Size == 0 && help.Size == 0 {
		fmt.Println("stack is empty!")
		return 0
	}
	if help.Size == 0 {
		for data.Size > 1 {
			help.Push(data.Pop())
		}
	}
	top := data.Pop()
	help.Push(top)
	swap()
	fmt.Println(top)
	return top
}
func swap() {
	data, help = help, data
}
func printStack() {
	//不能用data.Size作为约束条件，因为一直在改变
	for i := 0; i < 3; i++ {
		pop()
	}
}
func main() {
	data.InitQueue(4)
	help.InitQueue(4)
	push(1)
	push(2)
	push(3)
	printStack()
	// pop()
	// pop()
	// fmt.Println("data:")
	// data.Printqueue()
	// fmt.Println("help")
	// help.Printqueue()

}
