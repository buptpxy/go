/*
1、生成窗口最小值数组
有一个整型数组arr和一个大小为w的窗口从数组的最左边滑到最右边，窗口每次向右边滑一个位置。
例如，数组为[4,3,5,4,3,3,6,7]，窗口大小为3时：

[4 3 5] 4 3 3 6 7 窗口中最小值为3
4 [3 5 4] 3 3 6 7 窗口中最小值为3
4 3 [5 4 3] 3 6 7 窗口中最小值为3
4 3 5 [4 3 3] 6 7 窗口中最小值为3
4 3 5 4 [3 3 6] 7 窗口中最小值为3
4 3 5 4 3 [3 6 7] 窗口中最小值为3

如果数组长度为n，窗口大小为w，则一共产生n-w+1个窗口的最小值。请实现一个函数，给定一个数组arr，窗口大小w。返回一个长度为n-w+1的数组res,res[i]表示每一种窗口状态下的最小值。以本题为例，结果应该返回[3,3,3,3,3,3]。

思路：
使用一个双向队列来存储当前窗口内最小值，双向队列从左到右存储的arr[index]依次增大，
如果当前arr[i]>qMin.Back()则队列最后一个元素出队，直到最后一个元素大于当前元素为止。然后当前元素的index入队，否则当前元素的index直接入队。
qMin.Front()如果等于i-w，则Front因窗口过期而出队。
当i走到w-1时，就可开始依次把当前队列的Front及最小值放入结果数组中了。
*/
package queue

import (
	"container/list"
	"fmt"
)

func minNumInWindow(arr []int, w int) []int {
	n := len(arr)
	if n < w || w < 1 {
		return nil
	}
	index := 0
	res := make([]int, n-w+1)
	qMin := list.New()
	for i := 0; i < n; i++ {
		for qMin.Len() != 0 && arr[i] <= arr[qMin.Back().Value.(int)] {
			qMin.Remove(qMin.Back())
		}
		qMin.PushBack(i)
		if qMin.Front().Value.(int) == i-w {
			qMin.Remove(qMin.Front())
		}
		if i >= w-1 {
			res[index] = arr[qMin.Front().Value.(int)]
			index++
		}
	}
	return res
}
func main() {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	res := minNumInWindow(arr, 3)
	fmt.Println(res)
}
