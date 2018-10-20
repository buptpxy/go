/*
Q26. Channel
	1. 修改在练习 Q1 中创建的程序，换句话说，主体中调用的函数现在是一个
		goroutine 并且使用 channel 通讯。不用担心 goroutine 是如何停止的。
	2. 在完成了问题 1 后，仍有一些待解决的问题。其中一个麻烦是 goroutine 在
		main.main() 结束的时候，没有进行清理。更糟的是，由于 main.main() 和
		main.shower() 的竞争关系，不是所有数字都被打印了。本应该打印到 9，但
		是有时只打印到 8。添加第二个退出 channel，可以解决这两个问题。试试吧。
	//main()是最先执行的goroutine,要先在main()函数里面把值传给channel，channel才可以把值传给别的goroutine
	//go 只能加在函数前
*/

package main

import (
	"fmt"
)

// func shower() chan int {
// 	next := make(chan int)
// 	// done := make(chan bool)
// 	var i int
// 	go func() {
// 		for {
// 			next <- i
// 			i++
// 		}
// 	}()
// 	// done <- true

// 	return next
// }
// func main() {
// 	counter := shower()

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%d \n", <-counter)
// 	}
// 	// defer func() {
// 	// 	close(counter)
// 	// }()

// }

func main() {
	ch := make(chan int, 10)
	defer func() {
		close(ch)
	}()
	done := make(chan bool)
	defer func() {
		close(done)
	}()

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("i: ", i)
	}
	//当声明的通道是默认缓冲时（默认缓冲是1），go 语句不能放在for循环语句之后，会报错fatal error: all goroutines are asleep - deadlock!
	//因为通道会阻塞，一次只能存一个值，但是for循环给它传第二个值时第一个值未被读出，通道就会阻塞
	//当加大通道缓冲或者把这句go语句放在赋值语句之前，则不会报错
	go shower(ch, done)
	done <- true

}
func shower(c chan int, b chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Println(j)
		case <-b:
			break
		}
	}
}
