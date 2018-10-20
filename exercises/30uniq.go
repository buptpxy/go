package main

/*
30. 编写一个 Go 程序模仿 Unix 命令 uniq 的功能。程序应当像下面这样运行，提供一个下面这样的列表：
	'a' 'b' 'a' 'a' 'a' 'c' 'd' 'e' 'f' 'g'
	它将打印出没有后续重复的项目：
	'a' 'b' 'a' 'c' 'd' 'e' 'f'
*/
import (
	"fmt"
)

func main() {
	slice := []string{"a", "b", "a", "a", "a", "c", "d", "e", "f", "g"}
	first := slice[0]
	fmt.Printf("%s ", first)
	for _, v := range slice[1:] {
		if v != first {
			fmt.Printf("%s ", first)
			first = v
		}
	}
	fmt.Println()
}
