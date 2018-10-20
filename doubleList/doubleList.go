package doubleList

/*
type Element
    func (e *Element) Next() *Element
    func (e *Element) Prev() *Element
type List
    func New() *List
    func (l *List) Back() *Element
    func (l *List) Front() *Element
    func (l *List) Init() *List
    func (l *List) InsertAfter(v interface{}, mark *Element) *Element
    func (l *List) InsertBefore(v interface{}, mark *Element) *Element
    func (l *List) Len() int
    func (l *List) MoveAfter(e, mark *Element)
    func (l *List) MoveBefore(e, mark *Element)
    func (l *List) MoveToBack(e *Element)
    func (l *List) MoveToFront(e *Element)
    func (l *List) PushBack(v interface{}) *Element
    func (l *List) PushBackList(other *List)
    func (l *List) PushFront(v interface{}) *Element
    func (l *List) PushFrontList(other *List)
    func (l *List) Remove(e *Element) interface{}
*/
type Element struct {
	Value      interface{}
	next, prev *Element //next,prev,list都是指针
	list       *List    //list是指向root的指针
}
type List struct {
	root Element //root不是指针，是链表的头结点
	Len  int
}

//element的方法
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

//List的方法
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.Len = 0
	return l
}
func New() *List {
	return new(List).Init()
}
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}
func (l *List) Front() *Element {
	if l.Len == 0 {
		return nil
	}
	return l.root.next //l.root不是*Element类型，故此处不能是l.root.Next()
}
func (l *List) Back() *Element {
	if l.Len == 0 {
		return nil
	}
	return l.root.prev
}
func (l *List) insert(e, after *Element) *Element {
	n := after.next
	after.next = e
	e.next = n
	e.prev = after
	n.prev = e
	e.list = l //e属于l
	l.Len++
	return e
}
func (l *List) insertValue(v interface{}, after *Element) *Element {
	return l.insert(&Element{Value: v}, after)
}
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}
func (l *List) remove(e *Element) *Element {
	e.next.prev = e.prev
	e.prev.next = e.next
	e.next = nil
	e.prev = nil
	e.list = nil
	l.Len--
	return e
}
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || mark.list != l || e == mark {
		return
	}
	l.insert(l.remove(e), mark.prev)
}
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || mark.list != l || e == mark {
		return
	}
	l.insert(l.remove(e), mark)
}
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.insert(l.remove(e), &l.root) //因为root不是指针，故要加&
}
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.insert(l.remove(e), l.root.prev)
}
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}
func (l *List) PushFrontList(other *List) {
	// l.lazyInit()
	for i, e := other.Len, other.Back(); i > 0; i, e = i-1, e.prev {
		l.PushFront(e.Value)
	}
}
func (l *List) PushBackList(other *List) {
	// l.lazyInit()
	for i, e := other.Len, other.Front(); i > 0; i, e = i-1, e.next {
		l.PushBack(e.Value)
	}
}
