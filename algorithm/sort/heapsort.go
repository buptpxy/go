/*
每一个数组都可以看成一个完全二叉树（满二叉树是完全二叉树的一种）
位置i上的数的左孩子为arr[2*i+1]，右孩子为arr[2*i+2]，父节点为arr[(i-1)/2]，就算i为0，(i-1)/2也为零，即父节点是其本身
堆分为大根堆和小根堆，大根堆即为每个子树的根节点一定是最大的。
建立一个大根堆的时间复杂度为O(n)
堆排序思路：
	1. 先把数组调整为一个大根堆：把当前位置上的数与其父节点比较，如果比父节点大则替换父节点
	2. 此大根堆的根节点是数组中最大的数，让它和数组最后一个数交换，并把堆的大小减1
	3. 把新的根节点与其左右孩子比较，小的话就往下沉
	4. 重复2,3过程，直到堆的大小为0
*/
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func getMaxHeap(arr []int) {
	for i := 0; i < len(arr); i++ {
		index := i
		for arr[index] > arr[(index-1)/2] {
			arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
			index = (index - 1) / 2
		}
		// fmt.Println(arr)
	}
}

//index位置上的数的下沉过程
func heapify(arr []int, index int, size int) {
	// size := len(arr)
	left := index*2 + 1

	for left < size {
		var largest int
		if left+1 < size && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if arr[index] < arr[largest] {
			arr[index], arr[largest] = arr[largest], arr[index]
			index = largest
		} else {
			break
		}
		left = index*2 + 1
	}

}
func heapSort(arr []int) {
	size := len(arr)
	if size < 2 {
		return
	}
	getMaxHeap(arr)
	arr[size-1], arr[0] = arr[0], arr[size-1]
	size--
	for size > 0 {
		heapify(arr, 0, size)
		arr[size-1], arr[0] = arr[0], arr[size-1]
		size--
	}
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
		heapSort(arr2)
		// fmt.Println("heapSort: ", arr2)
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

	// arr := []int{1, 4, 2, 3, 5, 7, 3}
	// arr1 := arr[0 : len(arr)-1]
	// fmt.Println(arr1)
	// getMaxHeap(arr)
	// heapSort(arr)
	// fmt.Println(arr)
	check(10, 100, 500000)
}
