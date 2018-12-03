/*
并查集结构。
一个集合中，指针指向自己的元素为集合的代表元素，其他的元素最终指针都会指向代表元素。
在查询两个元素是否属于同一个集合时，只需查他们的代表元素是否相同。
在合并两个集合时，只需让集合元素少的那个集合的代表元素的指针指向集合元素多的那个集合的代表元素。
把这些集合都存入一个map中，map中的key就是元素的值，value就是这个元素指向的元素的值。若一个key的value为自身，则这个key为一个集合的代表元素
*/
package main

import (
	"fmt"
)

type Node interface{}

var fatherMap map[Node]Node
var sizeMap map[Node]int

func makeSets(head Node, nodes []Node) {
	fatherMap = make(map[Node]Node)
	sizeMap = make(map[Node]int)
	fatherMap[head] = head
	sizeMap[head] = 1
	for _, v := range nodes {
		fatherMap[v] = head
		sizeMap[head]++
	}
}
func updateSizeMap() {
	sizeMap = make(map[Node]int)
	for _, v := range fatherMap {
		sizeMap[v]++
	}
}
func findHead(node Node) Node {
	if fatherMap == nil {
		return nil
	}
	_, ok := fatherMap[node]
	if !ok {
		return nil
	}
	n := node
	for fatherMap[n] != n {
		n = fatherMap[n]
	}
	return n
}
func isSameSets(node1, node2 Node) bool {
	return findHead(node1) == findHead(node2)
}
func unionFind(node1, node2 Node) {
	if isSameSets(node1, node2) {
		fmt.Println("they are in same sets!")
		return
	}
	head1 := findHead(node1)
	head2 := findHead(node2)
	if sizeMap[head1] >= sizeMap[head2] {
		fatherMap[head2] = head1
		sizeMap[head1] = sizeMap[head1] + sizeMap[head2]
	} else {
		fatherMap[head1] = head2
		sizeMap[head2] = sizeMap[head2] + sizeMap[head1]
	}
}
func main() {
	// makeSets(2, []Node{2, 3, 5, 6})
	// fmt.Println(fatherMap)
	fatherMap = map[Node]Node{1: 1, 2: 1, 3: 2, 4: 3, 5: 1, 6: 6, 7: 6, 8: 6}
	updateSizeMap()
	// fmt.Println(findHead(3))
	// fmt.Println(findHead(8))
	unionFind(2, 7)
	fmt.Println(fatherMap)
}
