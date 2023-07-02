package buffer

import (
	"sync"
)

type Node struct {
	data []uint8
	prev *Node
	next *Node
}

type DoublyLinkedd struct {
	InputChannel  chan []uint8
	OutputChannel chan []uint8
	head          *Node
	tail          *Node
	len           int
	m             sync.RWMutex
}

func NewBuffer(inputChannel, outputChannel chan []byte) *DoublyLinkedd {
	cb := &DoublyLinkedd{
		InputChannel:  inputChannel,
		OutputChannel: outputChannel,
		len:           0,
	}

	go cb.run()
	return cb
}

func (cb *DoublyLinkedd) run() {
	for {
		// this prevents the len function from not needing the sync lock.
		if cb.head == nil {
			//cb.LenChannel <- 0
			val := <-cb.InputChannel
			node := &Node{data: val}
			cb.head, cb.tail = node, node
			cb.len++
		} else {
			select {
			//case cb.LenChannel <- cb.len:
			case val := <-cb.InputChannel:
				node := &Node{data: val, prev: cb.tail}
				cb.tail.next = node
				cb.tail = node
				cb.len++
			case cb.OutputChannel <- cb.head.data:
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

func (cb *DoublyLinkedd) Len() int {
	// TODO:(dbubel) figure out a way to do this with a channel
	cb.m.RLock()
	defer cb.m.RUnlock()
	return cb.len
}
