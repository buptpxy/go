/*
定义数组的异或和的概念：
异或：相同为0，相异为1。0异或任何数都为那个数本身。0异或0为0。
数组中所有的数异或起来， 得到的结果叫做数组的异或和，比如数组{3,2,1}的异或和是， 3^2^1 = 0
给定一个数组arr， 你可以任意把arr分成很多不相容的子数组， 你的目的是：
	分出来的子数组中，异或和为0的子数组最多。
	请返回：分出来的子数组中，异或和为0的子数组最多是多少？
【思路】求这种子数组数量的题，可使用动态规划。
0到i位置的数组的异或0子数组数量与0到i前面位置的数组有关。
当i所在的子数组不属于最优划分子数组时，它的最优子数组数量与0到i-1位置的最优子数组数量相同。
当i所在的子数组属于最优划分子数组时，先求出离i最近的异或和为0的位置k，它的最优子数组数量等于0到k-1位置的最优子数组数量加1。
遍历arr，从0到每个位置i的异或和xor和i都被记录在map中。
所以dp数组为一维数组，记录当前位置i和0到i的最优划分子数组数量。
如果xor为0上次出现的位置为-1，则dp[0]=1
如果xor为0上次出现的位置不为-1，则记此位置为k，dp[i]=dp[k]+1
当i>0，dp[i]=max(dp[i-1],dp[i])
maxLen=max(maxLen,dp[i])
*/

package main

import "fmt"

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func mostSum0XorSubArray(arr []int) int {
	xorPos := make(map[int]int)
	xorPos[0] = -1
	xor := 0
	maxLen := 0
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		xor = xor ^ arr[i]
		if pos, ok := xorPos[xor]; ok {
			if pos == -1 {
				dp[i] = 1
			} else {
				pre := pos
				dp[i] = dp[pre] + 1
			}
		}
		if i > 0 {
			dp[i] = max(dp[i-1], dp[i])
		}
		xorPos[xor] = i
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}
func main() {
	arr := []int{3, 2, 1, 2, 2}
	fmt.Println(mostSum0XorSubArray(arr))
}
