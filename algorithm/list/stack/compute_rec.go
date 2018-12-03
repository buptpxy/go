package main

import (
	"container/list"
	"fmt"
	"strconv"
	"unicode"
)

func strToRune(str string) []rune{
	return []rune(str)
}

func doMulDiv(stack *list.List,num int)  {
	if stack.Len()!=0  {
		top:=stack.Remove(stack.Back()).(string)
		if top=="+" || top=="-"{
			stack.PushBack(top)
		}else {
			value:=stack.Remove(stack.Back()).(int)
			if top=="*" {
				num=value*num
			}else {
				num=value/num
			}
		}
	}
	stack.PushBack(num)
}

func doAddSub(stack *list.List) int {
	res:=0
	preNum:=0
	for stack.Len()!=0  {
		top:=stack.Remove(stack.Back())
		switch top.(type) {
		case int:
			preNum=top.(int)
		case string:
			if top.(string)=="+" {
				res=res+preNum
			}else {
				res=res-preNum
			}
		}
	}
	return res+preNum
}

func compute(r []rune,i int) (int,int) {
	stack:=list.New()
	preNum:=0
	for i<len(r) && string(r[i])!=")"  {
		char:=string(r[i])
		if unicode.IsDigit(r[i]) {
			n,_:=strconv.Atoi(char)
			preNum=preNum*10+n
			i++
		}else if char!="(" {
			doMulDiv(stack,preNum)/*注意先让num入栈*/
			stack.PushBack(char)/*再让符号入栈*/
			preNum=0
			i++
		}else {
			res,pos:=compute(r,i+1)
			preNum=res
			i=pos+1
		}
	}
	doMulDiv(stack,preNum)/*注意此时preNum还没入栈，故还需调用doMulDiv*/
	resN:=doAddSub(stack)
	return resN,i
}

func main() {
	//str1 := "48*((70-65)-43)+8*1"
	str2 := "4*(-3)"
	//str3:="3+(1*4)"
	//str:="3+1*4"
	r := strToRune(str2)
	res,_ := compute(r, 0)
	fmt.Println(res)
}