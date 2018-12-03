/*
三种内存替换算法：LIFO（先来先走，用队列实现），LRU（least recently used，清除最近未使用的，用双向链表实现），LFU（least frequently used，清除使用频率最小的，用二维链表实现）

设计可以变更的缓存结构（ LRU）
【 题目】
设计一种缓存结构， 该结构在构造时确定大小， 假设大小为K， 并有两个功能：
set(key,value)： 将记录(key,value)插入该结构。
get(key)： 返回key对应的value值。
【 要求】
1． set和get方法的时间复杂度为O(1)。
2． 某个key的set或get操作一旦发生， 认为这个key的记录成了最经常使用的。
3． 当缓存的大小超过K时， 移除最不经常使用的记录， 即set或get最久远的。
【 举例】
假设缓存结构的实例是cache， 大小为3， 并依次发生如下行为：
1． cache.set("A",1)。 最经常使用的记录为("A",1)。
2． cache.set("B",2)。 最经常使用的记录为("B",2)， ("A",1)变为最不经常的。
3． cache.set("C",3)。 最经常使用的记录为("C",2)， ("A",1)还是最不经常的。
4． cache.get("A")。 最经常使用的记录为("A",1)， ("B",2)变为最不经常的。
5． cache.set("D",4)。 大小超过了3， 所以移除此时最不经常使用的记录("B",2)，
加入记录 ("D",4)， 并且为最经常使用的记录， 然后("C",2)变为最不经常使用的记录

【思路】
1、使用一个双向链表来实现cache，set时把节点加到head，如果cache的容量已满，则先从tail删除一个元素，再从head添加，get时也把节点换到队头。
2、把一个Node set进cache时，同时把它的key和对应的Node记录在一个map中，这样每次set或get一个Node之前就可以查看它是否已经在cache中。
map中的value存的是Node的内存地址，因此直接修改从map中得到的Node就相当于在链表中对Node修改。
*/
package main

import "fmt"

type Node struct {
	key   string
	value int
	pre   *Node
	next  *Node
}
type dbList struct {
	cap  int
	size int
	head *Node
	tail *Node
}

var nodes = make(map[string]*Node)

/*
删除步骤：
1、保证head不为n，即不会删除第一个元素，只会更改第一个元素
1、判断tail是否为n：若tail为n，则更改tail：1、tail=n.pre;2、tail.next=nil;3、n.pre=nil
2、若tail不为n，则更改n前后元素的指针：1、n.pre.next=n.next 2、n.next.pre=n.pre 3、n.pre=nil 4、n.next=nil
3、删除元素后记得size--，并从nodeMap里面删除n
*/
func del(cache *dbList, n *Node) {
	if cache.tail == n {
		cache.tail = n.pre
		n.pre.next = nil
		n.pre = nil
	} else {
		n.next.pre = n.pre
		n.pre.next = n.next
		n.pre = nil
		n.next = nil
	}
	cache.size--
	delete(nodes, n.key)
}

/*
插入步骤：
1、若head,tail都为空，则直接插入，并让head=tail=n
2、若head,tail不为空，则看size是否等于cap，如果等于，则删除tail
3、然后把元素插入到头部：1、head.pre=n 2、n.next=head 3、head=n
4、最后让size++，并在nodeMap里面加入n
*/
func insert(cache *dbList, n *Node) {
	if cache.head == nil && cache.tail == nil {
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
	nodes[n.key] = n
}

/*
set步骤：
0、如果cap为0则直接退出，如果cap为1则直接让head,tail都为n，size为1。
1、判断newN的key是否已在nodeMap中，若在，判断oldN是否为head，为head则直接更改value。不为head则先删除oldN再插入newN。
2、如果newN的key不在node中，则直接插入newN
*/
func (cache *dbList) set(k string, v int) {
	if cache.cap == 0 {
		fmt.Println("cap should more than 0!")
		return
	}
	newN := &Node{k, v, nil, nil}
	if cache.cap == 1 {
		cache.head = newN
		cache.tail = newN
		cache.size = 1
		return
	}
	if oldN, ok := nodes[k]; ok {
		if cache.head != oldN {
			del(cache, oldN)
			insert(cache, newN)
		} else {
			oldN.value = v
		}
	} else {
		insert(cache, newN)
	}
}

/*
get步骤：
0、如果cache中没有元素，则直接返回-1，如果查询的k为head.key则直接返回head.value
1、判断n的key是否在nodeMap中，若不在直接返回-1
2、若在则判断n是否为head，若为head则直接返回n.value
3、若不为head，则先删除n，再插入n，并返回n.value
*/
func (cache *dbList) get(k string) int {
	if cache.size == 0 {
		fmt.Printf("There is no item in cache! ")
		return -1
	}
	if cache.head.key == k {
		return cache.head.value
	}
	if n, ok := nodes[k]; ok {
		v := n.value
		if n == cache.head {
			return v
		}
		del(cache, n)
		insert(cache, n)
		return v
	}
	fmt.Printf("Not existed! ")
	return -1
}

func main() {
	cache := &dbList{1, 0, nil, nil}
	cache.set("a", 1)
	cache.set("b", 2)
	cache.set("c", 3)
	cache.set("d", 4)
	cache.set("a", 5)
	cache.set("a", 6)
	fmt.Println(cache.get("a"))
	fmt.Println(cache.get("b"))
}
