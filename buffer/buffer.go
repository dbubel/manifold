package buffer

type Node struct {
	data interface{}
	prev *Node
	next *Node
	len  int
}

type DoublyLinkedd struct {
	inputChannel  chan interface{}
	outputChannel chan interface{}
	head          *Node
	tail          *Node
}

func DoublyLinked() *DoublyLinkedd {
	cb := &DoublyLinkedd{
		inputChannel:  make(chan interface{}),
		outputChannel: make(chan interface{}),
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

func (cb *DoublyLinkedd) Write(val interface{}) {
	cb.inputChannel <- val
}

func (cb *DoublyLinkedd) Read() interface{} {
	return <-cb.outputChannel
}
