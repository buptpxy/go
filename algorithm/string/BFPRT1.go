package main

import "fmt"

func insertionSort1(arr []int, begin int, end int) []int {
	arr1 := arr
	for i := begin + 1; i < end+1; i++ {
		for j := i; j > begin; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
	return arr1
}

func getMedian1(arr []int, begin int, end int) int {
	arr1 := arr
	arr1 = insertionSort1(arr1, begin, end)
	index := (begin + end) / 2
	return arr1[index]
}

func medianOfMedians1(arr []int, begin int, end int) int {
	arr1 := arr
	num := end - begin + 1
	if num < 5 {
		return getMedian1(arr1, begin, end)
	}
	remain := num % 5
	offSet := 0
	if remain != 0 {
		offSet = 1
	}
	mArr := make([]int, num/5+offSet)
	for i := 0; i < len(mArr)-1; i++ {
		beginI := begin + i*5
		endI := beginI + 4
		mArr[i] = getMedian1(arr1, beginI, endI)
	}
	mArr[len(mArr)-1] = getMedian1(arr1, len(arr1)-remain, len(arr1)-1)
	return selects1(mArr, 0, len(mArr)-1, len(mArr)/2)
}
func partition1(arr []int, begin int, end int, pivot int) (int, int) {
	arr1 := arr
	less := begin - 1
	more := end + 1
	i := begin
	for i < more && less < more {
		if arr1[i] < pivot {
			arr1[i], arr1[less+1] = arr1[less+1], arr1[i]
			less++
			i++
		} else if arr1[i] > pivot {
			arr1[i], arr1[more-1] = arr1[more-1], arr1[i]
			more--
		} else {
			i++
		}
	}
	return less + 1, more - 1
}
func selects1(arr []int, begin int, end int, k int) int {
	arr1 := arr
	if begin == end {
		return arr1[begin]
	}
	pivot := medianOfMedians1(arr1, begin, end)
	less, more := partition1(arr1, begin, end, pivot)
	if k >= less && k <= more {
		return arr1[k]
	} else if k < less {
		return selects1(arr1, begin, less-1, k)
	} else {
		return selects1(arr1, more+1, end, k)
	}
}

func getKthByBFPRT1(arr []int, k int) int {
	arr1 := arr
	if k < 1 || k > len(arr1) {
		return -1
	}
	if len(arr1) <= 5 {
		insertionSort1(arr1, 0, len(arr1)-1)
		return arr1[k-1]
	}
	kthValue := selects1(arr1, 0, len(arr1)-1, k-1)
	return kthValue
}

func main() {
	arr := []int{4, 2, 6, 3, 2, 2}
	//insertionSort1(arr,5,5)
	//m:=getMedian1(arr,5,5)
	//m:=medianOfMedians1(arr,5,5)
	//m1,m2:=partition1(arr,5,5,2)
	//m:= selects1(arr,5,5,2)
	m := getKthByBFPRT1(arr, 4)
	fmt.Println(m)
}
