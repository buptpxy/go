/*
一些项目要占用一个会议室宣讲， 会议室不能同时容纳两个项目的宣讲。
 给你每一个项目开始的时间和结束的时间(给你一个数组，里面是一个个具体的项目)，
 你来安排宣讲的日程，要求会议室进行的宣讲的场次最多。
 返回这个最多的宣讲场次。
 【思路】应安照每个项目的结束时间来排序项目，则可让宣讲最多的次数
*/
package main

import (
	"fmt"
	"sort"
)

type project struct {
	start int
	end   int
}
type projects []project

func (p projects) Len() int {
	return len(p)
}
func (p projects) Less(i, j int) bool {
	return p[i].end < p[j].end
}
func (p projects) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func mostTimes(p projects, cur int) int {
	sort.Sort(p)
	count := 0
	for i := 0; i < len(p); i++ {
		if p[i].start >= cur {
			count++
			cur = p[i].end
		}
	}
	return count
}
func main() {
	p := projects{
		{9, 10},
		{6, 9},
	}
	fmt.Println(mostTimes(p, 7))
}
