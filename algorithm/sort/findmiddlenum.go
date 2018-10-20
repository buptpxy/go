/*
堆的应用：
	有一个源头一直在往外输出数字，要求实时求出当前所有已输出数字的中位数。
	中位数是一串数中大小排在中间的数，若共有奇数个数字，中位数只有一个；若共有偶数个数字，中位数定义为大小排在中间的两个数的平均值。
思路：
	建立一个大根堆一个小根堆，把较小的N/2个数放进大根堆，较大的N/2个数放进小根堆。
	这样就可保证size大的那个堆的根即为中位数，size一样大的话，即取两个堆的根的平均数
具体操作：
	1. 把第一个数放进大根堆，
	2. 如果接下来的数小则放进大根堆，否则放进小根堆。
	3. 当两个堆size差异大于2时，把size大的那个堆的堆顶元素取出放入另一个堆。
	4. 重复2，3直到n个元素放完
*/
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// type heap interface {
// 	Pop() int
// 	Push(int)
// }

type Maxheap []int

func (maxh Maxheap) Push(num int) []int {
	//插入到堆尾
	maxh = append(maxh, num)
	//调整过程
	index := len(maxh) - 1
	for maxh[index] > maxh[(index-1)/2] {
		maxh[index], maxh[(index-1)/2] = maxh[(index-1)/2], maxh[index]
		index = (index - 1) / 2
	}
	return maxh
}
func (maxh Maxheap) Pop() (newmaxh []int, root int) {
	//首尾交换
	size := len(maxh)
	if size < 1 {
		return
	}
	root = maxh[0]
	// fmt.Printf("root: %d\n", root)
	maxh[0], maxh[size-1] = maxh[size-1], maxh[0]
	size--
	//调整过程
	index := 0
	left := index*2 + 1
	for left < size {
		var largest int
		if left+1 < size && maxh[left+1] > maxh[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if maxh[largest] > maxh[index] {
			maxh[largest], maxh[index] = maxh[index], maxh[largest]
			index = largest
		} else {
			break
		}
		left = index*2 + 1
	}
	newmaxh = maxh[0:size]

	return
}

type Minheap []int

func (minh Minheap) Push(num int) []int {
	//插入到堆尾
	minh = append(minh, num)
	//调整过程
	index := len(minh) - 1
	for minh[index] < minh[(index-1)/2] {
		minh[index], minh[(index-1)/2] = minh[(index-1)/2], minh[index]
		index = (index - 1) / 2
	}
	return minh
}
func (minh Minheap) Pop() (newminh []int, root int) {
	//首尾交换
	size := len(minh)
	if size < 1 {
		return
	}
	root = minh[0]
	// fmt.Printf("root: %d\n", root)
	minh[0], minh[size-1] = minh[size-1], minh[0]
	size--
	//调整过程
	index := 0
	left := index*2 + 1
	for left < size {
		var least int
		if left+1 < size && minh[left+1] < minh[left] {
			least = left + 1
		} else {
			least = left
		}
		if minh[least] < minh[index] {
			minh[least], minh[index] = minh[index], minh[least]
			index = least
		} else {
			break
		}
		left = index*2 + 1
	}
	newminh = minh[0:size]

	return
}

func middlenum(maxh Maxheap, minh Minheap, num int) (Maxheap, Minheap, int) {

	rootmax := maxh[0]
	var rootmin int
	var middle int
	if num < rootmax {
		maxh = maxh.Push(num)
		// fmt.Printf("num < rootmax: maxh:%d  minh:%d \n", maxh, minh)

	} else {
		minh = minh.Push(num)
		// fmt.Printf("num >= rootmax: maxh:%d  minh:%d \n", maxh, minh)

	}
	if len(maxh)-len(minh) == 2 {
		maxh, rootmax = maxh.Pop()
		minh = minh.Push(rootmax)
		// fmt.Printf("len(maxh)-len(minh) == 2: maxh:%d  minh:%d \n", maxh, minh)

	}
	if len(minh)-len(maxh) == 2 {
		minh, rootmin = minh.Pop()
		maxh = maxh.Push(rootmin)
		// fmt.Printf("len(minh)-len(maxh) == 2: maxh:%d  minh:%d \n", maxh, minh)

	}
	if len(maxh) == len(minh) {
		middle = (maxh[0] + minh[0]) / 2
		// fmt.Printf("len(maxh) == len(minh): maxh:%d  minh:%d middle:%d\n", maxh, minh, middle)
	} else if len(maxh) > len(minh) {
		middle = maxh[0]
		// fmt.Printf("len(maxh) > len(minh): maxh:%d  minh:%d middle:%d\n", maxh, minh, middle)
	} else {
		middle = minh[0]
		// fmt.Printf("len(maxh) < len(minh): maxh:%d  minh:%d middle:%d\n", maxh, minh, middle)
	}
	// fmt.Printf("in maxh:%d minh:%d middle:%d\n", maxh, minh, middle)
	return maxh, minh, middle
}

func findmiddlenum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	//注意考虑这种情况
	if len(arr) == 1 {
		return arr[0]
	}
	var maxh Maxheap
	var minh Minheap
	var middle int
	maxh = maxh.Push(arr[0])

	for i := 1; i < len(arr); i++ {
		maxh, minh, middle = middlenum(maxh, minh, arr[i]) //maxh, minh, middle不能在for循环里面定义，否则每次循环都会初始化
	}
	return middle
}
func rightFunc(arr []int) int {
	var middle int
	if len(arr) == 0 {
		return 0
	}
	if len(arr) > 2 {
		sort.Ints(arr)
	}
	if len(arr)%2 == 0 {
		middle = (arr[(len(arr)-1)/2] + arr[len(arr)/2]) / 2
	} else {
		middle = arr[len(arr)/2]
	}
	return middle
}

func randomSlice(size int, value int) (arr []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Intn(size+1); i++ {
		v := rand.Intn(value) - rand.Intn(value)
		arr = append(arr, v)
	}
	return
}

func isEqual(num1, num2 int) bool {
	if num1 == num2 {
		return true
	} else {
		return false
	}

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
		num1 := rightFunc(arr1)
		// fmt.Println("sort: ", arr1)
		num2 := findmiddlenum(arr2)
		// fmt.Println("heapSort: ", arr2)
		if !isEqual(num1, num2) {
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

//函数里包含对slice进行append()操作时，一定要返回操作后的slice，因为append()相当于是新生成了一个slice
func addtest(arr []int, num int) ([]int, []int) {
	arr = append(arr, num)
	arr1 := append(arr, num+1)
	return arr, arr1
}

//因此要对slice进行append操作时应传入slice的指针
func addtest1(arr *[]int, num int) {
	*arr = append(*arr, num)
}
func main() {
	// arr := []int{-26}
	// addtest1(&arr, 1)
	// fmt.Println(arr)
	// fmt.Println("findmiddlenum: ", findmiddlenum(arr))
	// fmt.Println("rightFunc:", rightFunc(arr))
	check(10, 100, 500000)

}
