//用heap包（小根堆）来排序自定义类型
package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	oldh := *h
	n := len(oldh)
	x := oldh[n-1]
	*h = oldh[0 : n-1]
	return x
}
func main() {
	h := &IntHeap{2, 1, 5}
	//把h生成一个小根堆
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("最小值为：%d\n", (*h)[0])
	//此处写为h.Len()也是对的
	for (*h).Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
