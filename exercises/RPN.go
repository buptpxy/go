package main

/*
	优先级：1. + -  2. * /  3.(
	中缀转后缀：
	1.遇到操作数直接输出
	2.当栈为空，栈顶为'('，当前元素为'('，当前元素优先级比栈顶元素优先级高，这四种情况直接入栈
	3.当前元素为')'，一直出栈直到栈里面的'('
	4.当前元素不为'),且当前元素优先级比栈顶元素优先级低，一直出栈直到当前元素比栈顶元素优先级高（优先级相同也出栈）或栈为空，当前元素入栈
	5.读到输入的末尾，栈内元素依次出栈

	计算后缀表达式的值：
	将该表达式的字符以及符号，按照从左到右的顺序，依次入栈，一碰到符号则将栈顶前两个元素取出，做运算然后放入栈内，重复该操作，直到表达式结束。

	go run RPN.go --expression "33-(5+1)+2"
*/
import (
	"errors"
	"flag"
	"fmt"
	"stack"
	"strconv"
)

//判断string里面的每个字符是否是数字 ，需要遍历字符串中的每个字符时，可以把string转换成[]byte 或 []rune
func byteIsNumber(op byte) bool {
	if op == '+' || op == '-' || op == '*' || op == '/' || op == '(' || op == ')' {
		return false
	}
	return true
}

//当一个 if 语句不会进入下一个语句流程 – 也就是说，语句体结束于 break， continue， goto 或者 return – 不必要的 else 会被省略。
//把输入的无空格的string变成[]string
func convertToStrings(str string) (strs []string) {
	bys := []byte(str)
	var tmp string //存储数字
	for _, b := range bys {
		if !byteIsNumber(b) {
			if tmp != "" {
				strs = append(strs, tmp)
				tmp = ""
			}
			strs = append(strs, string(b))
		} else {
			tmp = tmp + string(b)
		}
	}
	strs = append(strs, tmp)
	return
}

//设定操作符优先级
func prefer(op1 string, op2 string) bool {
	priority := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "(": 3}
	if priority[op1] > priority[op2] {
		return true
	}
	return false

}

//判断[]string 里面的每个字符串是否为数字
func stringIsNumber(op string) bool {
	if op == "+" || op == "-" || op == "*" || op == "/" || op == "(" || op == ")" {
		return false
	}
	return true
}

//中缀转后缀
func generateRPN(s1 []string) (s2 []string) {
	st := stack.NewStack()
	for _, v := range s1 {
		switch {
		//1.数字直接添加到s2
		case stringIsNumber(v):
			s2 = append(s2, v)
			//2.四种情况直接入栈 ,注意st.IsEmpty()要再prefer(v, st.Top().(string))之前
		case st.IsEmpty(), prefer(v, st.Top().(string)), v == "(", st.Top().(string) == "(":
			st.Push(v)
			//3.当前元素为')'，一直出栈直到栈里面的'('
		case v == ")":
			for !st.IsEmpty() && st.Top().(string) != "(" {

				s2 = append(s2, st.Pop().(string)) //出栈并添加到s2中
			}
			st.Pop() //"("出栈
			//4.当前元素不为'),且当前元素优先级比栈顶元素优先级低，一直出栈直到当前元素比栈顶元素优先级高（优先级相同也出栈）或栈为空
		case v != ")" && !prefer(v, st.Top().(string)): //不能写成prefer(st.Top().(string),v )
			for prefer(st.Top().(string), v) && !st.IsEmpty() {
				//不能再加一句st.Pop()，加了就出栈两次了
				s2 = append(s2, st.Pop().(string))
			}
			st.Push(v) //当前元素入栈
		}
	}
	//5. 读到输入的末尾，栈内元素依次出栈
	for !st.IsEmpty() {

		s2 = append(s2, st.Pop().(string))
	}
	return
}

//定义四则运算
func operate(num1 float64, num2 float64, op string) (answer float64, err error) {
	switch op {
	case "+":
		answer = num1 + num2
	case "-":
		answer = num1 - num2
	case "*":
		answer = num1 * num2
	case "/":
		if num2 != 0 {
			answer = num1 / num2
		} else {
			err = errors.New("math: divide by 0")
		}
		//也可用 return 0, fmt.Errorf("math: divide of negative number %g", num2)
		// return 0, errors.New("math: divide by 0")
	}
	return
}

//计算后缀表达式
func calculate(strs []string) (result float64, err error) {
	st := stack.NewStack()
	var f float64
	for _, str := range strs {
		if stringIsNumber(str) {
			f, err = strconv.ParseFloat(str, 64)
			st.Push(f)
		} else {
			num2 := st.Pop().(float64) //要加.(float64)，不然会报错cannot use num1 (type interface {}) as type float64 in argument to operate: need type assertion
			// fmt.Printf("num2为：%f \n", num2)
			num1 := st.Pop().(float64)
			// fmt.Printf("num1为：%f \n", num1)
			result, err = operate(num1, num2, str)
			// fmt.Printf("后缀表达式的计算结果为：%f \n", result)
			st.Push(result)
		}
	}
	return
}

func main() {
	var exp *string = flag.String("expression", "", "") //第一个参数为参数名称，第二个参数为默认值，第三个参数是说明
	//解析函数将会在碰到第一个非flag命令行参数时停止，非flag命令行参数是指不满足命令行语法的参数，如命令行参数为cmd --flag=true abc则第一个非flag命令行参数为“abc”
	flag.Parse()
	fmt.Println("表达式为：", *exp)
	infix := convertToStrings(*exp)
	fmt.Printf("中缀表达式为：%s \n", infix)
	fmt.Printf("优先级为：%t \n", prefer("*", "("))
	suffix := generateRPN(infix)
	fmt.Printf("后缀表达式为：%s \n", suffix)
	answer, err := calculate(suffix)
	if err == nil {
		fmt.Printf("后缀表达式的计算结果为：%f \n", answer)
	} else {
		fmt.Println("error: ", err)
	}

}
