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
	从右边开始遍历这个数组，如果p前面的数q比p小，则p的小和为q的小和加q，如果q>=p，则继续向前寻找直到数组开头
	此方法失败
*/
package main

import (
	"fmt"
)

var sump, sumi, q int

// var sumarr := []int{}
//计算比pos位置上的数小的数的和
func smallNumSum(arr []int, pos int) int {
	if pos < 1 {
		return 0
	}
	for i := pos - 1; i >= 0; i-- {
		if arr[i] < arr[pos] { //找到pos前面最近的小于arr[pos]的数arr[i]
			if arr[i] >= arr[q] {
				fmt.Printf("arr[%d]:%d arr[%d]:%d sump:%d q:%d\n", i, arr[i], pos, arr[pos], sump, q)
				q = i
				sumi := smallNumSum(arr, q)
				sump = sumi + arr[i] //则pos的小和等于i的小和加i
				// return nsum
				fmt.Printf("sump:%d = arr[%d]:%d + sumi: %d \n", sump, i, arr[i], sumi)
			}
			if i == 0 {
				fmt.Printf("sump:%d\n", sump)
				return sump + arr[i] - arr[i]
			}

		}
	}
	return sump
}

func smallSum(arr []int) (ssum int) {
	for i := 0; i < len(arr); i++ {
		nsum := smallNumSum(arr, i)
		fmt.Println("nsum", nsum)
		ssum = ssum + nsum
		fmt.Println("ssum", ssum)
	}
	return
}

func main() {
	arr := []int{1, 3, 4, 2, 5}
	// ssum := smallSum(arr)
	nsum := smallNumSum(arr, 1)
	fmt.Println(nsum)
	// arr1 := [...]int{1, 2, 3}
	// fmt.Println(arr1)
}
