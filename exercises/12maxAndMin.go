package main

/*
Q12. 最小值和最大值
	1. 编写一个函数，找到 int slice ([]int) 中的最大值。
	2. 编写一个函数，找到 int slice ([]int) 中的最小值。
*/
import (
	"fmt"
)

func maxAndMin(slice []int) (max, min int) {
	max = slice[0]
	min = slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}
func main() {
	s := []int{4, 2, 7, 4, 7}
	max, min := maxAndMin(s)
	fmt.Printf("max is : %d , min is : %d \n", max, min)
}
