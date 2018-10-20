/*
【要求】给你一个数组arr， 和一个整数aim。
		如果可以任意选择arr中的数字， 能不能累加得到aim， 返回true或者false。
【思路】与求解字符串的子序列问题一样，每个元素有选和不选两种状态，如果有n个元素相加则最后有2^n种和，即查看这2^n个和中是否有和为aim的。

baseCase为sum=aim时，直接返回true; n=len(arr)时，表示遍历完sum也不等于aim，直接返回false。
其他位置的状态依赖于它下一个位置的状态的值，下一个位置的和有两种可能，包含当前位置或不包含当前位置。

改为动态规划矩阵dp时，行为数组中的元素索引0~len(arr)，列为0~MaxSum，
由递归终止状态可知，最后一行除了sum=aim那一列的元素为true，其他元素都为false。
由递归步骤可知，其他位置依赖的是它的下一行的元素或下一行、右边arr[i]列的元素
由递归初始状态可知，最后需要求的是dp[0][0]位置
*/
package main

import (
	"fmt"
)

func getSumRec(arr []int, index int, sum int, aim int) bool {
	if sum == aim {
		return true
	}
	if index == len(arr) {
		return false
	}
	return getSumRec(arr, index+1, sum, aim) || getSumRec(arr, index+1, sum+arr[index], aim)
}

func getSumDp(arr []int, aim int) bool {
	maxSum := 0
	for i := 0; i < len(arr); i++ {
		maxSum = maxSum + arr[i]
	}
	dp := make([][]bool, len(arr)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, maxSum+1)
	}
	for i := 0; i < maxSum+1; i++ {
		if i == aim {
			dp[len(arr)][i] = true
		} else {
			dp[len(arr)][i] = false
		}
	}
	//前面的值要依赖后面的值，故应该让后面的值先算出来
	for i := len(arr) - 1; i >= 0; i-- {
		for j := maxSum; j >= 0; j-- {
			dp[i][j] = dp[i+1][j]
			if j+arr[i] <= maxSum { //防止溢出
				dp[i][j] = dp[i+1][j] || dp[i+1][j+arr[i]]
			}
		}
	}
	// fmt.Println(dp)
	return dp[0][0]
}

func printAllSum(arr []int, index int, sum int) {
	if index == len(arr) {
		fmt.Println(sum)
		return
	}
	printAllSum(arr, index+1, sum)
	printAllSum(arr, index+1, sum+arr[index])
}

func main() {
	arr := []int{2, 3, 4, 1}
	fmt.Println(getSumRec(arr, 0, 0, 13))
	fmt.Println(getSumDp(arr, 13))
	// printAllSum(arr, 0, 0)
}
