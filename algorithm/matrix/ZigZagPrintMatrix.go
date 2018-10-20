/*
“ 之” 字形打印矩阵
【 题目】 给定一个矩阵matrix， 按照“ 之” 字形的方式打印这个矩阵，
例如： 1 2 3 4 5 6 7 8 9 10 11 12
“ 之” 字形打印的结果为： 1， 2， 5， 9， 6， 3， 4， 7， 10， 11，
8， 12
【 要求】 额外空间复杂度为O(1)。
思路： 分解为每一条斜边的打印
*/

package main

import (
	"fmt"
)

func printLevel(aRow int, aCol int, bRow int, bCol int, m [][]int, flag bool) {
	if !flag {
		cRow, cCol := bRow, bCol
		for cCol <= aCol {
			fmt.Printf("%d ", m[cRow][cCol])
			cRow--
			cCol++
		}

	} else {
		cRow, cCol := aRow, aCol
		for cRow <= bRow {
			fmt.Printf("%d ", m[cRow][cCol])
			cRow++
			cCol--
		}

	}
}

func printmatrixZigZag(m [][]int) {
	aRow, aCol, bRow, bCol := 0, 0, 0, 0
	endR := len(m) - 1
	endC := len(m[0]) - 1
	flag := false
	for aRow <= endR {
		printLevel(aRow, aCol, bRow, bCol, m, flag)
		if aCol < endC {
			aCol++
		} else {
			aRow++
		}
		if bRow < endR {
			bRow++
		} else {
			bCol++
		}
		flag = !flag
	}
	fmt.Printf("\n")
}
func main() {
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	printmatrixZigZag(matrix1)
	matrix := [][]int{{1}}
	printmatrixZigZag(matrix)
}
