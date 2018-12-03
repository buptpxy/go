/*
实现LFU中的set 和 get，时间复杂度都为O（1）
【思路】
1、用二维双链表实现实现，这个mdbList由多条dbList组成，每条dbList上面放一种times的node。
2、mdbList有它的cap,size,headList,还有一个nodesMap记录元素是否被加入过，还有一个timesMap记录每个times对应的dbList。
3、set规则为：如果key未在nodesMap中出现,则直接把node insert 进times为1的dbList;如果key在nodesMap中出现，则先取得这个节点的times，再del这个node，再把这个node insert进为times+1的dbList
4、get规则为：如果key未在nodesMap中出现，则返回-1；如果key在nodesMap中出现，则先取得这个节点的times，再del这个node，再把这个node insert进为times+1的dbList
5、del规则为：如果要删除的node=tail，则判断node.pre是否等于head，等于的话则直接删除这个cList： 若删掉的cList==headList，则headList=cList.head.right.list;headList.head.left=nil;
																						若删掉的cList!=headList，则head.left.right=head.right;head.right.left=head.left;head.left=nil;head.right=nil，
																					然后delete(timesMap,node.time)
															不等于的话则删除node并更改tail: tail=node.pre;tail.next=nil;node.pre=nil
			如果要删除的node!=tail，则直接删除node： node.pre.next=node.next;node.next.pre=node.pre;node.pre=nil;node.next=nil
			最后把mdbList.size--并delete(nodesMap,node.key)
6、insert规则为：如果size==cap，则先删掉headList的tail。
如果要插入的node.times不存在timesMap中，则新建一个dbList：若headList为空，则newList:=&dbList{&Node{0,0,node.times,nil,node,nil,nil},node};headList=newList
													若headList不为空，则从headList一直往右找，直到找到dbList.head.times>node.times为止，然后在找到的aimList之前插入一个新的dbList:
															若aimList==headList，则newList:=&dbList{&Node{0,0,node.times,nil,node,nil,aimHead,newList},node};headList=newList;aimList.head.left=newList.head;node.pre=head
															若aimList!=headList，则newList:=&dbList{&Node{0,0,node.times,nil,node,aimHead.left,aimHead,newList},node};aimList.head.left=newList;node.pre=head
									最后让timesMap[node.times]=node.list
如果要插入的node.times存在timesMap中，则把nodes添加到得到的cList的head后面：nodes.pre=cList.head;node.next=cList.head.next;cList.head.next.pre=node;cList.head.next=node
最后把mdbList.size++并nodesMap[key]=node
*/
package main

import "fmt"

type Node struct {
	key   string
	value int
	times int
	pre   *Node
	next  *Node
	left  *Node
	right *Node
	list  *dbList
}

type dbList struct {
	head *Node
	tail *Node
}

type mdbList struct {
	cap      int
	size     int
	headList *dbList
	nodesMap map[string]*Node
	timesMap map[int]*dbList
}

func del(cache *mdbList, n *Node) {
	cList := cache.timesMap[n.times]
	head := cList.head
	tail := cList.tail
	if n == tail { /*n为尾节点，则需注意更改尾节点*/
		if n.pre == head { /*n的上一个为头节点，即链表里面只有n一个元素，则删除这个元素的同时还需删除这个链表*/
			if head.right != nil {
				head.left.right = head.right
				head.right.left = head.left
				head.right = nil
			} else { /*如果当前链表是最右边的一个链表，则直接让左边链表头的right=nil*/
				head.left.right = nil
			}
			head.left = nil
			delete(cache.timesMap, n.times) /*既然链表已经删除，则在timesMap中删除对应times的记录*/
		} else {
			tail = n.pre
			tail.next = nil
			n.pre = nil
		}
	} else {
		n.pre.next = n.next
		n.next.pre = n.pre
		n.pre = nil
		n.next = nil
	}
	/*操作完后记得对cache中的有更改的属性赋值*/
	cache.size--
	delete(cache.nodesMap, n.key)
	if _, ok := cache.timesMap[n.times]; ok {
		cache.timesMap[n.times].head = head
		cache.timesMap[n.times].tail = tail
	}
}

