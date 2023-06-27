package buffer

type Node struct {
	data []uint8
	prev *Node
	next *Node
}

type DoublyLinkedd struct {
	inputChannel  chan []uint8
	outputChannel chan []uint8
	head          *Node
	tail          *Node
	len           int
}

func NewBuffer() *DoublyLinkedd {
	cb := &DoublyLinkedd{
		inputChannel:  make(chan []uint8),
		outputChannel: make(chan []uint8),
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
		} else {
			select {
			case val := <-cb.inputChannel:
				node := &Node{data: val, prev: cb.tail}
				cb.tail.next = node
				cb.tail = node
			case cb.outputChannel <- cb.head.data:
				if cb.head == cb.tail {
					cb.head, cb.tail = nil, nil
				} else {
					cb.head = cb.head.next
					cb.head.prev = nil
				}
			}
		}
	}
}

func (cb *DoublyLinkedd) Write(val []uint8) {
	cb.inputChannel <- val
}

func (cb *DoublyLinkedd) Read() []uint8 {
	return <-cb.outputChannel
}
