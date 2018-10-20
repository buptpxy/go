/*
给你一个二维数组， 二维数组中的每个数都是正数，
要求从左上角走到右下角， 每一步只能向右或者向下。
沿途经过的数字要累加起来。 返回最小的路径和。
【思路】
baseCase为[0,0]位置，最短路径即为当前位置元素值。
其他位置的最短路径为其左边和上边元素的最小值加当前位置的元素值。
第一列位置的元素值为上边元素值加当前位置的元素值。
第一行位置的元素值为左边元素值加当前位置的元素值。
*/
package main

import (
	"fmt"
)

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
func minPathRec(m [][]int, i int, j int) int {
	if i == 0 && j == 0 {
		return m[0][0]
	}
	if i == 0 {
		return minPathRec(m, i, j-1) + m[0][j]
	}
	if j == 0 {
		return minPathRec(m, i-1, j) + m[i][0]
	}
	return min(minPathRec(m, i-1, j), minPathRec(m, i, j-1)) + m[i][j]
}
func minPathDp(m [][]int) int {
	row := len(m)
	if row == 0 {
		return 0
	}
	col := len(m[0])
	dp := m
	dp[0][0] = m[0][0]
	//第一列
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + m[i][0]
	}
	//第一行
	for i := 1; i < col; i++ {
		dp[0][i] = dp[0][i-1] + m[0][i]
	}
	//其他位置
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + m[i][j]
		}
	}
	return dp[row-1][col-1]
}
func main() {
	m := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}
	fmt.Println(minPathRec(m, len(m)-1, len(m[0])-1))
	fmt.Println(minPathDp(m))
}
