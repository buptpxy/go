/*
图的存储方式
1） 邻接表
2） 邻接矩阵
	如何表达图？ 生成图？
	用来存储图的矩阵：第一列是边的权重，第二列是from节点，第三列是to节点
*/
package main

import (
	"fmt"
	// mapset "github.com/deckarep/golang-set"
	"container/list"
)

type Edge struct {
	weight int
	from   Node
	to     Node
}
type Node struct {
	value int
	in    int
	out   int
	nexts []Node
	edges []Edge
}
type Graph struct {
	nodes map[int]Node
	edges &list.List
}

func newGraph(m [][]int) Graph {
	g := Graph{make(map[int]Node), list.New()}
	for i := 0; i < len(m); i++ {
		weight := m[i][0]
		from := m[i][1]
		to := m[i][2]
		if _, ok := g.nodes[from]; !ok {
			g.nodes[from] = Node{0, 0, 0, nil, nil}
		}
		if _, ok := g.nodes[to]; !ok {
			g.nodes[to] = Node{0, 0, 0, nil, nil}
		}
		fromNode := g.nodes[from]
		toNode := g.nodes[to]
		fromNode.out++
		fromNode.nexts = append(g.nodes[from].nexts, g.nodes[to])
		toNode.in++
		newEdge := Edge{weight, fromNode, toNode}

		if !g.edges.Contains(newEdge) {
			g.edges.Add(newEdge)
			fromNode.edges = append(g.nodes[from].edges, newEdge)
		}
		g.nodes[from] = fromNode
		g.nodes[to] = toNode
	}
	return g
}

type test struct {
	s []interface{}
}

func main() {
	// m := make(map[int]int)
	// m[1] = 1
	// v, ok := m[1] //直接v:= m[1]也可以
	// fmt.Println(v)
	// fmt.Println(ok)

	// m := [][]int{
	// 	{5, 1, 2},
	// 	{8, 1, 3},
	// 	{7, 2, 4},
	// 	{6, 4, 3},
	// }
	// g := newGraph(m)
	// fmt.Println(g)

	s := []interface{}{1, 2, 3}
	// t := test{s}
	// a := mapset.NewSetFromSlice(t) //cannot use t (type test) as type []interface {} in argument to mapset.NewSetFromSlice
	// b := mapset.NewSet([]interface{}{1, 2, 3}) // hash of unhashable type []interface {}
	a := mapset.NewSetFromSlice(s) //Set{1, 2, 3}
	fmt.Println(a)
	o := []rune{1, 2, 3}
	str := string(o)
	b := []byte(str)
	h := sha256.New()
	h.Write(b)
	hashed := h.Sum(nil)
	c := hex.EncodeToString(hashed)
	fmt.Println(c)
}
