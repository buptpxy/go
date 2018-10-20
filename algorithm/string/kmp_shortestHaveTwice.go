/*
给定一个字符串str1，只能网str1的后面添加字符变成str2。
要求：str2必须包含两个str1，两个str1可以有重合，但是不能以同一位置开头。str2尽量短。最终返回str2。
思路：先求出str1的next数组，多求一位即可得到整个字符串的最长前缀和后缀长度，在str1后面加上不属于最长前缀的部分即可。
*/
package main

import "fmt"

func next(str string) []int {
	r := []rune(str)
	next := make([]int, len(r)+1) //next[i]默认为0
	next[0] = -1
	next[1] = 0
	for i := 2; i < len(r)+1; i++ {
		j := i - 1
		for j > 0 {
			if r[i-1] == r[next[j]] {
				next[i] = next[i-1] + 1
				break
			}
			j = next[j]
		}
	}
	return next
}

func getShortestStr(str1 string) string {
	nex := next(str1)
	n := nex[len(str1)]
	r1 := []rune(str1)
	r2 := r1[n:]
	r2 = append(r1, r2...)
	str2 := string(r2)
	return str2
}

func main() {
	str1 := "acdkacd"
	nex := next(str1)
	fmt.Println(nex)
	str2 := getShortestStr(str1)
	fmt.Println(str2)
}
