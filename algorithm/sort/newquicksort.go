package main

/*
与经典快速排序的区别是，经典快速排序一遍只确定一个元素x的位置，即把数组分为了<=x,和>x的两部分两部分分别递归。
改进后把数组分为了<x,=x,>x三部分
经典快速排序存在的问题：因为总拿最后一个元素去和其它元素比较，当数组本身就为有序时，时间复杂度可达到O(n^2)，最好的时间复杂度为O(n*logn)
改进为随机快排可达到期望为O(nlogn)的时间复杂度,随机快排的额外空间复杂度为O(logn)，主要用在存储划分点的位置
*/
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func partition(arr []int, L, R int) (el, er int) {
	less := L - 1
	more := R + 1
	num := arr[R] //arr[R]会发生变化，所以先把arr[R]赋值给num
	// more := R //或者arr[R]不参与排序，因为知道它肯定是在等于区域的。但测试结果错误?
	for L < more {
		if arr[L] < num {
			arr[less+1], arr[L] = arr[L], arr[less+1]
			// fmt.Printf("arr[p] < num less:%d more:%d L:%d arr:%d \n ", less, more, L, arr)
			less++
			L++
		} else if arr[L] > num {
			arr[more-1], arr[L] = arr[L], arr[more-1]
			// fmt.Printf("arr[p] > num less:%d more:%d L:%d arr:%d \n ", less, more, L, arr)
			more--
		} else {
			// fmt.Printf("arr[p] == num less:%d more:%d L:%d arr:%d \n ", less, more, L, arr)
			L++
		}
	}
	el = less + 1
	er = more - 1
	// er = more
	return
}

func quicksort(arr []int, L, R int) {
	if L < R {
		//加入以下两行代码就是随机快排
		p := L + rand.Intn(R-L)         //生成一个在（L，R）区间的随机数p
		arr[p], arr[R] = arr[R], arr[p] //将arr[p]与arr[R]交换

		el, er := partition(arr, L, R)
		quicksort(arr, L, el-1)
		quicksort(arr, er+1, R)
	}
}
func qsort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quicksort(arr, 0, len(arr)-1)
}

func rightFunc(arr []int) {
	sort.Ints(arr)
}

func randomSlice(size int, value int) (arr []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Intn(size+1); i++ {
		v := rand.Intn(value) - rand.Intn(value)
		arr = append(arr, v)
	}
	return
}

func isEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		// fmt.Printf("len(arr1): %d,len(arr2): %d", len(arr1), len(arr2))
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			// fmt.Printf("arr1[%d]:%d  arr2[%d]:%d\n", i, arr1[i], i, arr2[i])
			return false
		}
	}
	// fmt.Println("true")
	return true
}

func check(size, value, testTimes int) {
	var arr1, arr2, arr3 []int
	var succeed bool
	for i := 0; i < testTimes; i++ {
		arr1, arr2, arr3 = []int{}, []int{}, []int{}
		succeed = true
		arr1 = randomSlice(size, value)
		// fmt.Println("arr1:", arr1)
		arr2 = append(arr2, arr1...)
		// fmt.Println("arr2:", arr2)
		arr3 = append(arr3, arr1...)
		rightFunc(arr1)
		// fmt.Println("sort: ", arr1)
		qsort(arr2)
		// fmt.Println("qsort: ", arr2)
		if !isEqual(arr1, arr2) {
			succeed = false
			break
		}
	}
	if succeed {
		fmt.Printf("%s \n", "Nice!")
	} else {
		fmt.Printf("%s \n", "Fuck!")
		fmt.Printf("出错的数组是：%d \n", arr3)
	}
}
func main() {
	// arr := []int{-39, 50, -19, 3, -39, 30}
	// qsort(arr)
	// fmt.Println(arr)
	check(10, 100, 500000)
}
