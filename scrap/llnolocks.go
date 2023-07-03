package scrap

import (
	"context"
)

type Node struct {
	data uint8
	prev *Node
	next *Node
}

type LinkedListChan struct {
	head      *Node
	tail      *Node
	pushBack  chan *Node
	getFront  chan chan *Node
	removeVal chan uint8
	quit      chan struct{}
}

func NewLinkedList() *LinkedListChan {
	ll := &LinkedListChan{
		pushBack:  make(chan *Node),
		getFront:  make(chan chan *Node),
		removeVal: make(chan uint8),
		quit:      make(chan struct{}),
	}

	go ll.run()

	return ll
}

func (ll *LinkedListChan) PushBack(data uint8) {
	ll.pushBack <- &Node{data: data}
}

func (ll *LinkedListChan) Front(ctx context.Context) *Node {
	frontChan := make(chan *Node)
	defer close(frontChan)
	ll.getFront <- frontChan
	select {
	case node := <-frontChan:
		return node
	case <-ctx.Done():
		return nil
	}
}

func (ll *LinkedListChan) Remove(data uint8) {
	ll.removeVal <- data
}

func (ll *LinkedListChan) run() {
	for {
		select {
		case node := <-ll.pushBack:
			if ll.head == nil {
				ll.head = node
				ll.tail = node
			} else {
				ll.tail.next = node
				node.prev = ll.tail
				ll.tail = node
			}
		case frontChan := <-ll.getFront:
			frontChan <- ll.head
		case data := <-ll.removeVal:
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
					break
				}
			}
		case <-ll.quit:
			return
		}
	}
}
