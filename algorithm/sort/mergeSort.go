/*
归并排序
把数组逐次二分为小的部分，每次先把左右两部分数组排好序，然后把整个数组排好序
时间复杂度为O(nlogn) 空间复杂度为O(n)
*/
package main

import (
	"fmt"
)

func merge(arr []int, L int, mid int, R int) {
	// help := make([]int, len(arr)) 用make初始化会使help的每个元素用0填充
	help := []int{}
	p := L
	q := mid + 1
	for p <= mid && q <= R {
		if arr[p] < arr[q] {
			help = append(help, arr[p])
			p++
		} else {
			help = append(help, arr[q])
			q++
		}
	}
	if p <= mid {
		help = append(help, arr[p:mid+1]...) //注意此处是mid+1但实际上追加的是0-mid的元素
	}
	if q <= R {
		help = append(help, arr[q:R+1]...)
	}

	help1 := []int{}
	help1 = append(help1, arr[R+1:]...)
	arr = append(arr[:L], help...) //此处是L而不是L+1，表示从L开始追加，可以理解成arr[a:b]实际上是arr[a:b)，后面是个开区间
	arr = append(arr[:R+1], help1...)

}
func sortProcess(arr []int, L int, R int) {
	if L == R {
		return
	}
	// mid := (L + R) / 2 //L+R容易溢出，可改成L+(R-L)/2,即L+(R-L)>>1，右移位表示/2，位运算比除快
	mid := L + (R-L)>>1
	sortProcess(arr, L, mid)
	sortProcess(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func mergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	sortProcess(arr, 0, len(arr)-1)
}

// func add(arr []int, num int) {
// 	arr1 := []int{}
// 	arr = append(arr1, num)
// }
func main() {
	arr := []int{8, 5, 4, 7, 1, 9, 3}
	// arr := []int{4, 5, 6, 7, 1, 2, 3}
	// merge(arr, 1, 3, 5)
	// sortProcess(arr, 1, 5)
	// mergeSort(arr)
	// add(arr, 1)
	fmt.Printf("%d\n", arr)
}
