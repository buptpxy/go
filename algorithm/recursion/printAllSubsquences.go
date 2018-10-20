/*
打印一个字符串的全部子序列， 包括空字符串。子序列和子串不同，子串必须是原字符串中一串连续的字符构成的
【思路】
把字符串变成字符数组，
每一位上的字符都有存在或不存在两种状态，
第n位的状态不依赖于前一位传给这一位的状态值，
当前的总状态为前一位传过来的状态加当前位子的状态
baseCase为最后一位时，总的状态即为前一位传过来的状态

打印全部子序列是一个过程，终止条件是遍历完整个序列时，直接打印“待打印序列”，
每一步的过程是先把不包含当前元素的序列加入“待打印序列”，再把包含当前元素的序列加入“待打印序列”，然后index+1
*/

package main

import (
	"fmt"
)

func process(str []rune, res []rune, index int) {
	if index == len(str) {
		fmt.Println(string(res))
		return
	}
	process(str, res, index+1)
	res = append(res, str[index])
	process(str, res, index+1)
}

func printSub(str string) {
	r := []rune(str)
	process(r, []rune{}, 0)
}

func main() {
	str := "hel"
	printSub(str)
}
