/*
【暴力递归】
1， 把问题转化为规模缩小了的同类问题的子问题
2， 有明确的不需要继续进行递归的条件(base case)
3， 有当得到了子问题的结果之后的决策过程
4， 不记录每一个子问题的解

【动态规划】
1， 从暴力递归中来
2， 将每一个子问题的解记录下来， 避免重复计算
3， 把暴力递归的过程， 抽象成了状态表达
4， 并且存在化简状态表达， 使其更加简洁的可能

【汉诺塔问题】: 打印n层汉诺塔从最左边移动到最右边的全部过程
【要求】一次只能移动一个块，大的块不能放在小的上面，有三个杆子，一个为left，一个help，一个right
【思路】n层汉诺塔移动的过程要依赖n-1层汉诺塔移动的过程，baseCase为1层汉诺塔移动的情况

*/
package main

import (
	"fmt"
)

func hanoi(n int, from string, to string, help string) {
	if n == 1 {
		fmt.Printf("move from %s to %s \n", from, to)
		return
	}
	hanoi(n-1, from, help, to)
	hanoi(1, from, to, help)
	hanoi(n-1, help, to, from)
}

func main() {
	hanoi(3, "left", "right", "middle")
}
