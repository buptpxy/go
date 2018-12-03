/*求最大子矩阵的大小
给定一个整型矩阵map，其中的值只有0和1两种，求其中全是1的所有矩形区域中，全部为1最大的矩形区域的1的数量。
例如：1 1 1 0
其中，最大的矩形区域有3个1，所以返回3
再如：
1 0 1 1
1 1 1 1
1 1 1 0
其中，最大的矩形区域有6个1，所以返回6
思路：
可利用从底到顶单调递增的单调栈，来得到每个元素左边和右边离他最近的比他小的元素。
与单调栈稍微不同的是，1、遇到相等的元素时下标小的元素直接出栈，下标大的元素进展，不需要再把两个下标压在一起.2、且遇到比栈顶元素小的元素时不用更新MaxAera，直接入栈。
一个数组{3,5,2,5,6,0,1,5}可以表示一个直方图，数组中的每个元素为直方图的高度。然后计算每一条直方可以向左边和右边扩的大小，显然，遇到比它小的就扩不动了，此时就可以计算它扩的区域的面积。
对于一个矩阵来说，可以分别计算以每一行为底的直方图的MaxAera，最后找出最大的MaxAera。
此数组的MaxAera为8
*/
package main

import (
	"fmt"
	"github.com/pengpeng1314/go/stack"
)

func maxRecSize(m [][]int) int {
	if len(m) < 1 || len(m[0]) < 1 {
		return 0
	}
	height := make([]int, len(m[0]))
	maxAera := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] != 0 {
				height[j] = height[j] + 1
			} else {
				height[j] = 0
			}
		}
		maxAera = max(maxAera, maxRecFromBottom(height))
	}
	return maxAera
}
func maxRecFromBottom(arr []int) int {
	if len(arr) < 1 {
		return 0
	}
	maxAera := 0
	s := stack.NewStack()
	for i := 0; i < len(arr); i++ {
		for !s.IsEmpty() && arr[i] <= arr[s.Top().(int)] {
			//注意L和R是一直在变化的
			R := i
			L := -1
			index := s.Pop().(int)
			if !s.IsEmpty() {
				L = s.Top().(int)
			}
			aera := (R - L - 1) * arr[index]
			maxAera = max(aera, maxAera)
		}
		s.Push(i)
	}
	for !s.IsEmpty() {
		R := len(arr)
		L := -1
		index := s.Pop().(int)
		if !s.IsEmpty() {
			L = s.Top().(int)
		}
		aera := (R - L - 1) * arr[index]
		maxAera = max(aera, maxAera)
	}
	return maxAera
}
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func main() {
	arr := []int{3, 5, 3, 5, 6, 0, 1, 5}
	maxAera := maxRecFromBottom(arr)
	fmt.Println(maxAera)
	matrix := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 0},
	}
	fmt.Println(maxRecSize(matrix))
}
