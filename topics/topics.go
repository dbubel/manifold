package topics

import (
	"fmt"
	"sync"
)

type Topics struct {
	addTopic      chan string
	input         chan TopicEnqueueWrapper
	output        chan []byte
	outputRequest chan string
	lenRequest    chan string
	lenResult     chan int
	Topics        map[string]*Topic
	lock          sync.Mutex
	Cond          *sync.Cond
}

func New() *Topics {
	t := Topics{
		Topics:        make(map[string]*Topic),
		addTopic:      make(chan string),
		input:         make(chan TopicEnqueueWrapper),
		output:        make(chan []byte),
		outputRequest: make(chan string),
		lenRequest:    make(chan string),
		lenResult:     make(chan int),
	}
	t.Cond = sync.NewCond(&t.lock)

	go t.run()
	return &t
}

type TopicEnqueueWrapper struct {
	TopicName string
	Data      []byte
}

func (t *Topics) run() {
	for {
		select {
		case topicName := <-t.lenRequest:
			t.lenResult <- t.Topics[topicName].Len()
		case topicName := <-t.addTopic:
			if _, ok := t.Topics[topicName]; !ok {
				t.Topics[topicName] = newTopic(topicName)
			}
		case val := <-t.input:
			if _, ok := t.Topics[val.TopicName]; !ok {
				t.Topics[val.TopicName] = newTopic(val.TopicName)
			}

			t.Topics[val.TopicName].Enqueue(val.Data)

		case topicName := <-t.outputRequest:
			if _, ok := t.Topics[topicName]; !ok {
				//t.Topics[topicName] = newTopic(topicName)

				val := <-t.input
				fmt.Println(val)
				//t.Topics[val.TopicName].Enqueue(val.Data)
				fmt.Println("return nil")
				t.output <- val.Data
			} else {
				x := t.Topics[topicName].Dequeue()
				t.output <- x
			}
		}
	}
}

func (t *Topics) AddTopic(topicName string) {
	t.addTopic <- topicName
}

func (t *Topics) Enqueue(topicName string, data []byte) {
	t.input <- TopicEnqueueWrapper{
		TopicName: topicName,
		Data:      data,
	}
}

func (t *Topics) Len(topicName string) int {
	t.lenRequest <- topicName
	return <-t.lenResult
}

func (t *Topics) Dequeue(topicName string) []byte {
	t.outputRequest <- topicName
	return <-t.output
}
