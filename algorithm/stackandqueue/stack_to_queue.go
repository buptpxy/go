/*
如何仅用栈结构实现队列结构？
思路：
	使用两个栈，一个栈用来push，一个栈用来pop
	入队时，直接进入push栈
	出队时，如果pop栈没有元素，则push栈里面的元素全部出栈，进入pop栈，pop栈出栈；pop栈有元素则直接从pop栈出栈
*/
package main

import (
	"array/stack"
	"fmt"
)

var pushstack stack.ArrayStack
var popstack stack.ArrayStack

func push(num int) {
	pushstack.Push(num)
}

func pop() int {
	if pushstack.Size == 0 && popstack.Size == 0 {
		fmt.Println("queue is empty!")
		return 0
	}
	if popstack.Size == 0 {
		for pushstack.Size > 0 {
			popstack.Push(pushstack.Pop())
		}
	}
	top := popstack.Pop()
	fmt.Println("top: ", top)
	return top
}
func peek() int {
	if pushstack.Size == 0 && popstack.Size == 0 {
		fmt.Println("queue is empty!")
		return 0
	}
	if popstack.Size == 0 {
		for pushstack.Size > 0 {
			popstack.Push(pushstack.Pop())
		}
	}
	top := popstack.Peek()
	fmt.Println("top: ", top)
	return top
}
func main() {
	pushstack.InitStack(4)
	popstack.InitStack(4)
	push(1)
	push(2)
	push(3)
	pop()
	peek()
	fmt.Println("pushstack:")
	pushstack.Printstack()
	fmt.Println("popstack:")
	popstack.Printstack()
}
