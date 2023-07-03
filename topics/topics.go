package topics

import (
	"context"
	"github.com/dbubel/manifold/buffer"
	"sync"
)

type Topics struct {
	topics map[string]*buffer.Queue
	lock   sync.RWMutex
}

func NewTopics() *Topics {
	return &Topics{
		topics: make(map[string]*buffer.Queue),
	}
}

func (t *Topics) Enqueue(topicName string, data []byte) {
	t.lock.Lock()
	if _, ok := t.topics[topicName]; !ok {
		t.topics[topicName] = buffer.NewQueue()
	}
	t.lock.Unlock()
	t.topics[topicName].Enqueue(data)
}

func (t *Topics) Dequeue(ctx context.Context, topicName string) []byte {
	t.lock.Lock()
	if _, ok := t.topics[topicName]; !ok {
		t.topics[topicName] = buffer.NewQueue()
	}
	t.lock.Unlock()
	return t.topics[topicName].BlockingDequeue(ctx)
}

func (t *Topics) Len(topicName string) int {
	t.lock.Lock()
	if _, ok := t.topics[topicName]; !ok {
		t.topics[topicName] = buffer.NewQueue()
	}
	t.lock.Unlock()
	return t.topics[topicName].Len()
}
