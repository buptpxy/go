/*
【题目】给定一个数组arr， 和一个整数num， 求在arr中， 累加和等于num的最长子数组的长度。
	例子：
	arr = {7,3,2,1,1,7,7,7}， num = 7
	其中有很多的子数组累加和等于7， 但是最长的子数组是{3,2,1,1}， 所以返回其长度4。
【思路】
把求最长子数组长度转换成求最短子数组长度。
1、遍历数组先求从0位置开始到i位置结束的所以元素的和，把他们放在一个map中，map中记录这个和sum和最早出现的位置pos。
2、看aim=sum-num是否存在map中，若存在，则可知从0到pos累加和等于aim的最短子数组的长度，则pos到i为累加和等于num的最长子数组的长度。
然后看sum是否存在map中，若不存在，则把sum和i添加进map中
*/
package main

import "fmt"

func maxLength(arr []int, num int) int {
	posSum := make(map[int]int)
	posSum[0] = -1
	sum := 0
	maxLen := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if pos, ok := posSum[sum-num]; ok {
			pos = posSum[sum-num]
			maxLen = max(i-pos, maxLen)
		}
		if _, ok := posSum[sum]; !ok {
			posSum[sum] = i
		}
	}
	return maxLen
}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func main() {
	//m:=make(map[int]int)
	//fmt.Println(m[0])
	arr := []int{7, -3, -2, -1, -1, 7, 7, 7}
	fmt.Println(maxLength(arr, 7))
}
