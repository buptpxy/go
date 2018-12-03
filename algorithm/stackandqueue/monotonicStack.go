/*单调栈
用来求一个数组中所有的数左边离它最近的大于它的数和右边离它最近的大于它的数，时间复杂度O(n)
例如：{3,5,2,5,6,0,1,5} =>{{3,-1,5},{5,-1,6},{2,5,5},{5,-1,6},{6,-1,-1},{0,6,1},{1,6,5},{5,6,-1}}
栈s从栈底到栈顶为单调递减，
当栈为空或者当前元素小于栈顶元素时直接入栈，
栈中每个元素的左边最大值即为它下面的元素，
否则依次弹出栈顶元素直到栈顶元素大于当前元素或栈空，弹出的这个元素的右边最大值即为使它出栈的当前元素，
若是因为数组遍历完而出站，则这个弹出的元素的右边最大值记为-1。
然后让当前元素进栈，若为栈底，则左边最大值为-1。
当有相同大小的元素时，把下标记在栈的同一位置
*/
package main

import (
	"fmt"
	"github.com/pengpeng1314/go/stack"
)

func findNearBiger(arr []int) [][]int {
	if len(arr) < 1 {
		return nil
	}
	res := make([][]int, len(arr))
	s := stack.NewStack()
	for i := 0; i < len(arr); i++ {
		res[i] = make([]int, 3)
		res[i][0] = arr[i]
		if s.IsEmpty() {
			res[i][1] = -1
			s.Push([]int{i})
		} else if arr[i] == arr[s.Top().([]int)[0]] {
			el := s.Pop().([]int)
			el = append(el, i)
			if s.IsEmpty() {
				res[i][1] = -1
			} else {
				indexs := s.Top().([]int)
				res[i][1] = arr[indexs[len(indexs)-1]]
			}
			s.Push(el)
		} else if arr[i] < arr[s.Top().([]int)[0]] {
			indexs := s.Top().([]int)
			res[i][1] = arr[indexs[len(indexs)-1]]
			s.Push([]int{i})
		} else {
			for !s.IsEmpty() && arr[i] > arr[s.Top().([]int)[0]] {
				indexs := s.Top().([]int)
				for _, v := range indexs {
					res[v][2] = arr[i]
				}
				s.Pop()
			}
			if s.IsEmpty() {
				res[i][1] = -1
				s.Push([]int{i})
			} else {
				indexs := s.Top().([]int)
				res[i][1] = arr[indexs[len(indexs)-1]]
				s.Push([]int{i})
			}
		}
	}
	for !s.IsEmpty() {
		indexs := s.Top().([]int)
		for _, v := range indexs {
			res[v][2] = -1
		}
		s.Pop()
	}
	return res
}

func findNearBiger1(arr []int) [][]int {
	if len(arr) < 1 {
		return nil
	}
	res := make([][]int, len(arr))
	s := stack.NewStack()
	//先遍历数组
	for i := 0; i < len(arr); i++ {
		res[i] = make([]int, 3)
		res[i][0] = arr[i]
		el := []int{i}
		//先考虑需要出栈的情况，即当前元素大于栈顶元素，出栈时可得到栈顶元素的右大值
		for !s.IsEmpty() && arr[i] > arr[s.Top().([]int)[0]] {
			indexs := s.Top().([]int)
			for _, v := range indexs {
				res[v][2] = arr[i]
			}
			s.Pop()
		}
		//分情况讨论左大值
		//若栈为空，则当前元素的左大值为-1
		if s.IsEmpty() {
			res[i][1] = -1
		} else {
			//栈不为空时，可分为当前元素等于栈顶元素或小于栈顶元素两种情况
			//默认情况为当前元素小于栈顶元素，当前入栈元素为i，左大值为栈顶元素
			indexs := s.Top().([]int)
			res[i][1] = arr[indexs[len(indexs)-1]]
			//当前元素等于栈顶元素时，求左大值又分两种情况，但都需栈顶元素先出栈
			if arr[i] == arr[s.Top().([]int)[0]] {
				el = s.Pop().([]int)
				el = append(el, i)
				//栈顶元素出栈后，若栈为空，则左大值为-1
				if s.IsEmpty() {
					res[i][1] = -1
				} else {
					//栈顶元素出栈后，若栈不为空，则左大值为当前栈顶元素
					indexs = s.Top().([]int) //此时的s.Top()已经不同于上面的s.Top()
					res[i][1] = arr[indexs[len(indexs)-1]]
				}
			}
		}
		//全部情况讨论完后，处理元素入栈
		s.Push(el)
	}
	//数组遍历完再遍历栈中的数，栈中的数左大值已确定，右大值都为-1
	for !s.IsEmpty() {
		indexs := s.Top().([]int)
		for _, v := range indexs {
			res[v][2] = -1
		}
		s.Pop()
	}
	return res
}

func main() {
	arr := []int{3, 5, 2, 5, 6, 0, 1, 5}
	//arr1:=[]int{3,3,3}
	res := findNearBiger1(arr)
	fmt.Println(res)
}
