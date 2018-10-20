package stack

/*
	为了让 go test 能够工作，需要将包所在文件放到 $GOPATH/src：
	% mkdir $GOPATH/src/stack
	% cp pushpop_test.go $GOPATH/src/stack
	% cp stack-as-package.go $GOPATH/src/stack
	输出：
	% go test stack
	ok stack 0.001s
*/

import (
	"strconv"
)

type Stack struct {
	top int //元素即将被存入的下标位置
	arr [10]int
}

func (s *Stack) Push(d int) {
	if s.top > 9 {
		return
	}
	s.arr[s.top] = d
	s.top++
}
func (s *Stack) Pop() (d int) {
	s.top--
	if s.top < 0 {
		s.top = 0
		return
	}
	d = s.arr[s.top]
	return
}
func (s *Stack) String() (str string) {

	for n := 0; n < s.top; n++ {
		str = str + "[" + strconv.Itoa(n) + ":" + strconv.Itoa(s.arr[n]) + "] "
	}
	return
}
