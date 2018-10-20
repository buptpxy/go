/*
【要求】给你一个数组arr， 和一个整数aim。
		如果可以任意选择arr中的数字， 能不能累加得到aim， 返回true或者false。
【思路】与求解字符串的子序列问题一样，每个元素有选和不选两种状态，如果有n个元素相加则最后有2^n种和，即查看这2^n个和中是否有和为aim的。
*/
package main

import (
	"fmt"
)

func findAimRec(arr []int, i int, aim int) bool {
	if i == 0 {
		if arr[0] == aim {
			return true
		}
		return false
	}
	if aim-arr[i] < 0 {
		return findAimRec(arr, i-1, aim)
	}
	return findAimRec(arr, i-1, aim) || findAimRec(arr, i-1, aim-arr[i])
}
func getSum(arr []int) (sum int) {
	if arr == nil {
		return 0
	}
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}
func findAimDp(arr []int, aim int) bool {
	if getSum(arr) < aim {
		return false
	}
	dp := make([][]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]bool, aim+1)
	}
	for i := 0; i < aim+1; i++ {
		if arr[0] == i {
			dp[0][i] = true
		} else {
			dp[0][i] = false
		}
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < aim+1; j++ {
			if j-arr[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i]]
			}
		}
	}
	return dp[len(arr)-1][aim]
}
func main() {
	arr := []int{2, 3, 4, 1}
	fmt.Println(findAimRec(arr, len(arr)-1, 5))
	fmt.Println(findAimDp(arr, 5))
}
