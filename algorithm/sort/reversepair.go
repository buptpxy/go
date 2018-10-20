/*
逆序对问题
	在一个数组中， 左边的数如果比右边的数大， 则折两个数构成一个逆序对， 请打印所有逆序对。
	思路：
	可利用归并排序的比较过程
*/
package main

import (
	"fmt"
)

func merge2(arr []int, L int, mid int, R int) {
	help := []int{}
	p := L
	q := mid + 1
	for p <= mid && q <= R {
		if arr[p] > arr[q] {
			for i := p; i <= mid; i++ {
				fmt.Printf("(%d,%d) \n", arr[i], arr[q])
			}
			help = append(help, arr[q])
			q++
		} else {
			help = append(help, arr[p])
			p++
		}
	}
	if p <= mid {
		help = append(help, arr[p:mid+1]...)
	}
	if q <= R {
		help = append(help, arr[q:R+1]...)
	}
	help1 := []int{}
	help1 = append(help1, arr[R+1:]...)
	arr = append(arr[:L], help...)
	arr = append(arr[:R+1], help1...)
}
func sortProcess2(arr []int, L int, R int) {
	if L == R {
		return
	}
	mid := L + (R-L)>>1
	sortProcess2(arr, L, mid)
	sortProcess2(arr, mid+1, R)
	merge2(arr, L, mid, R)
}

func reversePair(arr []int) {
	if len(arr) < 2 {
		return
	}
	sortProcess2(arr, 0, len(arr)-1)
}
func main() {
	arr := []int{1, 3, 4, 2, 5, 1}
	reversePair(arr)
}
