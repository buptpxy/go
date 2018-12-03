/*
给定一个字符串str， str表示一个公式， 公式里可能有整数、 加减乘除符号和左右括号， 返回公式的计算结果。
【举例】
str="48*((70-65)-43)+8*1"， 返回-1816。
str="3+1*4"， 返回7。 str="3+(1*4)"， 返回7。
【说明】
1．可以认为给定的字符串一定是正确的公式， 即不需要对str做公式有效性检查。
2．如果是负数， 就需要用括号括起来， 比如"4*(-3)"。 但如果负数作为公式的开头或括号部分的开头， 则可以没有括号， 比如"-3*4"和"(-3*4)"都是合法的。
3．不用考虑计算过程中会发生溢出的情况
【思路】
1、如果是没有括号的形式，则直接用栈，当前字符为数字时处理成数字然后入栈，为+或-则直接进栈，为*或/则先计算再进栈
2、如果是有括号的形式，则把括号部分当黑盒处理，由子函数返回给父函数
*/
package main

import (
	"container/list"
	"fmt"
	"strconv"
	"unicode"
)

type Res struct {
	num int
	pos int
}

func strToRune(str string) []rune {
	return []rune(str)
}

func value(r []rune, i int) Res {
	stack := list.New()
	preNum := 0
	res := Res{}
	for i < len(r) && string(r[i]) != ")" {
		char := string(r[i])
		if unicode.IsDigit(r[i]) { /*如果是数字则生成数字*/
			c, _ := strconv.Atoi(char)
			preNum = preNum*10 + c
			i++
		} else if char != "(" { /*不为左括号，不为右括号，也不为数字，则为加减乘除*/
			mulDiv(stack, preNum)
			stack.PushBack(char)
			preNum = 0
			i++
		} else { /*为左括号，则从左括号下一个位置开始递归*/
			res = value(r, i+1)
			preNum = res.num
			i = res.pos + 1
		}
	}
	mulDiv(stack, preNum)
	return Res{addSub(stack), i}
}

func mulDiv(stack *list.List, num int) {
	if stack.Len() != 0 {
		cur := 0
		top := stack.Remove(stack.Back())
		if top == "+" || top == "-" {
			stack.PushBack(top)
		} else { /*不是加或减，则top是乘或者除*/
			cur = stack.Remove(stack.Back()).(int) /*cur一定是数字*/
			if top == "*" {
				num = cur * num
			} else {
				num = cur / num
			}
		}
	}
	stack.PushBack(num)
}

func addSub(stack *list.List) int {
	res := 0
	preNum := 0
	for stack.Len() != 0 {
		top := stack.Remove(stack.Back())
		switch top.(type) {
		case int:
			preNum = top.(int)
		case string:
			if top.(string) == "+" {
				res = res + preNum
			} else {
				res = res - preNum
			}
		}
	}
	return res+preNum
}

func main() {
	str1 := "48*((70-65)-43)+8*1"
	//str2 := "4*(-3)"
	//str3:="3+(1*4)"
	//str:="3+1*4"
	r := strToRune(str1)
	res := value(r, 0)
	fmt.Println(res.num)
	/*r := strToRune(str)
	num := 0
	for i := 0; i < len(r); i++ {
		char := string(r[i])
		if unicode.IsDigit(r[i]) {
			c, _ := strconv.Atoi(char)
			num = num*10 + c
		} else {
			fmt.Printf("%d ", num)
			num = 0
		}
	}*/
}
