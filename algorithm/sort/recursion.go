package main

import (
	"fmt"
	"math"
)

func getMax(arr []int, L int, R int) int {
	if L == R {
		return arr[L]
	}
	mid := (L + R) / 2
	leftMax := getMax(arr, L, mid)
	fmt.Println("leftMax: ", leftMax)
	rightMax := getMax(arr, mid+1, R)
	fmt.Println("rightMax: ", rightMax)
	max := math.Max(float64(leftMax), float64(rightMax))
	fmt.Println("max: ", max)
	return int(max)
}
func main() {
	arr := []int{2, 4, 3, 1}
	max := getMax(arr, 0, len(arr)-1)
	fmt.Println("max: ", max)
}
