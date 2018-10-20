/*
输入：参数1，正数数组costs 参数2，正数数组profits 参数3，正数k 参数4， 正数m
		costs[i]表示i号项目的花费
		profits[i]表示i号项目在扣除花费之后还能挣到的钱(利润)
		k表示你不能并行、 只能串行的最多做k个项目
		m表示你初始的资金
说明： 你每做完一个项目，马上获得的收益，可以支持你去做下一个项目。
输出： 你最后获得的最大钱数。
【思路】利用一个小根堆和一个大根堆，先把所有的项目都根据花费放入小根堆，然后依次弹出可以做的项目，根据收益放入大根堆
		弹出大根堆堆顶的同时更新总资金m，再从小根堆中弹出可以做的项目放入大根堆。直到做完k个项目，或大根堆为空，或无项目花费低于总资金
*/
package main

import (
	"container/heap"
	"fmt"
)

type project struct {
	cost   int
	profit int
}
type minCost []project
type maxProfit []project

//小根堆
func (p minCost) Len() int {
	return len(p)
}
func (p minCost) Less(i, j int) bool {
	return p[i].cost < p[j].cost
}
func (p minCost) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p *minCost) Push(x interface{}) {
	*p = append(*p, x.(project))
}
func (p *minCost) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

//大根堆
func (p maxProfit) Len() int {
	return len(p)
}
func (p maxProfit) Less(i, j int) bool {
	return p[i].profit > p[j].profit
}
func (p maxProfit) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p *maxProfit) Push(x interface{}) {
	*p = append(*p, x.(project))
}
func (p *maxProfit) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
func getProject(costs []int, profits []int) []project {
	n := len(costs)
	p := make([]project, n)
	for i := 0; i < n; i++ {
		p[i].cost = costs[i]
		p[i].profit = profits[i]
	}
	return p
}
func minhToMaxh(minheap *minCost, maxheap *maxProfit, m int) {
	for i := 0; i < minheap.Len(); i++ {
		p := heap.Pop(minheap).(project)
		if p.cost <= m {
			heap.Push(maxheap, p)
		} else {
			heap.Push(minheap, p) //注意要把未放入大根堆的元素重新放回小根堆
		}
	}
}
func maxMoney(costs []int, profits []int, k int, m int) int {
	minheap := &minCost{}
	heap.Init(minheap)
	maxheap := &maxProfit{}
	heap.Init(maxheap)
	projects := getProject(costs, profits)
	for i := 0; i < len(projects); i++ {
		heap.Push(minheap, projects[i])
	}
	minhToMaxh(minheap, maxheap, m)
	for maxheap.Len() > 0 && k > 0 {
		p := heap.Pop(maxheap).(project)
		m = m + p.profit
		k--
		minhToMaxh(minheap, maxheap, m)
	}
	return m
}
func main() {
	costs := []int{3, 5, 7}
	profits := []int{1, 2, 3}
	// fmt.Println(getProject(costs, profits))
	fmt.Println(maxMoney(costs, profits, 3, 4))
}