func insert(cache *mdbList, n *Node) {
	if cache.size == cache.cap { /*当cache容量满了，需先删除一个元素，删除的是头链表后面第一条链表的尾节点*/
		del(cache, cache.headList.head.right.list.tail)
	}
	timesMap := cache.timesMap
	newList := &dbList{}
	if cList, ok := timesMap[n.times]; !ok {
		aimHead := cache.headList.head
		aimList := cache.headList
		for aimHead.times < n.times {
			if aimHead.right == nil { /*如果aimList已经是最右边的list，则让newList直接连到aimList的右边*/
				aimList = timesMap[aimHead.times]
				newList = &dbList{&Node{"", 0, n.times, nil, n, aimHead, nil, nil}, n}
				newList.head.list = newList
				aimList.head.right = newList.head
				newList.head.left = aimList.head
				break
			}
			aimHead = aimHead.right
		}
		if aimHead.right != nil && aimHead.times > n.times { /*如果aimList不是最右边的List，且是times大于n.times的List，所以让newList连到aimList的左边*/
			aimList = timesMap[aimHead.times]
			newList = &dbList{&Node{"", 0, n.times, nil, n, aimHead.left, aimHead, nil}, n}
			newList.head.list = newList
			aimList.head.left = newList.head
		}
		n.pre = newList.head
		n.list = newList
		/*更改完list后timesMap需要更新一下*/
		timesMap[aimHead.times] = aimList
		timesMap[n.times] = n.list
	} else {
		n.pre = cList.head
		n.next = cList.head.next
		cList.head.next.pre = n
		cList.head.next = n
	}
	/*更新cache的变化的属性*/
	cache.size++
	cache.nodesMap[n.key] = n
	cache.timesMap = timesMap
}

func (cache *mdbList) set(k string, v int) {
	if cache.cap == 0 { /*cap为0直接返回*/
		return
	}
	n := &Node{k, v, 1, nil, nil, nil, nil, nil}
	if cache.cap == 1 { /*cap为1则将元素直接存在headList的头结点中*/
		cache.size = 1
		cache.headList = &dbList{&Node{k, v, 1, nil, n, nil, nil, nil}, n}
		cache.nodesMap = make(map[string]*Node) /*每次set先清空nodesMap*/
		cache.nodesMap[k] = cache.headList.head
		cache.timesMap[1] = cache.headList
		return
	}
	if oldN, ok := cache.nodesMap[k]; !ok {
		insert(cache, n)
	} else {
		del(cache, oldN)
		newN := oldN
		newN.value = v
		newN.times = oldN.times + 1
		insert(cache, newN)
	}
}

func (cache *mdbList) get(k string) int {
	if cache.cap == 0 { /*cap为0直接返回*/
		return -1
	}
	if oldN, ok := cache.nodesMap[k]; ok {
		if cache.cap == 1 { /*cap为1直接返回value*/
			return oldN.value
		}
		del(cache, oldN)
		newN := oldN
		newN.times = oldN.times + 1
		insert(cache, newN)
		return newN.value
	}
	return -1
}

func main() {
	/*不在头链表或头结点存储数据，减少特殊边界情况*/
	cache := &mdbList{3, 0, &dbList{&Node{"", 0, 0, nil, nil, nil, nil, nil}, nil}, make(map[string]*Node), make(map[int]*dbList)}
	cache.headList.head.list = cache.headList /*初始化时要注意头链表和头结点*/
	cache.timesMap[0] = cache.headList
	cache.set("a", 1)
	cache.set("b", 2)
	cache.set("c", 3)
	cache.set("d", 4)
	cache.set("a", 5)
	cache.set("a", 6)
	fmt.Println(cache.get("a"))
	fmt.Println(cache.get("b"))
}
