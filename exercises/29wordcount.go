package main

/*
29. 单词和字母统计
	编写一个从标准输入中读取文本的小程序，并进行下面的操作：
	1. 计算字符数量（包括空格）；
	2. 计算单词数量；
	3. 计算行数。
	换句话说，实现一个 wc(1)（参阅本地的手册页面），然而只需要从标准输入读取。
*/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var chars, words, lines int
	buf := bufio.NewReader(os.Stdin)
	// fmt.Println(buf.ReadString('\n'))
	// s, err := buf.ReadString('\n')
	// fmt.Println("s: ", s)
	// fmt.Println("err: ", err)
	for {
		switch s, ok := buf.ReadString('\n'); true {
		case s == "eof\n", ok != nil:
			fmt.Printf("字符数：%d 单词数：%d 行数：%d \n", chars, words, lines)
			return
		default:
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
