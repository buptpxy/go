package queue

import (
	"fmt"
)

type Element interface{}

type Queue struct {
	element []Element
}

//初始化
func NewQueue() *Queue {
	return &Queue{}
}
func (q *Queue) Size() int {
	return len(q.element)
}

func (q *Queue) Top() Element {
	if q.Size() > 0 {
		return q.element[0]
	}
	return nil
}
func (q *Queue) Push(value ...Element) {
	q.element = append(q.element, value...)
}
func (q *Queue) Pop() (top interface{}) { //返回值为interface类型即可以返回任何类型的值，这里也可写为Element类型
	if q.Size() > 0 {
		top = q.Top()
		q.element = q.element[1:q.Size()]
		return
	}
	return nil
}
func (q *Queue) IsEmpty() bool {
	if q.element == nil || q.Size() == 0 {
		return true
	}
	return false
}
func (q *Queue) Print() {
	for i := 0; i < q.Size(); i++ {
		fmt.Println(i, "=>", q.element[i])
	}
}
