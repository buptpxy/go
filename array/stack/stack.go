package stack

import (
	"fmt"
)

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
