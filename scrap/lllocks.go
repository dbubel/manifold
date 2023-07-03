package scrap

import (
	"context"
	"sync"
	"time"
)

type LinkedListLocks struct {
	head *Node
	tail *Node
	lock sync.RWMutex
}

func NewLinkedListLocks() *LinkedListLocks {
	return &LinkedListLocks{}
}

func (ll *LinkedListLocks) PushBack(data uint8) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	node := &Node{data: data}

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		node.prev = ll.tail
		ll.tail = node
	}
}

func (ll *LinkedListLocks) Front(ctx context.Context) *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	if ll.head == nil {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(100 * time.Millisecond):
			return ll.head
		}
	}

	return ll.head
}

func (ll *LinkedListLocks) Remove(data uint8) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	for node := ll.head; node != nil; node = node.next {
		if node.data == data {
			if node == ll.head {
				ll.head = node.next
				if ll.head != nil {
					ll.head.prev = nil
				}
			} else if node == ll.tail {
				ll.tail = node.prev
				if ll.tail != nil {
					ll.tail.next = nil
				}
			} else {
				node.prev.next = node.next
				node.next.prev = node.prev
			}
			return
		}
	}
}
