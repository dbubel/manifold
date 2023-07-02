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
	t.m.Lock()
	defer t.m.Unlock()
	t.Topics[name] = newTopic(name)
}

//func (t *Topics) GetTopic(name string) *Topic {
//	t.m.RLock()
//	defer t.m.RUnlock()
//	if _, exists := t.Topics[name]; exists {
//		return t.Topics[name]
//	}
//	return nil
//}

func (t *Topics) Enqueue(topicName string, data []byte) {
	if _, exists := t.Topics[topicName]; !exists {
		fmt.Println("create", topicName)
		t.AddTopic(topicName)
	}
	t.Topics[topicName].Queue.Write(data)
}

func (t *Topics) Dequeue(ctx context.Context, topicName string) []byte {
	//if _, exists := t.Topics[topicName]; !exists {
	//	t.AddTopic(topicName)
	//}
	return t.Topics[topicName].Queue.Read(ctx)
}

//
//// Enqueue adds a value to the queue with the given id.
//func (t *Topics) Enqueue(id string, value []uint8) {
//	t.m.Lock()
//	defer t.m.Unlock()
//	topic := t.GetTopic(id)
//	if topic == nil {
//		// Create a new queue if it doesn't exist yet
//		t.AddTopic(id)
//	}
//	topic = t.GetTopic(id)
//	topic.Queue.Write(value)
//}
//
//func (t *Topics) Dequeue(ctx context.Context, id string) []uint8 {
//	t.m.Lock()
//	defer t.m.Unlock()
//	topic := t.GetTopic(id)
//	if topic == nil {
//		// Create a new queue if it doesn't exist yet
//		t.AddTopic(id)
//	}
//	topic = t.GetTopic(id)
//	return topic.Queue.Read(ctx)
//}

//
//func (t Topics) List(ctx context.Context) map[string]int32 {
//	var result = make(map[string]int32)
//	for k, v := range t.Topics {
//		result[k] = int32(v.Queue.Len(ctx))
//	}
//	return result
//}
