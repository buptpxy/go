/*
 荷兰国旗问题
	给定一个数组arr， 和一个数num， 请把小于num的数放在数组的左边， 等于num的数放在数组的中间， 大于num的数放在数组的右边。
    要求额外空间复杂度O(1)， 时间复杂度O(N)
 思路：
 	在数组左边划定一个小于num的域，令less=-1，在数组右边划定一个大于num的域，令more=len(arr)。
 	遍历数组，当arr[p]=num，则直接p++,
 	当arr[p]<num，则swap(arr[less+1],arr[p]),less++,p++
 	当arr[p]>num，则swap(arr[more-1],arr[p]),more--
 	直到p=more，则遍历结束
*/
package main

import (
	"fmt"
)

func partition(arr []int, L, R, num int) (el, er int) {
	less := L - 1
	more := R + 1
	p := L
	for p < more {
		//最好用if else-if else的形式代替三个if并列
		if arr[p] == num {
			// fmt.Printf("arr[p] == num less:%d more:%d p:%d arr:%d \n ", less, more, p, arr)
			p++
		} else if arr[p] < num {
			arr[less+1], arr[p] = arr[p], arr[less+1]
			// fmt.Printf("arr[p] < num less:%d more:%d p:%d arr:%d \n ", less, more, p, arr)
			less++
			p++
		} else {
			arr[more-1], arr[p] = arr[p], arr[more-1]
			// fmt.Printf("arr[p] > num less:%d more:%d p:%d arr:%d \n ", less, more, p, arr)
			more--
		}

	}
	el = less + 1
	er = more - 1
	return
}

func main() {
	arr := []int{1, 4, 2, 3, 5, 7, 3}
	el, er := partition(arr, 0, 6, 3)
	fmt.Println(arr)
	fmt.Printf("el:%d er:%d \n", el, er)
}
