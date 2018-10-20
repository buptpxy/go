package main

/*
Q6. 整数顺序 (考察多值返回)
	编写函数，返回其（两个）参数正确的（自然）数字顺序：
	f(7,2) -> 2,7
	f(2,7) -> 2,7
*/
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func intorder(a int, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
func intsliceorder(is []int) []int {
	sort.Ints(is)
	return is
}
func main() {
	buf := bufio.NewReader(os.Stdin)
	str, _ := buf.ReadString('\n')
	// a := strings.Fields(str)[0]
	// b := strings.Fields(str)[1]
	// c, _ := strconv.Atoi(a)
	// d, _ := strconv.Atoi(b)
	// c, d = intorder(c, d)
	// fmt.Printf("正确的参数顺序为： %d %d \n", c, d)
	//intslice := []int(str) //错误的代码：cannot convert str (type string) to type []int
	intslice := []int{}
	for _, v := range strings.Fields(str) {
		ints, _ := strconv.Atoi(v)
		intslice = append(intslice, ints)
	}
	fmt.Printf("正确的参数顺序为： %d \n", intsliceorder(intslice))
}
