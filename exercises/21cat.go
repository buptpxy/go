/*
Q21. Cat
	1. 编写一个程序，模仿 Unix 的 cat 程序。对于不知道这个程序的人来说，下面的调用显示了文件 blah 的内容：
		% cat blah
	2. 使其支持 n 开关，用于输出每行的行号。
	go run 21cat.go --n=true
	go run 21cat.go --n=true "/Users/gaohongmin/Desktop/test/go/firstGo/src/firstGo/var.go"
	3. 上面问题中， 1 提供的解决方案存在一个 Bug。你能定位并修复它吗？
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var lineFlag = flag.Bool("n", false, "num each line")

func cat(r *bufio.Reader) {
	// ReadBytes 功能同 ReadSlice，只不过返回的是缓存的拷贝。
	//func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
	// ReadString 功能同 ReadBytes，只不过返回的是字符串。
	//func (b *Reader) ReadString(delim byte) (line string, err error)
	//cannot use "\n" (type string) as type byte in argument to r.ReadString

	line := 1
	for {
		s, err := r.ReadString('\n') //应在for循环内
		//invalid operation: s == io.EOF (mismatched types string and error)
		if err == io.EOF {
			break
		}
		if *lineFlag { //如果不加*会报错:lineFlag (type *bool) used as if condition
			fmt.Printf("%d  %s \n", line, s)
			line++
		} else {
			fmt.Printf("%s \n", s)
		}
	}
	return

}
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	flag.Parse()
	if flag.NArg() == 0 {
		cat(reader)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
			//fmt.Fprintf() 依据指定的格式向第一个参数内写入字符串，第一个参数必须实现了 io.Writer 接口。
			//Fprintf() 能够写入任何类型，只要其实现了 Write 方法，包括 os.Stdout,文件（例如 os.File），管道，网络连接，通道等等，
			//同样的也可以使用 bufio 包中缓冲写入。
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
