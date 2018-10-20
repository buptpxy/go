/*
【要求】打印一个字符串的全部排列
【思路】如果这个字符串包含n个字符，就相当于每一位上的数可以和其它n-1个位子上的数交换
用一个index记录交换到了第几轮，后一轮的状态依赖于前一轮的index值，后一轮从index+1位置开始交换
【进阶】打印一个字符串的全部排列， 要求不要出现重复的排列
【思路】要求不重复则可以考虑把排列的字符串都存进哈希表的key中
*/

package main

import (
	"fmt"
)

//直接对字符数组进行操作，每次小递归都依赖的前一次小递归传入的值，故最后得到的值不对
func process(chs []rune, i int) {
	// fmt.Println("chs: ", string(chs))
	if i == len(chs) {
		fmt.Println("%s ", string(chs))
	}
	for j := i; j < len(chs); j++ {
		chs[i], chs[j] = chs[j], chs[i]
		process(chs, i+1)
	}
}

var m = make(map[string]int)

//字符串是3位则分为三次大递归，每次小递归中使用的str和大递归传入的str相同
func process1(str string, i int) {
	// fmt.Println("str: ", str)
	chs := []rune(str) //并未对原字符串进行操作，而是构造了一个字符串的拷贝
	if i == len(chs) {
		fmt.Printf("%s ", string(chs))
	}
	for j := i; j < len(chs); j++ {
		chs[i], chs[j] = chs[j], chs[i]
		process1(string(chs), i+1)
	}
}
func process2(str string, i int) {
	chs := []rune(str)
	if i == len(chs) {
		m[string(chs)] = i
	}
	for j := i; j < len(chs); j++ {
		chs[i], chs[j] = chs[j], chs[i]
		process2(string(chs), i+1)
	}
}
func printPermutations(str string) {
	// r := []rune(str)
	// process(r, 0)
	process1(str, 0)
	fmt.Println()
	process2(str, 0)
	for k, _ := range m {
		fmt.Printf("%s ", k)
	}
	fmt.Println()
}
func test(str []rune) {
	str[0], str[1] = str[1], str[0]
}
func main() {
	str := "hhl"
	printPermutations(str)
	// r := []rune(str)
	// test(r)
	// fmt.Println("r: ", string(r))  //ehl，原字符数组会被改变
	// fmt.Println("str: ", str)      //hel，原字符串不会被改变
	// m := make(map[string]int)
	// m["key"] = 1
	// m["key"] = 2
	// fmt.Println(m) //2会覆盖1
}
