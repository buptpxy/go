package main

import "fmt"

/*
BFPRT算法及其复杂度估计。
在一堆数中求其前k大或前k小的问题，简称TOP-K问题。而目前解决TOP-K问题最有效的算法即是BFPRT算法，又称为中位数的中位数算法。
在快速排序的基础上，首先通过判断主元位置与k的大小使递归的规模变小，其次通过修改快速排序中主元的选取方法来降低快速排序在最坏情况下的时间复杂度。
BFPRT算法步骤如下：

选取主元；
1.1. 将n个元素按顺序分为⌊n/5⌋个组，每组5个元素，若有剩余，再分一个组；
1.2. 对于这⌊n/5⌋个组中的每一组使用插入排序找到它们各自的中位数；
1.3. 对于 1.2 中找到的所有中位数，调用BFPRT算法求出它们的中位数，作为主元；
以 1.3 选取的主元为分界点，把小于主元的放在左边，大于主元的放在右边；
判断主元的位置与k的大小，有选择的对左边或右边递归。
*/

//不要直接对传进来的slice进行操作，会直接改变这个slice的值
func insertionSort(arr []int, begin int, end int) []int {
	arr1 := arr
	for i := begin + 1; i < end+1; i++ {
		for j := i; j > begin; j-- {
			if arr1[j-1] > arr1[j] {
				arr1[j-1], arr1[j] = arr1[j], arr1[j-1]
			} else {
				break
			}
		}
	}
	return arr1
}
func getMedian(arr []int, begin int, end int) int {
	arr1 := arr
	arr1 = insertionSort(arr1, begin, end)
	index := (begin+end)/2 + (begin+end)%2 //+(begin+end)%2 ???
	return arr1[index]
}

//求整个数组的中位数即中位数数组mArr求第len(mArr)/2小的数
func medianOfMedians(arr []int, begin int, end int) int {
	arr1 := arr
	num := (end - begin) + 1
	if num < 5 {
		return getMedian(arr1, begin, end)
	}
	remain := num % 5
	offset := 0
	if remain != 0 {
		offset = 1
	}
	mArr := make([]int, num/5+offset)
	for i := 0; i < len(mArr)-1; i++ {
		beginI := begin + i*5
		endI := beginI + 4
		mArr[i] = getMedian(arr1, beginI, endI)
	}
	mArr[len(mArr)-1] = getMedian(arr1, len(arr1)-remain, len(arr1)-1)
	return selects(mArr, 0, len(mArr)-1, len(mArr)/2)
}

func partition(arr []int, begin int, end int, pivot int) (int, int) {
	arr1 := arr
	equal1 := begin - 1
	equal2 := end + 1
	i := begin
	for i < equal2 && equal1 <= equal2 {
		if arr1[i] < pivot {
			arr1[i], arr1[equal1+1] = arr1[equal1+1], arr1[i]
			equal1++
			i++
		} else if arr1[i] > pivot {
			arr1[i], arr1[equal2-1] = arr1[equal2-1], arr1[i]
			equal2--
		} else {
			i++
		}
	}
	return equal1 + 1, equal2 - 1
}

//求数组中第k+1小的数
func selects(arr []int, begin int, end int, k int) int {
	arr1 := arr
	if begin == end {
		return arr1[begin]
	}
	pivot := medianOfMedians(arr1, begin, end)
	equal1, equal2 := partition(arr1, begin, end, pivot)
	if k >= equal1 && k <= equal2 {
		return arr1[k]
	} else if k < equal1 {
		return selects(arr1, begin, equal1-1, k)
	} else {
		return selects(arr1, equal2+1, end, k)
	}
}

//求数组中第k小的数
func getKthByBFPRT(arr []int, k int) int {
	arr1 := arr
	if k < 1 || len(arr1) < k {
		return -1
	}
	if len(arr1) < 5 {
		arr1 = insertionSort(arr1, 0, len(arr1)-1)
		return arr1[k-1]
	}
	kthValue := selects(arr1, 0, len(arr1)-1, k-1)
	return kthValue //不能把selects(arr1, 0, len(arr1)-1, k-1)直接赋在return中
}

//求数组中最小的k个数
func getTopkNumsByBFPRT(arr []int, k int) []int {
	arr1 := arr
	if k < 1 || len(arr1) < k {
		return nil
	}
	topK := make([]int, k)
	if len(arr1) < 5 {
		arr1 = insertionSort(arr1, 0, len(arr1)-1)
		for i := 0; i < k; i++ {
			topK[i] = arr1[i]
		}
		return topK
	}
	kValue := getKthByBFPRT(arr1, k)
	j := 0
	for i := 0; i < len(arr1) && j < k; i++ {
		if arr1[i] <= kValue {
			topK[j] = arr1[i]
			j++
		}
	}
	return topK
}
func main() {
	arr := []int{4, 2, 6, 3, 2, 2}
	//arr1 := insertionSort(arr, 0, 2)
	//fmt.Println(arr1)
	//arr2:= []int{1,2,3,4,5}
	//e1, e2 := partition(arr2, 0, 4, 3)
	//fmt.Printf("%d %d\n", e1, e2)
	fmt.Println(getKthByBFPRT(arr, 4))
	//fmt.Println(getTopkNumsByBFPRT(arr, 3)) //getKthByBFPRT和getTopkNumsByBFPRT不能同时都运行
	//fmt.Println(selects(arr,0,6,3))
}
