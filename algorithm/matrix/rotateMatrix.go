/*
旋转正方形矩阵
【 题目】 给定一个整型正方形矩阵matrix， 请把该矩阵调整成顺时针旋转90度的样子。
【 要求】 额外空间复杂度为O(1)。
思路：选择90度即以正方形的中间为轴，上下左右四条边的每一条边的其中一个元素之间的互换，上到右，右到下，下到左，左到上
	而且是一圈一圈的换
*/
package main

import (
	"fmt"
)

func rotateEdge(aRow int, aCol int, bRow int, bCol int, m [][]int) {

	times := bRow - aRow
	for i := 0; i < times; i++ {
		m[aRow][aCol+i], m[aRow+i][bCol], m[bRow][bCol-i], m[bRow-i][aCol] = m[bRow-i][aCol], m[aRow][aCol+i], m[aRow+i][bCol], m[bRow][bCol-i]
	}
}
func rotate(m [][]int) {
	aRow, aCol, bRow, bCol := 0, 0, len(m)-1, len(m[0])-1
	for aRow < bRow {
		rotateEdge(aRow, aCol, bRow, bCol, m)
		aRow++
		aCol++
		bRow--
		bCol--
	}
}
func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix1)
	fmt.Println(matrix1)
}
