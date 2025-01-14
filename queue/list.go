package queue

import (
	"time"

	"github.com/google/uuid"
)

type Element struct {
	// TODO: investigate stream
	Value       []uint8
	EnqueueTime time.Time
	ID          uuid.UUID
	Complete    chan struct{}
	next        *Element
	prev        *Element
	list        *List
}

type List struct {
	root Element
	len  int
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Remove(e *Element) []uint8 {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *List) PushFront(v *Element) *Element {
	l.lazyInit()
	return l.insert(v, &l.root)
}

func (l *List) PushBack(v *Element) *Element {
	l.lazyInit()
	return l.insert(v, l.root.prev)
}

func (l *List) PushBackElement(v *Element) *Element {
	l.lazyInit()
	return l.insert(v, l.root.prev)
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

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// func (l *List) insertValue(v []uint8, at *Element) *Element {
// 	return l.insert(&Element{ID: uuid.New(), Value: v, EnqueueTime: time.Now()}, at)
// }

func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}
