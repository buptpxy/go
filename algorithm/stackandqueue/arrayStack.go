/*
用数组结构实现大小固定的栈
实现栈的思路：
	申请一个固定大小的slice，
	给栈一个记录大小的size变量，判断栈空或满
	给栈一个记录栈顶元素位置的index变量，若index超过了数组的length-1则回到0继续循环
	元素从栈顶进从栈顶出
*/

package main

type ArrayStack struct {
	Arr   []int
	Size  int
	Index int
}

func (as *ArrayStack) InitStack(length int) {
	as.Arr = make([]int, length)
	as.Size = 0
	as.Index = 0
	// as.Printstack()
}

func (as *ArrayStack) Push(num int) {
	length := len(as.Arr)
	if as.Size == length {
		fmt.Println("stack is full!")
		return
	}
	as.Arr[as.Index] = num
	as.Size++
	if as.Index == length-1 {
		as.Index = 0
	} else {
		as.Index++
	}
	// as.Printstack()
}

func (as *ArrayStack) Pop() int {
	length := len(as.Arr)
	if as.Size == 0 {
		fmt.Println("stack is empty!")
		return 0
	}
	if as.Index == 0 {
		as.Index = length - 1
	} else {
		as.Index--
	}
	top := as.Arr[as.Index]
	as.Size--
	return top
}

func (as *ArrayStack) Peek() int {
	length := len(as.Arr)
	if as.Size == 0 {
		fmt.Println("stack is empty!")
		return 0
	}
	var top int
	if as.Index == 0 {
		top = as.Arr[length-1]
	} else {
		top = as.Arr[as.Index-1]
	}
	return top
}
func (as *ArrayStack) Printstack() {
	for k, v := range as.Arr {
		fmt.Printf("%d => %d\n", k, v)
	}
	fmt.Printf("------------Index:%d Size:%d top:%d-------------\n", as.Index, as.Size, as.Peek())
}

func main() {
	var as ArrayStack //不能是var as *ArrayStack
	as.InitStack(3)
	as.Push(1)
	as.Push(4)
	as.Push(8)
	as.Pop()
	as.Push(5)
}
