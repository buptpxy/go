/*
【要求】给你一个栈， 请你逆序这个栈， 不能申请额外的数据结构， 只能使用递归函数。 如何实现？
递归函数本来就是系统栈。输出是逆序的，故可以先让栈中的元素递归出栈，再递归进栈
*/
package main

import (
	"fmt"
	"stack"
)

//依次让栈内元素出栈并在下一个元素出栈后入栈
func getAndRemoveLastElement(s *stack.Stack) int {
	result := s.Pop().(int)
	if s.IsEmpty() {
		return result
	}
	last := getAndRemoveLastElement(s)
	s.Push(result)
	return last
}

func reverseStack(s *stack.Stack) {
	if s.IsEmpty() {
		return
	}
	result := getAndRemoveLastElement(s)
	reverseStack(s)
	fmt.Println(result) //打印的顺序为4,3,2,1，可以看出系统把每次得到的result先暂存在系统栈中，执行getAndRemoveLastElement再出栈
	s.Push(result)
}
func main() {
	s := stack.NewStack()
	s.Push(1, 2, 3, 4)
	reverseStack(s)
	s.Print()
}
