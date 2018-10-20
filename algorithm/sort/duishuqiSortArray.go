package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func bubblesort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
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
		bubblesort(arr2)
		// fmt.Println("bubblesort: ", arr2)
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
	// s := randomSlice(10, 100)
	// fmt.Printf("%d\n", s)
	check(10, 100, 500000)
}
