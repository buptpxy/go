/*
小和问题
在一个数组中， 每一个数左边比当前数小的数累加起来， 叫做这个数组的小和。 求一个数组的小和。
例子：
	[1,3,4,2,5]

	1左边比1小的数， 没有；
	3左边比3小的数， 1；
	4左边比4小的数， 1、 3；
	2左边比2小的数， 1；
	5左边比5小的数， 1、 3、 4、 2；
所以小和为1+1+3+1+1+3+4+2=16

思路：
	找一个数的左边有多少个数比它小，并把它们加起来，可以转换为找每个数右边有多少个数比它大，即知道了这个数要被加多少次。
	可以利用归并排序的每次比较过程，找到比它大的数
*/
package main

import (
	"fmt"
)

func merge1(arr []int, L int, mid int, R int) (sum int) {
	help := []int{}
	p := L
	q := mid + 1
	sum = 0
	for p <= mid && q <= R {
		if arr[p] < arr[q] {
			help = append(help, arr[p])
			fmt.Printf("help:%d\n", help)
			sum = sum + arr[p]*(R-q+1)
			fmt.Printf("sum:%d\n", sum)
			p++
		} else {
			help = append(help, arr[q])
			//sum = sum + arr[q]*(mid-p+1) 只算右边大的，不算左边大的
			q++
		}
	}
	if p <= mid {
		help = append(help, arr[p:mid+1]...)
		fmt.Printf("phelp:%d\n", help)
	}
	if q <= R {
		help = append(help, arr[q:R+1]...)
		fmt.Printf("qhelp:%d\n", help)
	}
	help1 := []int{}
	help1 = append(help1, arr[R+1:]...)
	arr = append(arr[:L], help...)
	arr = append(arr[:R+1], help1...)
	fmt.Printf("arr:%d\n", arr)
	return sum
}

func sortProcess1(arr []int, L int, R int) (sum int) {
	if L == R {
		return
	}
	mid := L + (R-L)>>1
	// sortProcess(arr, L, mid)
	// sortProcess(arr, mid+1, R)
	// merge(arr, L, mid, R)
	sum = sortProcess1(arr, L, mid) + sortProcess1(arr, mid+1, R) + merge1(arr, L, mid, R)
	return
}

func smallSum(arr []int) (sum int) {
	if len(arr) < 2 {
		return 0
	}
	sum = sortProcess1(arr, 0, len(arr)-1)
	return
}
func main() {
	arr := []int{1, 3, 4, 2, 5}
	sum := smallSum(arr)
	// sum := merge1(arr, 0, 2, 4)
	fmt.Println(sum)
}
