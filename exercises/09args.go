package main

/*
Q9. 变参
	编写函数接受整数类型变参，并且每行打印一个数字。(要使用 ... 语法来实现函数接受若干个数字作为变参。)
*/
import (
	"fmt"
)

func args(numbers ...int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
}

func main() {
	args(1, 2, 5, 3, 4)
	args(4, 5, 3)
}
