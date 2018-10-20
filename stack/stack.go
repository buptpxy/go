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
	"fmt"
)

type Element interface{}

type Stack struct {
	element []Element
}

//初始化
func NewStack() *Stack {
	return &Stack{}
}
func (s *Stack) Size() int {
	return len(s.element)
}

func (s *Stack) Top() Element {
	if s.Size() > 0 {
		return s.element[s.Size()-1]
	}
	return nil
}
func (s *Stack) Push(value ...Element) {
	s.element = append(s.element, value...)
}
func (s *Stack) Pop() (top interface{}) { //返回值为interface类型即可以返回任何类型的值，这里也可写为Element类型
	if s.Size() > 0 {
		top = s.Top()
		s.element = s.element[:s.Size()-1]
		return
	}
	return nil
}
func (s *Stack) IsEmpty() bool {
	if s.element == nil || s.Size() == 0 {
		return true
	}
	return false
}
func (s *Stack) Print() {
	for i := s.Size() - 1; i >= 0; i-- {
		fmt.Println(i, "=>", s.element[i])
	}
}
