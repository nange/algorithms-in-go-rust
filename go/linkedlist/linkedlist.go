package linkedlist

type Lister interface {
	Clear()
	Len() int
	Front() *Element
	Back() *Element
	Remove(e *Element) interface{}
	PushFront(v interface{}) *Element
	PushBack(v interface{}) *Element
	InsertBefore(v interface{}, mark *Element) *Element
	InsertAfter(v interface{}, mark *Element) *Element
	MoveBefore(e, mark *Element)
	MoveAfter(e, mark *Element)
}

type Element struct {
	prev, next *Element
	list       *List
	Value      interface{}
}

func (e *Element) Next() *Element {
	if n := e.next; n != nil && n != &e.list.root {
		return n
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; p != nil && p != &e.list.root {
		return p
	}
	return nil
}

type List struct {
	root Element
	len  int
}

func New() *List {
	l := new(List)
	l.Clear()
	return l
}

func (l *List) Clear() {
	if firstElem := l.root.next; firstElem != nil && firstElem.list == l {
		firstElem.prev = nil
	}
	if lastElem := l.root.prev; lastElem != nil && lastElem.list == l {
		lastElem.next = nil
	}
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	e.list = nil
	l.len--
	return e
}

func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *List) PushFront(v interface{}) *Element {
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Element {
	return l.insertValue(v, l.root.prev)
}

func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
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
