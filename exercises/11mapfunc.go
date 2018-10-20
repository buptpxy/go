package main

/*
Q11. map 函数
	map() 函数是一个接受一个函数和一个列表作为参数的函数。
	函数应用于列表中的每个元素，而一个新的包含有计算结果的列表被返回。因此：
	map(f(), (a1, a2, ..., an−1, an)) = (f(a1), f(a2), ..., f(an−1), f(an))
	1. 编写 Go 中的简单的 map() 函数。它能工作于操作整数的函数就可以了。
	2. 扩展代码使其工作于字符串列表。
*/
import (
	"fmt"
)

func Map(f func(int) int, intslice []int) []int {
	slice := []int{} //搭配slice[k] = f(v)是错误的写法，会报错index out of range，需搭配slice = append(slice, f(v))
	// slice := make([]int, len(intslice))
	for _, v := range intslice {
		// slice[k] = f(v)
		slice = append(slice, f(v))
	}
	return slice
}
func main() {
	f := func(a int) int {
		return a * a
	}
	s := []int{3, 4, 2}
	fmt.Println(Map(f, s))
}
