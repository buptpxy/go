package main

import "fmt"

type Node struct {
	key   string
	value int
	pre   *Node
	next  *Node
}

type dbList struct {
	cap   int
	size  int
	head  *Node
	tail  *Node
	nodes map[string]*Node
}

func del(cache *dbList, n *Node) {
	if n == cache.tail {
		cache.tail = n.pre
		n.pre.next = nil
		n.pre = nil
	} else {
		n.pre.next = n.next
		n.next.pre = n.pre
		n.pre = nil
		n.next = nil
	}
	cache.size--
	delete(cache.nodes, n.key)
}

func insert(cache *dbList, n *Node) {
	if cache.head == nil {
		cache.tail = n
	} else {
		if cache.size == cache.cap {
			del(cache, cache.tail)
		}
		cache.head.pre = n
		n.next = cache.head
	}
	cache.head = n
	cache.size++
	cache.nodes[n.key] = n
}

func (cache *dbList) set(k string, v int) {
	if cache.cap == 0 {
		return
	}
	newN := &Node{k, v, nil, nil}
	if cache.cap == 1 {
		cache.head = newN
		cache.tail = newN
		cache.size = 1
		return
	}
	if oldN, ok := cache.nodes[k]; ok {
		if oldN == cache.head {
			oldN.value = v
		} else {
			del(cache, oldN)
			insert(cache, newN)
		}
	} else {
		insert(cache, newN)
	}
}

func (cache *dbList) get(k string) int {
	if cache.size == 0 {
		return -1
	}
	if cache.head.key == k {
		return cache.head.value
	}
	if n, ok := cache.nodes[k]; ok {
		v := n.value
		del(cache, n)
		insert(cache, n)
		return v
	}
	return -1
}

func main() {
	cache := &dbList{3, 0, nil, nil, make(map[string]*Node)}
	cache.set("a", 1)
	cache.set("b", 2)
	cache.set("c", 3)
	cache.set("d", 4)
	cache.set("a", 5)
	cache.set("a", 6)
	fmt.Println(cache.get("a"))
	fmt.Println(cache.get("b"))
}
