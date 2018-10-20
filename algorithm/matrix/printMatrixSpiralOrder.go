/*转圈打印矩阵
【 题目】 给定一个整型矩阵matrix， 请按照转圈的方式打印它。
例如： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
打印结果为： 1， 2， 3， 4， 8， 12， 16， 15， 14， 13， 9， 5， 6， 7， 11， 10
【 要求】 额外空间复杂度为O(1)。
	思路：把每一个圈的打印当成一整个步骤，而不是把每一步的打印当成一个步骤。知道左上角和右下角的点即可知道整个圈
*/
package main

import (
	"fmt"
)

func printMatrixSpiralOrder(m [][]int) {
	aRow, aCol, bRow, bCol := 0, 0, len(m)-1, len(m[0])-1
	for aRow <= bRow && aCol <= bCol {
		printedge(aRow, aCol, bRow, bCol, m)
		aRow++
		aCol++
		bRow--
		bCol--
	}
	fmt.Printf("\n")
}

func printedge(aRow int, aCol int, bRow int, bCol int, m [][]int) {
	if aRow == bRow {
		for i := aCol; i <= bCol; i++ {
			fmt.Printf("%d ", m[aRow][i])
		}
	} else if aCol == bCol {
		for i := aRow; i <= bRow; i++ {
			fmt.Printf("%d ", m[i][aCol])
		}
	} else {
		cRow, cCol := aRow, aCol
		for cCol != bCol {
			fmt.Printf("%d ", m[cRow][cCol])
			cCol++
		}
		for cRow != bRow {
			fmt.Printf("%d ", m[cRow][cCol])
			cRow++
		}
		for cCol != aCol {
			fmt.Printf("%d ", m[cRow][cCol])
			cCol--
		}
		for cRow != aRow {
			fmt.Printf("%d ", m[cRow][cCol])
			cRow--
		}
	}
}
func main() {

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	printMatrixSpiralOrder(matrix)
}
