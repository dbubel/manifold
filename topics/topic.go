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
	return &t
}

func (t *Topic) Enqueue(data []byte) {
	t.inputChannel <- data
}

func (t *Topic) Dequeue() []byte {
	return <-t.outputChannel
}

func (t *Topic) Len() int {
	return t.Queue.Len()
}
