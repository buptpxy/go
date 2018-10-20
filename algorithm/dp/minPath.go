/*
给你一个二维数组， 二维数组中的每个数都是正数，
要求从左上角走到右下角， 每一步只能向右或者向下。
沿途经过的数字要累加起来。 返回最小的路径和。
【思路】先写出暴力递归版本。
		每一步到右下角位置的最短路径与它前面走过的路径无关，是无后效性的。
		baseCase为矩阵右下角的位置，它到右下角的最短路径即为它自己的值。
		矩阵的最后一行，某个元素到右下角的最短路径即为它自己的值加它右边元素到右下角的最短路径。
		矩阵的最后一列，某个元素到右下角的最短路径即为它自己的值加它下边元素到右下角的最短路径。
		矩阵中间的某个元素到右下角的最短路径即为它自己的值加它右边元素的最短路径和下边元素的最短路径中最小者。
		但此递归版本存在多次重复计算的情况。故可以改成动态规划。
*/
package main

import (
	"fmt"
)

func minPathRec(m [][]int, i int, j int) int {
	if i == len(m)-1 && j == len(m[0])-1 {
		return m[i][j]
	}
	if i == len(m)-1 && j < len(m[0])-1 {
		return m[i][j] + minPathRec(m, i, j+1)
	}
	if i < len(m)-1 && j == len(m[0])-1 {
		return m[i][j] + minPathRec(m, i+1, j)
	}

	return m[i][j] + min(minPathRec(m, i, j+1), minPathRec(m, i+1, j))
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
func minPathDp(m1 [][]int) int {
	if len(m1) == 0 || len(m1[0]) == 0 {
		return 0
	}
	m := m1
	r := len(m) - 1
	c := len(m[0]) - 1
	for i := c - 1; i >= 0; i-- {
		m[r][i] = m1[r][i] + m[r][i+1]
	}
	for i := r - 1; i >= 0; i-- {
		m[i][c] = m1[i][c] + m[i+1][c]
	}
	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			m[i][j] = m1[i][j] + min(m[i+1][j], m[i][j+1])
		}
	}
	return m[0][0]
}
func main() {
	m := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}
	fmt.Println(minPathRec(m, 0, 0))
	fmt.Println(minPathDp(m))
}
