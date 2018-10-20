package main

/*
Q3. 字符串
	1. 建立一个 Go 程序打印下面的内容（到 100 个字符）：
		A
		AA
		AAA
		AAAA
		AAAAA
		AAAAAA
		AAAAAAA
		...
	2. 建立一个程序统计字符串里的字符数量：
		asSASA ddd dsjkdsjs dk
		同时输出这个字符串的字节数。 提示： 看看 unicode/utf8 包。
	3. 扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”。
	4. 编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”。 提示：你需要知道一些关于转换的内容
*/
import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//*************************1.错误的解法：cannot convert slice (type []string) to type string
	// slice := []string{}
	// for i := 0; i < 100; i++ {
	// 	slice = append(slice, "A")
	// 	str := string(slice)
	// 	fmt.Printf("%s\n", str)
	// }

	//**************************1.正确的解法：
	// var str string
	// for i := 0; i < 100; i++ {
	// 	str += "A"
	// 	fmt.Printf("%s\n", str)
	// }

	//**************************2.
	str1 := "hwlwljoljs你好 w"
	fmt.Printf("这个字符串的字节数为： %d\n", len(str1)) //在 Go 中，字符串是以 UTF-8 为格式进行存储的，在字符串上调用 len 函数，取得的是字符串包含的 byte(uint8) 的个数。英文字母是一个byte,中文是3个byte
	byteslice := []byte(str1)
	fmt.Printf("这个字符串的字数为： %d\n", utf8.RuneCount(byteslice)) //返回rune的个数,错误和短的编码被视为宽度为1个rune(int32)
	runeslice := []rune(str1)
	fmt.Printf("这个字符串的字数为：%d\n", len(runeslice))

	//**************************3.
	copy(byteslice[3:6], []byte("abc"))
	fmt.Printf("修改后的字符串为：%s \n", byteslice)

	//**************************4.
	str2 := "foobar"
	runeslice2 := []rune(str2)
	for i, j := 0, len(runeslice2)-1; i < j; i, j = i+1, j-1 { //不能写i++,j--,会报错syntax error: unexpected comma, expecting { after for clause
		runeslice2[i], runeslice2[j] = runeslice2[j], runeslice2[i]
	}
	fmt.Printf("逆置后的字符串为：%s \n", string(runeslice2))
}
