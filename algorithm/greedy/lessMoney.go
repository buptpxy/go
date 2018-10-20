/*
【题目】一块金条切成两半，是需要花费和长度数值一样的铜板的。 比如长度为20的金条， 不管切成长度多大的两半， 都要花费20个铜板。
一群人想整分整块金条， 怎么分最省铜板？
例如,给定数组{10,20,30}，代表一共三个人，整块金条长度为10+20+30=60. 金条要分成10,20,30三个部分。
如果，先把长度60的金条分成10和50,花费60;再把长度50的金条分成20和30,花费50 ;一共花费110铜板。
但是如果， 先把长度60的金条分成30和30，花费60;再把长度30金条分成10和20,花费30; 一共花费90铜板。
输入一个数组， 返回分割的最小代价。
【思路】贪心算法常用堆来解决。
	此题即一个哈夫曼编码的问题，即把数组元素都放入一个堆中。
	每次弹出最小的两个，再把这两个的和放回去。
	一直到堆中只剩下最后一个元素，即为最小代价。
*/
package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func huffman(h *IntHeap, s []int) int {
	for _, v := range s {
		heap.Push(h, v)
	}
	for h.Len() > 1 {
		m1 := heap.Pop(h).(int)
		fmt.Printf("m1: %d\n", m1)
		m2 := heap.Pop(h).(int)
		fmt.Printf("m2: %d\n", m2)
		m1 = m1 + m2
		heap.Push(h, m1)
	}
	return heap.Pop(h).(int)
}

func main() {
	h := &IntHeap{}
	heap.Init(h)
	s := []int{10, 20, 30}
	fmt.Println(huffman(h, s))
	// fmt.Println(heap.Pop(h).(int))
	// fmt.Println(heap.Pop(h).(int))
}
