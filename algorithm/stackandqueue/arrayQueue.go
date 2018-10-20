/*
用数组结构实现大小固定的队列
实现队列思路：
	申请一个固定大小的slice
	给队列一个记录大小的Size变量，判断队列空或满
	给队列一个记录队头元素位置的Start变量和一个记录队尾元素位置的End变量，
	若Start和End超过了数组的length-1则回到0继续循环
	元素从队尾进，从队头出
*/
package main

import (
	"fmt"
)

type ArrayQueue struct {
	Arr   []int
	Size  int
	Start int
	End   int
}

func (aq *ArrayQueue) InitQueue(length int) {
	aq.Arr = make([]int, length)
	aq.Size = 0
	aq.Start = 0
	aq.End = 0
	aq.Printqueue()
}

func (aq *ArrayQueue) Push(num int) {
	length := len(aq.Arr)
	if aq.Size == length {
		fmt.Println("queue is full!")
		return
	}
	aq.Arr[aq.End] = num
	aq.Size++
	if aq.End == length-1 {
		aq.End = 0
	} else {
		aq.End++
	}
	aq.Printqueue()
}

func (aq *ArrayQueue) Pop() int {
	length := len(aq.Arr)
	if aq.Size == 0 {
		fmt.Println("queue is empty!")
		return 0
	}
	End := aq.Arr[aq.Start]
	aq.Size--
	if aq.Start == length-1 {
		aq.Start = 0
	} else {
		aq.Start++
	}
	return End
}

func (aq *ArrayQueue) Peek() int {

	if aq.Size == 0 {
		fmt.Println("queue is empty!")
		return 0
	}
	End := aq.Arr[aq.Start]
	return End
}

func (aq *ArrayQueue) Printqueue() {
	for k, v := range aq.Arr {
		fmt.Printf("%d => %d\n", k, v)
	}
	fmt.Printf("------------Start:%d End:%d Size:%d top:%d-------------\n", aq.Start, aq.End, aq.Size, aq.Peek())
}

func main() {
	var aq ArrayQueue
	aq.InitQueue(3)
	aq.Push(2)
	aq.Push(3)
	aq.Pop()
	aq.Push(5)
	aq.Push(6)
	aq.Push(7)
}
