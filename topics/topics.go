package topics

import (
	"context"
	"fmt"
	"sync"
)

type Topics struct {
	m      sync.RWMutex
	Topics map[string]*Topic
}

func New() *Topics {
	return &Topics{
		Topics: make(map[string]*Topic),
	}
}

func (t *Topics) AddTopic(name string) {
	t.Topics[name] = newTopic(name)
}

func (t *Topics) Enqueue(topicName string, data []byte) {
	t.m.Lock()
	defer t.m.Unlock()
	if _, ok := t.Topics[topicName]; !ok {
		fmt.Println("create", topicName)
		t.AddTopic(topicName)
	}

	t.Topics[topicName].Queue.Write(data)
}

func (t *Topics) Dequeue(ctx context.Context, topicName string) []byte {
	t.m.Lock()
	defer t.m.Unlock()
	x := t.Topics[topicName].Queue.Read(ctx)
	return x
}
