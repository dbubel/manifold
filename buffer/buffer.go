package buffer

import (
	"context"
	"fmt"
	"sync"
)

type Node struct {
	data []uint8
	prev *Node
	next *Node
}

type DoublyLinkedd struct {
	inputChannel  chan []uint8
	outputChannel chan []uint8
	//lengthRes     chan int
	m    sync.Mutex
	head *Node
	tail *Node
	len  int
}

func NewBuffer() *DoublyLinkedd {
	cb := &DoublyLinkedd{
		inputChannel:  make(chan []uint8),
		outputChannel: make(chan []uint8),
		//lengthRes:     make(chan int),
		len: 0,
	}

	go cb.run()
	return cb
}

func (cb *DoublyLinkedd) run() {
	for {
		// this prevents the len function from not needing the sync lock.
		if cb.head == nil {
			val := <-cb.inputChannel
			node := &Node{data: val}
			cb.head, cb.tail = node, node
			cb.len++
		} else {
			select {
			//case cb.lengthRes <- cb.len:
			case val := <-cb.inputChannel:
				node := &Node{data: val, prev: cb.tail}
				cb.tail.next = node
				cb.tail = node
				cb.len++
			case cb.outputChannel <- cb.head.data:
				if cb.head == cb.tail {
					cb.head, cb.tail = nil, nil
				} else {
					cb.head = cb.head.next
					cb.head.prev = nil
				}
				cb.len--
			}
		}
	}
}

func (cb *DoublyLinkedd) Write(val []uint8) {
	cb.inputChannel <- val
}

func (cb *DoublyLinkedd) Read(ctx context.Context) []uint8 {
	select {
	case item := <-cb.outputChannel:
		//fmt.Println("item", item)
		//ctx.Done()
		return item
	case <-ctx.Done():
		fmt.Println("ctx done")
		return nil
	}
}

func (cb *DoublyLinkedd) Len(ctx context.Context) int {
	//timer := time.NewTimer(1 * time.Second)
	//return cb.len
	// TODO:(dean) this bugs me a lot. This is the only lock used.
	cb.m.Lock()
	defer cb.m.Unlock()
	return cb.len
	//select {
	//case i := <-cb.lengthRes:
	//	return i
	//case <-ctx.Done():
	//	return 7
	//}
}
