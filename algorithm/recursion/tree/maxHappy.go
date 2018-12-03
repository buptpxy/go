/*
一个公司的上下节关系是一棵多叉树， 这个公司要举办晚会， 你作为组织者已经摸清了大家的心理： 一个员工的直
接上级如果到场， 这个员工肯定不会来。 每个员工都有一个活跃度的值， 决定谁来你会给这个员工发邀请函， 怎么
让舞会的气氛最活跃？ 返回最大的活跃值。
举例：给定一个矩阵来表述这种关系
	matrix ={{1 , 6}， {1 , 5}，{1 , 4}}
这个矩阵的含义是：
matrix[0] = {1 , 6}， 表示0这个员工的直接上级为1,0这个员工自己的活跃度为6
matrix[1] = {1 , 5}， 表示1这个员工的直接上级为1（ 他自己是这个公司的最大boss） ,1这个员工自己的活跃度为5
matrix[2] = {1 , 4}， 表示2这个员工的直接上级为1,2这个员工自己的活跃度为4
为了让晚会活跃度最大， 应该让1不来， 0和2来。 最后返回活跃度为10
【思路】先把矩阵变成树，基于树再去计算以每个节点为根的树的最大活跃度
1、两种情况：1）自己来，则子节点都不来，则最大活跃度为每个子节点都不来的最大活跃度相加 2）自己不来，则子节点可以来也可以不来，则最大活跃度为每个子节点max(go,ungo)的最大活跃度之和
2、每个节点要收集的信息：自己来的最大活跃度goHappy，自己不来的最大活跃度ungoHappy
3、basecase:head==nil，则goHappy=0,ungoHappy=0
此递归的时间复杂度较低，不用转dp数组
*/
package main

import (
	"fmt"
)

type Node struct {
	value int
	child []*Node
}

type res struct {
	goHappy   int
	ungoHappy int
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//O(n)
func maxHappy(head *Node) res {
	//fmt.Println("rec")
	if head == nil {
		return res{0, 0}
	}
	goHappy := head.value
	ungoHappy := 0
	for i := 0; i < len(head.child); i++ {
		cRes := maxHappy(head.child[i])
		goHappy = goHappy + cRes.ungoHappy
		ungoHappy = ungoHappy + max(cRes.goHappy, cRes.ungoHappy)
	}
	return res{goHappy, ungoHappy}
}

func generateTree(m [][]int) *Node {
	nodeMap := make(map[int]*Node)
	parentMap := make(map[int]int)
	head := 0
	for i := 0; i < len(m); i++ {
		n := &Node{m[i][1], nil}
		nodeMap[i] = n
		parentMap[i] = m[i][0]
	}
	for n1, node := range nodeMap {
		for n2, p := range parentMap {
			if n2 != p {
				if n1 == p {
					node.child = append(node.child, nodeMap[n2])
					delete(parentMap, n2)
				}
			} else {
				head = n2
				delete(parentMap, n2)
			}
		}
	}
	return nodeMap[head]
}

func main() {
	//matrix:=[][]int{{1,6},{1,5},{1,4}}
	matrix := [][]int{{1, 4}, {3, 5}, {1, 6}, {3, 10}, {3, 112}, {4, 11}, {4, 13}}
	head := generateTree(matrix)
	fmt.Println(head)
	happy := maxHappy(head)
	fmt.Println(max(happy.ungoHappy, happy.goHappy))
}
