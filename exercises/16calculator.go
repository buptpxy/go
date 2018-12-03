package main

/*
Q16. 创建逆波兰（后缀）计算器计算给定的后缀表达式。
	解题思路：
	1.逆波兰表达式可以讲复杂的计算过程转化为简单的操作过程，进而得出答案。 比如 (a+b)*(b-c) 按照逆波兰表达式的规则得到 ：ab+bc-*
	2.然后将该表达式的字符以及符号，按照从左到右的顺序，依次入栈，一碰到符号则将栈顶前两个元素取出，做运算然后放入栈内，重复该操作，直到表达式结束。
	下面将结合栈与逆波兰表达式写一个简易计算器。
*/
import (
	"bufio"
	"fmt"
	stack "github.com/pengpeng1314/go/stack"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		st := stack.NewStack()
		var reader *bufio.Reader = bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		// var token string
		if err != nil {
			return
		}
		for _, c := range strings.Fields(s) {

			switch {
			case c >= "0" && c <= "9":
				r, _ := strconv.Atoi(c)
				// fmt.Printf("r: %d\n", r)
				st.Push(r)
				// st.Print()

			case c == "+":
				p := st.Pop().(int) //将interface{}转换为int
				q := st.Pop().(int)
				answer := p + q
				st.Push(answer)
				fmt.Printf("%d\n", answer)

			case c == "*":
				p := st.Pop().(int)
				q := st.Pop().(int)
				answer := p * q
				st.Push(answer)
				fmt.Printf("%d\n", answer)

			case c == "-":
				p := st.Pop().(int)
				q := st.Pop().(int)
				answer := q - p
				st.Push(answer)
				fmt.Printf("%d\n", answer)

			case c == "exit":
				return
			default:
				//error
			}
		}

	}
}
