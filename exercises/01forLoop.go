package main

/*
Q1. For-loop
	1. 创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值。
	2. 用 goto 改写 1 的循环。关键字 for 不可使用。
	3. 再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。
*/
import (
	"fmt"
)

func main() {
	//1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//2
	i := 0
Here:
	fmt.Println(i)
	i++
	if i < 10 {
		goto Here
	}

	//3
	array := [...]int{1, 2, 3, 4, 5}
	for _, v := range array {
		fmt.Println(v)
	}

}
