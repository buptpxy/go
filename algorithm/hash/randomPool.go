/*
设计RandomPool结构
【 题目】 设计一种结构， 在该结构中有如下三个功能：
insert(key)： 将某个key加入到该结构， 做到不重复加入。
delete(key)： 将原本在结构中的某个key移除。
getRandom()：等概率随机返回结构中的任何一个key。
【 要求】 Insert、 delete和getRandom方法的时间复杂度都是O(1)
【思路】由于hash表的key是不重复的，故可以用两个hash表来实现，
	第一个hash表的key对应的value为key加入的顺序，第二个hash表的key为加入的顺序，value为表一中的key
	要返回等概率的值可以利用一个0~size-1范围内的随机数，来查询对应的key
	为了使删除key之后仍能等概率查询，可以把最后加入的元素与要删除的元素互换，再删除要删除的元素
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateMap(m1 map[int]int) map[int]int {
	m2 := make(map[int]int)
	for k, v := range m1 {
		m2[v] = k
	}
	return m2
}
func insert(m1 map[int]int, m2 map[int]int, key int) {
	size := len(m1)
	if _, ok := m1[key]; ok {
		fmt.Printf("%d is already in map!\n", key)
		return
	}
	m1[key] = size
	m2[size] = key
}
func getRandom(m2 map[int]int) int {
	size := len(m2)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(size) //产生0~size-1范围内的随机数
	return m2[num]
}
func deleteKey(m1 map[int]int, m2 map[int]int, key int) {
	if _, ok := m1[key]; !ok {
		fmt.Printf("%d is not in map!\n", key)
		return
	}
	size := len(m1)
	valueK := m1[key]
	keyS := m2[size-1]
	m1[key], m1[keyS] = m1[keyS], m1[key]
	m2[valueK], m2[size-1] = m2[size-1], m2[valueK]
	delete(m1, key)
	delete(m2, size-1)
}
func main() {
	m1 := map[int]int{1: 0, 32: 1, 54: 2, 75: 3, 45: 4, 55: 5, 76: 6, 72: 7}
	m2 := generateMap(m1)
	insert(m1, m2, 9)
	deleteKey(m1, m2, 75)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(getRandom(m2))
}
