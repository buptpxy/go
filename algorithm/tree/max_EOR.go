/*
给定一个数组， 求子数组的最大异或和。一个数组的异或和为数组中所有的数异或起来的结果。
【思路】
在go的位移操作符中，要是int<<uint或者int>>uint形式，即右边必须是无符号整数;
在go中，int的最高位为符号位，正数的符号位为0，负数的符号位为1
1^1=0 1^0=1 因此 1^n 是对n取反
0^1=1 0^0=0 因此 0^n 还是n本身
1、可以采用前缀树来记录以i结尾的所有子数组的最大异或和，这样当下一个数a[i+1]到来时就知道了它跟哪个子数组异或得到的和最大，并把这个异或和也加入前缀树中。
2、求得最大异或和的方法为：对于int类型的数来说，都是32位，最高位为符号位，符号位的bestPath应保持与curPath保持相同；无论正数还是负数，其他位的bestPath应保持与curPath相反，才能得到最大异或和。
3、对a[i]的每一位，比如第k位，得到期望的bestPath后，还得看树中是否有这个bestPath，有则bestPath=bestPath否则bestPath=1^bestPath。然后把得到的第k位的最终异或值左移k位，并与当前已经得到的a[i]的前几位的异或值相或。
4、重复以上过程一直到a[i]的第32位，这样就可以得到以i位置结尾的所有子数组的最大异或和，然后更新全局变量maxEOR
*/
package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

type Node struct {
	next [2]*Node
}

func add(head *Node, num int) {
	var i uint
	cur := head
	for i = 31; i >= 0 && i < 32; i-- { /*如果写成 i=31;i>=0 ;i--  注意此处的i为无符号数，当i减到0后，会溢出变成最大值，因此循环不会停下来，所以限制条件要加上&& i<32*/
		path := (num >> i) & 1 /*&1的作用是只取最后一位*/
		if cur.next[path] == nil {
			cur.next[path] = &Node{}
		}
		cur = cur.next[path]
	}
}
func MaxEOR(head *Node, num int) int {
	res := 0
	var i uint
	cur := head
	for i = 31; i >= 0 && i < 32; i-- {
		path := (num >> i) & 1 /*path是当前数i位置上的值*/
		bestPath := 0          /*bestpath是如果想得到最大异或值，则当前i位置应该选择的异或对象*/
		if i < 31 {
			bestPath = path ^ 1 /*求理想的异或对象，应保证符号位为0，其他位为与自己相反的数，才能得到最大异或和*/
		}
		if cur.next[bestPath] == nil { /*实际上能存在的最理想的异或对象*/
			bestPath = bestPath ^ 1
		}
		res = res | ((path ^ bestPath) << i) /*得到i位置上的最大异或值，并把i位置上得到的最大异或值加到最终结果上*/
		cur = cur.next[bestPath]
	}
	return res
}
func getMaxEOR(arr []int) int {
	maxEOR := INT_MIN
	head := &Node{}
	eor := 0
	add(head, eor)
	for i := 0; i < len(arr); i++ {
		eor = eor ^ arr[i]
		res := MaxEOR(head, eor)
		maxEOR = max(maxEOR, res) /*把EOR(0~i)与树中记录的异或和依次异或，得到可能的最大异或和*/
		add(head, eor)            /*把arr[0]^arr[1]^...^arr[i-1]^arr[i]的值放入前缀树中，最终树中存储了从0到任意位置的异或值，实际上从任意位置p到i位置的异或值可由EOR(p~i)=EOR(0~i)^EOR(0~p)得到*/
	}
	return maxEOR
}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

/*对于数组中的每个元素，都计算出它到它后面的每个元素的异或和，然后选出最大的异或和。此种实现方式的时间复杂度为O（n^2）*/
func rightFunc1(arr []int) int {
	maxEOR := INT_MIN
	for i := 0; i < len(arr); i++ { /*外循环用来控制i从0到len*/
		eor := 0
		for j := i; j < len(arr); j++ { /*内循环用来计算i到len位置的异或和，可并得到最大值*/
			eor = eor ^ arr[j]
			maxEOR = max(maxEOR, eor)
		}
	}
	return maxEOR
}

/*使用dp数组记录0到i-1位置的异或和，就可知道任意位置p到i位置的异或和，从而选出最大的异或和。此种实现方式的时间复杂度为O（n^2）*/
func rightFunc2(arr []int) int {
	maxEOR := INT_MIN
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		dp[i] = dp[i-1] ^ arr[i]
		maxEOR = max(maxEOR, dp[i]) /*计算0～i位置的异或和的最大值*/
		for j := 0; j < i; j++ {
			eor := dp[j] ^ dp[i]
			maxEOR = max(maxEOR, eor) /*计算j到i位置的异或和的最大值*/
		}
	}
	return maxEOR
}
func main() {
	//arr:=[]int{1,0,2,5,0}
	arr := []int{1, 2, 4, 5}
	fmt.Println(rightFunc1(arr))
}
