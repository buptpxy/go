/*
给定一个数组， 求如果排序之后， 相邻两数的最大差值， 要求时间复杂度O(N)， 且要求不能用非基于比较的排序。
思路：
	借鉴桶排序的思想，有N个数则使用N+1个桶。
	首先变量数组找出数组的最大值和最小值，然后把两者的差值分成N+1个范围（N+1个桶）。
	接下来再遍历一次数组，把每个数放到相应的桶中。最小的数在0号桶，最大的数在N号桶。
	中间必定至少有一个桶是空的，所以最大的gap一定不在一个桶内
	所以应记录下每个桶的最小值和离它最近的左边的非空桶的最大值的差值，最大的gap一定就存在这些差值中
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func bucket(length int, num, min, max float64) int {
	return int((num - min) * float64(length) / (max - min))
}

func maxGap(arr []float64) float64 {
	if len(arr) < 2 {
		return 0
	}
	min := arr[0]
	max := arr[0]
	length := len(arr)
	//找出数组的最大最小值
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	if min == max {
		return 0
	}
	//把数组中的数放到对应的桶里，并求出每个桶的最大最小值
	// var hasnum []bool //默认都为false,这样声明数组默认大小为0
	// var mins []float64
	// var maxs []float64
	hasnum := make([]bool, length+1)
	mins := make([]float64, length+1)
	maxs := make([]float64, length+1)
	bid := 0 //桶的标号
	for _, num := range arr {
		bid = bucket(length, num, min, max)
		if hasnum[bid] {
			mins[bid] = math.Min(mins[bid], num)
			maxs[bid] = math.Max(maxs[bid], num)
		} else {
			mins[bid] = num
			maxs[bid] = num
			hasnum[bid] = true
		}
	}
	//记录每个桶的最小值和离它最近的左边的非空桶的最大值的差值
	var maxgap float64
	lastMax := maxs[0]
	for k, v := range hasnum {
		if v {
			maxgap = math.Max(maxgap, mins[k]-lastMax)
			lastMax = maxs[k]
		}
	}
	return maxgap
}

func rightFunc(arr []float64) float64 {
	if len(arr) < 2 {
		return 0
	}
	sort.Float64s(arr)
	last := arr[0]
	var maxgap float64
	for _, v := range arr {
		maxgap = math.Max(maxgap, v-last)
		last = v
	}
	return maxgap
}

func randomSlice(size int, value int) (arr []float64) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Intn(size+1); i++ {
		v := rand.Float64()*float64(value) - rand.Float64()*float64(value)
		arr = append(arr, v)
	}
	return
}

func isEqual(num1, num2 float64) bool {
	if num1 == num2 {
		return true
	} else {
		return false
	}
}

func check(size, value, testTimes int) {
	var arr1, arr2, arr3 []float64
	var succeed bool
	for i := 0; i < testTimes; i++ {
		arr1, arr2, arr3 = []float64{}, []float64{}, []float64{}
		succeed = true
		arr1 = randomSlice(size, value)
		// fmt.Println("arr1:", arr1)
		arr2 = arr1
		// fmt.Println("arr2:", arr2)
		arr3 = arr1
		num1 := rightFunc(arr1)
		// fmt.Println("sort: ", arr1)
		num2 := maxGap(arr2)
		// fmt.Println("bubblesort: ", arr2)
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

func main() {

	// arr := []float64{1, 4}
	// maxgap := maxGap(arr)
	// fmt.Println(maxgap)
	check(10, 100, 500000)

}
