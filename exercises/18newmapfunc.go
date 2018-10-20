package main

/*Q18. 使用 interface 的 map 函数
	使用练习 Q11 的答案，利用 interface 使其更加通用。让它至少能同时工作于
int 和 string。
*/
import (
	"fmt"
)

type m interface{}

func mul2(arg m) m {
	switch arg.(type) {
	case int:
		return arg.(int) * 2
	case string:
		return arg.(string) + arg.(string)
	default: //没有default且在switch外面也没有return时会报错：missing return at end of function
		return arg
	}
}

func mapFunc(f func(m) m, args []m) (funcs []m) {
	for _, v := range args {
		funcs = append(funcs, f(v))
	}
	return
}

func main() {
	// var ints m = []{1,2,3}
	// var strings m = []{"hello","pxy","!"}
	//m类型的slice应该如下这样写
	ints := []m{1, 2, 3}
	strings := []m{"hello", "pxy", "!"}
	mul2Ints := mapFunc(mul2, ints)
	mul2Strings := mapFunc(mul2, strings)
	fmt.Printf("%d \n", mul2Ints)
	fmt.Printf("%s \n", mul2Strings)
}
