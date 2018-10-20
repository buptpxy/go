package main

/*
Q2：解决这个叫做 Fizz-Buzz[23] 的问题：
	编写一个程序，打印从 1 到 100 的数字。当是三的倍数就打印 “Fizz”
	代替数字，当是5的倍数就打印 “Buzz”。当数字同时是三和五的倍数时，打印 “FizzBuzz”。
*/
import (
	"fmt"
)

func main() {
	//方法一
	for i := 1; i < 101; i++ {
		switch true {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
	//方法二，使用了常数量
	const (
		FIZZ = 3
		BUZZ = 5
	)
	var p bool
	for i := 1; i < 100; i++ {
		p = false
		if i%FIZZ == 0 {
			fmt.Printf("Fizz")
			p = true
		}
		if i%BUZZ == 0 {
			fmt.Printf("Buzz")
			p = true
		}
		if !p {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}
