/*
实现一个特殊的栈， 在实现栈的基本功能的基础上， 再实现返
回栈中最小元素的操作。
【 要求】
1． pop、 push、 getMin操作的时间复杂度都是O(1)。
2． 设计的栈类型可以使用现成的栈结构。
【 思路】
	使用两个栈，一个栈记录数据，一个栈同时记录栈中所有元素的最小值
*/

package main

import (
	"array/stack"
	"fmt"
)

type minStack struct {
	data stack.ArrayStack
	mins stack.ArrayStack
}

func (ms *minStack) initStack(length int) {
	ms.data.InitStack(length)
	ms.mins.InitStack(length)
}

func (ms *minStack) push(num int) {
	ms.data.Push(num)
	if ms.mins.Size == 0 || num < ms.mins.Peek() {
		ms.mins.Push(num)
	} else {
		ms.mins.Push(ms.mins.Peek())
	}
}
func (ms *minStack) pop() int {
	top := ms.data.Pop()
	ms.mins.Pop()
	fmt.Println("top: ", top)
	return top
}
func (ms *minStack) getMin() int {
	min := ms.mins.Peek()
	fmt.Println("min:", min)
	return min
}
func (ms *minStack) printStack() {
	fmt.Println("data stack:")
	ms.data.Printstack()
	fmt.Println("mins stack:")
	ms.mins.Printstack()
}
func main() {
	var ms minStack
	ms.initStack(3)
	ms.push(3)
	ms.push(2)
	ms.push(1)
	ms.printStack()
	ms.pop()
	ms.printStack()
	ms.getMin()
}
