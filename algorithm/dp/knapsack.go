/*
【背包问题】给定两个数组w和v，两个数组长度相等，w[i]表示第i件商品的重量，v[i]表示第i件商品的价值。
 再给定一个整数bag，要求你挑选商品的重量加起来一定不能超过bag，返回满足这个条件下，你能获得的最大价值。

*/

package main

import (
	"fmt"
)

//重量不允许超过maxW，遍历到i时，如果不要i，则bags剩余空间为maxW，总价值还是i-1时的返回值
//如果要i，则bag剩余空间为maxW-w[i1]，总价值为i-1时的价值加v[i]
func maxVInBagRec(w []int, v []int, i int, bag int) int {
	if i == 0 {
		return 0
	}
	if bag-w[i] < 0 {
		return maxVInBagRec(w, v, i-1, bag)
	}
	return max(maxVInBagRec(w, v, i-1, bag), maxVInBagRec(w, v, i-1, bag-w[i])+v[i])

}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
func getSum(arr []int) (sum int) {
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

//动态规划矩阵：行为i，列为总重量，值为总价值
func maxVInBagDp(w []int, v []int, bag int) int {
	if getSum(w) <= bag {
		return getSum(v)
	}
	col := bag + 1
	row := len(w)
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if j-w[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}
	return dp[row-1][bag]
}
func main() {
	w := []int{3, 2, 4, 7}
	v := []int{5, 6, 13, 19}
	fmt.Println(maxVInBagRec(w, v, 3, 5))
	fmt.Println(maxVInBagDp(w, v, 5))
}
