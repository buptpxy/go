package main

/*
Q8. 栈
	1. 创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数——将数据放入栈，和 pop 函数——从栈中取得内容。
		栈应当是后进先出（LIFO）的。
	2. 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的方式打印整个栈： fmt.Printf("My stack %v\n", stack)
		栈可以被输出成这样的形式： [0:m] [1:l] [2:k]
*/
import (
	"fmt"
	"stack"
	// "strconv"
)

// type stack struct {
// 	i   int //元素即将被存入的下标位置
// 	arr [10]int
// }

// func (s *stack) push(d int) {
// 	if s.i > 9 {
// 		return
// 	}
// 	s.arr[s.i] = d
// 	s.i++
// }
// func (s *stack) pop() (d int) {
// 	s.i--
// 	d = arr[s.i]
// 	return
// }
// func (s *stack) String() (str string) {

// 	for n := 0; n < s.i; n++ {
// 		str = str + "[" + strconv.Itoa(n) + ":" + strconv.Itoa(s.arr[n]) + "] "
// 	}
// 	return
// }
func main() {
	p := new(stack.Stack)
	p.String()
	fmt.Printf("My new stack %v\n", p)

	p.Push(3)
	p.Push(2)
	p.Push(7)
	p.String()
	fmt.Printf("My push stack %v\n", p)
	fmt.Printf("My pop %d\n", p.Pop())

	p.String()
	fmt.Printf("My pop stack %v\n", p)
}
