/*
一个数组表示一组环形山，数组元素越大山越高。相邻的山可以互相看见，如果两座山不相邻，但是中间的山高度不大于这两座山的高度最小值，这两座山也可以互相看见。问能够互相看见的山的对数。
1、如果数组中无等高的山，则
	1座山时，0对山能互相看见
	2座山时，1对山能互相看见
	n座山时（n>2），2*n-3对山能互相看见：
					选出最高和次高的山，为了避免重复，我们规定由低的山去找高的山，
					则其他的每一座山一定跟自己左边第一座比自己高的山，和右边第一座比自己高的山组成对，故一共(n-2)*2=2*n-4对，
					加上最高和次高这一对，故一共2*n-3对。
2、如果数组中有等高的山，则按照栈底最大的单调栈的方法来做。
	首先选出数组中最大的数，从这个数开始遍历数组，保证这个数一定在栈底。
	栈中记录的是这个数的大小和它出现的次数，当一个数被弹出栈时结算它的对数，比如3出现k次，则3的对数为C(k,2)+2*k
	当数组遍历完了，但是栈中还有数时，如果这个数出栈后，栈中还剩余多于一个数，则它的对数也为C(k,2)+2*k
				如果这个数出栈后，栈中只剩下一个数，就要剩下的这个数出现次数是否大于1，大于1则同上，等于1则它的对数为C(k,2)+k
				如果这个数出栈后栈为空，则这个数的对数为C(k,2)
*/
package main

import (
	"fmt"
	"github.com/pengpeng1314/go/stack"
)

func next(arr []int, i int) int {
	if i == len(arr)-1 {
		return 0
	}
	return i + 1
}

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}
func C(n, k int) int {
	if k > n || k == 0 {
		return 0
	}
	upside := factorial(n)
	below := factorial(n-k) * factorial(k)
	result := upside / below
	return result
}
func getMaxIndex(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return max
}

type mNum struct {
	value int
	times int
}

func communications(arr []int) int {
	n := len(arr)
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	//mountainNum:=make(map[int]int)
	s := stack.NewStack()
	index := getMaxIndex(arr)
	m := mNum{arr[index], 1}
	s.Push(m)
	res := 0
	for i := next(arr, index); i != index; i = next(arr, i) {
		for !s.IsEmpty() && arr[i] > s.Top().(mNum).value {
			m = s.Pop().(mNum)
			res = res + C(m.times, 2) + m.times*2
			//最大值肯定还在栈底，故此时栈肯定不为空
		}
		m = mNum{arr[i], 1}
		if !s.IsEmpty() && arr[i] == s.Top().(mNum).value {
			m = s.Pop().(mNum)
			m.times++
		}
		s.Push(m)
	}
	for !s.IsEmpty() {
		m = s.Pop().(mNum)
		res = res + C(m.times, 2)
		if !s.IsEmpty() {
			if s.Size() > 2 {
				res = res + m.times*2
			} else {
				m0 := s.Pop().(mNum)
				if m0.times == 1 {
					res = res + m.times
				} else {
					res = res + m.times*2
				}
			}
		}
	}
	return res
}
func main() {
	fmt.Println(factorial(5))
	fmt.Println(C(5, 2))
	arr := []int{1, 2, 4, 4, 4, 5}
	fmt.Println(communications(arr))
}
