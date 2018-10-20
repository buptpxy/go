/*
【要求】母牛每年生一只母牛， 新出生的母牛成长三年后也能每年生一只母牛， 假设不会死。 求N年后， 母牛的数量。
【思路】先写出递归版本。
		1、 baseCase为第一、二、三年 num(n)=n
		2、往后每一年的数量等于它的前一年的数量加前三年的数量 num(n) = num(n-1) + num(n-3)

【进阶】如果每只母牛只能活10年， 求N年后， 母牛的数量。
*/
package main

import (
	"fmt"
)

func cowNum(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	return cowNum(n-1) + cowNum(n-3)
}
func cowNum2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	res := 3
	pre := 2
	prepre := 1
	tmp1 := 0
	tmp2 := 0
	for i := 4; i <= n; i++ {
		tmp1 = res
		tmp2 = pre
		res = res + prepre
		pre = tmp1
		prepre = tmp2
	}
	return res
}
func main() {
	fmt.Println(cowNum(6))
	fmt.Println(cowNum2(6))
}
