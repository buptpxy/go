/*
求最大值减去最小值小于或等于num的子数组数量
给定数组arr和整数num，返回有多少个子数组满足如下情况:
max(arr[i..j]) - min(arr[i..j]) <= num
max(arr[i..j])表示子数组arr[i..j]中的最大值，min(arr[i..j])表示子数组arr[i..j]中的最小值。如果数组长度为 N，请实现时间复杂度为 O(N)的解法。

思路：
子数组里的元素一定是在原数组中连续的，所以子数组总数为N*(N+1)/2
下标从L到R的一个子数组如果满足max-min<=num，则L到R之间的子数组一定都满足
下标从L到R的一个子数组如果不满足max-min<=num，则L往左边扩和R往右边扩的子数组一定不满足
所以即求从L到R的窗口的最大值和最小值
先让L为0，R一直往右直到扩不动为止（max-min>num），符合条件的子数组个数为R-L-1
然后让L往右开始移动，R再往右直到扩不动为止
*/
package queue

import (
	"container/list"
	"fmt"
)

func getValidSubArrNum(arr []int, num int) int {
	if len(arr) < 1 || num < 0 {
		return -1
	}
	qMax := list.New()
	qMin := list.New()
	L := 0
	R := 0
	res := 0
	for L < len(arr) {
		for R < len(arr) {
			for qMax.Len() != 0 && arr[R] >= arr[qMax.Back().Value.(int)] {
				qMax.Remove(qMax.Back())
			}
			qMax.PushBack(R)
			for qMin.Len() != 0 && arr[R] <= arr[qMin.Back().Value.(int)] {
				qMin.Remove(qMin.Back())
			}
			qMin.PushBack(R)
			if arr[qMax.Front().Value.(int)]-arr[qMin.Front().Value.(int)] > num {
				arr1 := arr[L:R]
				fmt.Println(arr1)
				break
			}
			R++
		}
		if qMax.Front().Value.(int) == L {
			qMax.Remove(qMax.Front())
		}
		if qMin.Front().Value.(int) == L {
			qMin.Remove(qMin.Front())
		}
		res = res + R - L
		L++
	}
	return res
}
func main() {
	arr := []int{4, 3, 5, 6, 1}
	fmt.Println(getValidSubArrNum(arr, 1))
}
