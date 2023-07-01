package buffer

import "time"

type Node struct {
	data []uint8
	prev *Node
	next *Node
}

type DoublyLinkedd struct {
	inputChannel  chan []uint8
	outputChannel chan []uint8
	lengthRes     chan int
	head          *Node
	tail          *Node
	len           int
}

func NewBuffer() *DoublyLinkedd {
	cb := &DoublyLinkedd{
		inputChannel:  make(chan []uint8),
		outputChannel: make(chan []uint8),
		lengthRes:     make(chan int),
	}

	go cb.run()
	return cb
}

func (cb *DoublyLinkedd) run() {
	for {
		if cb.head == nil {
			val := <-cb.inputChannel
			node := &Node{data: val}
			cb.head, cb.tail = node, node
			cb.len++
		} else {
			select {
			case cb.lengthRes <- cb.len:
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

func (cb *DoublyLinkedd) Read() []uint8 {
	timer := time.NewTimer(time.Millisecond * 10)
	select {
	case item := <-cb.outputChannel:
		return item
	case <-timer.C:
		return nil
	}
}

func (cb *DoublyLinkedd) Len() int {
	timer := time.NewTimer(time.Millisecond * 10)
	select {
	case l := <-cb.lengthRes:
		return l
	case <-timer.C:
		return 0
	}
}
