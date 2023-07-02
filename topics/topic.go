package topics

import (
	"github.com/dbubel/manifold/buffer"
)

type Topic struct {
	Name          string
	Queue         *buffer.DoublyLinkedd
	inputChannel  chan []byte
	outputChannel chan []byte
}

func newTopic(name string) *Topic {
	t := Topic{
		Name:          name,
		inputChannel:  make(chan []byte),
		outputChannel: make(chan []byte),
	}

	t.Queue = buffer.NewBuffer(t.inputChannel, t.outputChannel)
	go t.run()
	return &t
}

func (t *Topic) run() {
	for {
		select {
		case val := <-t.inputChannel:
			t.Queue.Write(val)
			//case t.outputChannel <- t.Queue.Read(context.Background()):
		}
	}
}

func (t *Topic) Enqueue(data []byte) {
	t.inputChannel <- data
}

func (t *Topic) Dequeue() []byte {
	return <-t.outputChannel
	//timer := time.NewTicker(time.Second)
	//select {
	//case val := <-t.outputChannel:
	//	fmt.Println("HI", string(val))
	//	return []byte{1}
	//case <-timer.C:
	//	return nil
	//}
	//return
}
