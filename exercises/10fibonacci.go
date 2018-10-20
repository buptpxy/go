package main

/*
Q10 斐波那契
	斐波那契数列以： 1, 1, 2, 3, 5, 8, 13, ... 开始。或者用数学形式表达： x1 = 1; x2 =1; xn = xn−1 + xn−2
	编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列。
*/
import (
	"fmt"
)

func fibonacci(num int) []int {
	slice := make([]int, num)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < num; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}
func main() {
	for _, v := range fibonacci(10) {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
