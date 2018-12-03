/*
1、生成窗口最大值数组
有一个整型数组arr和一个大小为w的窗口从数组的最左边滑到最右边，窗口每次向右边滑一个位置。
例如，数组为[4,3,5,4,3,3,6,7]，窗口大小为3时：

[4 3 5] 4 3 3 6 7 窗口中最大值为5
4 [3 5 4] 3 3 6 7 窗口中最大值为5
4 3 [5 4 3] 3 6 7 窗口中最大值为5
4 3 5 [4 3 3] 6 7 窗口中最大值为4
4 3 5 4 [3 3 6] 7 窗口中最大值为6
4 3 5 4 3 [3 6 7] 窗口中最大值为7

如果数组长度为n，窗口大小为w，则一共产生n-w+1个窗口的最大值。请实现一个函数，给定一个数组arr，窗口大小w。返回一个长度为n-w+1的数组res,res[i]表示每一种窗口状态下的最大值。以本题为例，结果应该返回[5,5,5,4,6,7]。
*/
package queue

import (
	"container/list"
	"fmt"
)

func maxNumInWindow(arr []int, w int) []int {
	n := len(arr)
	if n < w {
		return nil
	}
	qMax := list.New()

	index := 0
	res := make([]int, n-w+1)
	for i := 0; i < n; i++ {
		for qMax.Len() != 0 && arr[i] >= arr[qMax.Back().Value.(int)] {
			qMax.Remove(qMax.Back())
		}
		qMax.PushBack(i)
		if qMax.Front().Value.(int) == i-w {
			qMax.Remove(qMax.Front())
		}
		if i >= w-1 {
			res[index] = arr[qMax.Front().Value.(int)]
			index++
		}
	}
	return res
}
func main() {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	res := maxNumInWindow(arr, 3)
	fmt.Println(res)
	//q:=list.New()
	//q.PushBack(1)
	//fmt.Println(q.Back().Value)
}
