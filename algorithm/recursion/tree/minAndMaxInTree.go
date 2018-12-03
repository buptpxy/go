/*
用递归的方法求一棵树的最大和最小值。
【思路】
1、类似的题可以转换成求出以每个节点的为跟的子树的最大最小值
2、划分出可能情况，有三中可能情况：1）最大最小值来自于左子树 2）最大最小值来自于右子树 3）最大最小值来自于自己
3、每个节点需要收集的信息：1）子树最大值 2）子树最小值
4、baseCase：head==nil时，向上传的返回值应不影响上层的判断，故让min=系统最大值，max=系统最小值
*/
/*
golang的标准库里没有定义系统最小值和系统最大值。不过可以用位操作运算，轻松定义这些常量。
无符号整型uint
其最小值是0，其二进制表示的所有位都为0，
const UINT_MIN uint = 0
其最大值的二进制表示的所有位都为1，那么，
const UINT_MAX = ^uint(0)

有符号整型int
根据补码，其最大值二进制表示，首位0（0表示正数），其余1，那么，
const INT_MAX = int(^uint(0) >> 1)
根据补码，其最小值二进制表示，首位1（1表示负数），其余0，那么，
const INT_MIN = ^INT_MAX
*/
package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

type Node struct {
	value int
	left  *Node
	right *Node
}

type res struct {
	min int
	max int
}

//O(2n-1)
func minAndMax(head *Node) res {
	fmt.Println("rec")
	if head == nil {
		return res{INT_MAX, INT_MIN}
	}
	lRes := minAndMax(head.left)
	rRes := minAndMax(head.right)
	minN := min(lRes.min, rRes.min)
	maxN := max(lRes.max, rRes.max)
	return res{min(head.value, minN), max(head.value, maxN)}
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func main() {
	head := &Node{10, nil, nil}
	head.left = &Node{11, nil, nil}
	head.right = &Node{9, nil, nil}
	head.left.left = &Node{13, nil, nil}
	head.left.right = &Node{17, nil, nil}
	head.right.left = &Node{15, nil, nil}
	head.right.right = &Node{16, nil, nil}
	r := minAndMax(head)
	fmt.Println(r)
}
