package main

import "fmt"

/*
Manacher算法扩展题目
给定一个字符串str1， 只能往str1的后面添加字符变成str2， 要求str2
整体都是回文串且最短。
举例：
str1 = ABC12321, 返回ABC12321CBA
思路：即求出str1的pArr[]数组，直到R到达最后一个元素，即求出来包含最后一个元素的最长回文序列，再把剩下的元素逆序加在str1后面即可得到str2。
*/
func manacherString1(str string) []rune {
	r1 := []rune(str)
	r2 := []rune{'#'}
	for i := 0; i < len(r1); i++ {
		r2 = append(r2, r1[i])
		r2 = append(r2, '#')
	}
	return r2
}

//返回包含最后一个元素的最长回文序列的第一个元素所在位置
func manacher1(str string) int {
	r := manacherString1(str)
	pArr := make([]int, len(r))
	R := -1
	C := -1
	index := 0
	for i := 0; R < len(r); i++ {
		if i < R {
			pArr[i] = min(R-i, pArr[2*C-i])
		}
		for i+pArr[i] < len(r) && i-pArr[i] > -1 {
			if r[i+pArr[i]] == r[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}
		index = i
	}
	return index - pArr[index] - 1
}
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
func getStr2(str1 string) string {
	r := []rune(str1)
	index := manacher1(str1)
	for i := index - 1; i >= 0; i-- {
		r = append(r, r[i])
	}
	return string(r)
}
func main() {
	str := "bcbaab"
	fmt.Println(manacher1(str))
	fmt.Println(getStr2(str))
}
