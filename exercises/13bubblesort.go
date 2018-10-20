package main

/*
Q13.  冒泡排序
	编写一个针对 int 类型的 slice 冒泡排序的函数.
	它在一个列表上重复步骤来排序，比较每个相䩪的元素，并且顺序错误的时候，交换它们。
	一遍一遍扫描列表，直到没有交换为止，这意味着列表排序完成。
	算法得名于更小的元素就像 “泡泡” 一样冒到列表的顶端。
*/
import (
	"fmt"
)

type Sorter interface {
	Len() int
	// Less(i int, j int) bool
	Swap(i int, j int)
}

type Si []int
type Ss []string

func (s Si) Len() int {
	return len(s)
}
func (s Ss) Len() int {
	return len(s)
}

func (s Si) Swap(i int, j int) {
	if s[j] < s[i] {
		s[j], s[i] = s[i], s[j]
	}
}
func (s Ss) Swap(i int, j int) {
	if s[j] < s[i] {
		s[j], s[i] = s[i], s[j]
	}
}

func Sort(s Sorter) {
	for i := 0; i < s.Len()-1; i++ {
		for j := i + 1; j < s.Len(); j++ {
			s.Swap(i, j)
		}
	}
}
func main() {
	intslice := Si{3, 7, 23, 0, 2, 5}
	stringslice := Ss{"a", "b", "sgf", "segd"}
	Sort(intslice)
	Sort(stringslice)
	fmt.Printf("intslice: %d \n", intslice)
	fmt.Printf("stringslice: %s \n", stringslice)
}
