/*
【要求】给定一个字符串类型的数组strs，找到一种拼接方式，
使得把所有字符串拼起来之后形成的字符串具有最低的字典序。
【思路】如果ab形式的字典序小于ba形式的字典序，则按照ab形式拼接。否则按ba形式拼接。
		技巧是不比较a和b的字典序，而是比较拼接后的字典序
*/
package main

import (
	"fmt"
	"sort"
)

type Strs []string

func (s Strs) Len() int {
	return len(s)
}

func (s Strs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Strs) Less(i, j int) bool {
	return s[i]+s[j] < s[j]+s[i]
}
func lowestLexicography(s Strs) string {
	sort.Sort(s)
	res := ""
	for _, v := range s {
		res = res + v
	}
	return res
}
func main() {
	s := Strs{"cd", "ab", "ef"}
	fmt.Println(lowestLexicography(s))
	// fmt.Println("a"+"b" < "b"+"a")
}
