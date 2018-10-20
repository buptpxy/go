/*
kmp 算法用来检测一个字符串str2是否是另一个字符串str1的子串。如果是，则返回匹配的第一个位置。
有一个寻找最长前缀和后缀的next数组。
比如abcabc的最长前缀就是3。
*/
package main

import (
	"fmt"
)

func getNext(str string) []int {
	r := []rune(str)
	next := make([]int, len(r)+1)
	next[0] = -1
	next[1] = 0
	for i := 2; i < len(next); i++ {
		next[i] = 0
		for j := i - 1; j > 0; {
			if r[next[j]] == r[i-1] {
				next[i] = next[j] + 1
				break
			} else {
				j = next[j]
			}
		}
	}
	return next
}

func kmp(str1 string, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 || len(str1) < len(str2) {
		return -1
	}
	next := getNext(str2)
	r1 := []rune(str1)
	r2 := []rune(str2)
	i := 0
	j := 0
	for i < len(r1) && j < len(r2) {
		if r1[i] == r2[j] {
			i++
			j++
		} else if next[j] == -1 {
			i++
		} else {
			j = next[j]
		}
	}
	if j == len(r2) {
		return i - j
	}
	return -1
}

func main() {
	str1 := "abcaababcabcde"
	str2 := "abcabc"
	next := getNext(str2)
	fmt.Println(next)
	fmt.Println(kmp(str1, str2))
}
