/*
求n!的结果
*/
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(4))
}
