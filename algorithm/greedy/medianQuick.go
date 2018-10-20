/*
堆的应用：
	有一个源头一直在往外输出数字，要求实时求出当前所有已输出数字的中位数。
	中位数是一串数中大小排在中间的数，若共有奇数个数字，中位数只有一个；若共有偶数个数字，中位数定义为大小排在中间的两个数的平均值。
思路：
	建立一个大根堆一个小根堆，把较小的N/2个数放进大根堆，较大的N/2个数放进小根堆。
	这样就可保证size大的那个堆的根即为中位数，size一样大的话，即取两个堆的根的平均数
具体操作：
	1. 把第一个数放进大根堆，
	2. 如果接下来的数小则放进大根堆，否则放进小根堆。
	3. 当两个堆size差异大于1时，把size大的那个堆的堆顶元素取出放入另一个堆。
	4. 重复2，3直到n个元素放完
*/
package main

import (
	"container/heap"
	"fmt"
	// "math/rand"
	// "time"
)

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func findMedian(minh *minHeap, maxh *maxHeap, x int) int {
	if maxh.Len() == 0 {
		heap.Push(maxh, x)
	}
	median := heap.Pop(maxh).(int)

	if x < median {
		heap.Push(maxh, x)
		heap.Push(maxh, median)
	} else {
		heap.Push(minh, x)
		heap.Push(maxh, median)
	}
	if minh.Len()-maxh.Len() > 1 {
		p := heap.Pop(minh).(int)
		heap.Push(maxh, p)
	}
	if maxh.Len()-minh.Len() > 1 {
		p := heap.Pop(maxh).(int)
		heap.Push(minh, p)
	}
	if maxh.Len() == minh.Len() {
		median = (heap.Pop(maxh).(int) + heap.Pop(minh).(int)) / 2
	} else if maxh.Len() > minh.Len() {
		median = heap.Pop(maxh).(int)
	} else {
		median = heap.Pop(minh).(int)
	}
	return median
}
func main() {
	minh := &minHeap{4, 7, 9}
	maxh := &maxHeap{1, 2, 3}
	heap.Init(minh)
	heap.Init(maxh)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	fmt.Println("num: ", num)
	fmt.Println("median", findMedian(minh, maxh, num))

}
