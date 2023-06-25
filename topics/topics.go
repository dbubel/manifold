package topics

import (
	"context"
	"fmt"
	q "github.com/dbubel/manifold/queue"
)

// Topics In Go, a map is a reference type, so you don't need to use pointer semantics to access the map.
type Topics map[string]*q.Queue

func (t Topics) Enqueue(id string, value []uint8) error {
	queue, ok := t[id]
	if !ok {
		// Create a new queue if it doesn't exist yet
		queue = q.NewQueue()
		t[id] = queue
	}
	return queue.Enqueue(value)
}

func (t Topics) List() map[string]int32 {
	var m map[string]int32
	m = make(map[string]int32)
	for k, v := range t {
		m[k] = int32(v.Len())
	}
	return m
}

func (t Topics) Dequeue(id string) ([]byte, error) {
	queue, ok := t[id]
	if !ok {
		// Create a new queue if it doesn't exist yet
		queue = q.NewQueue()
		t[id] = queue
	}
	return queue.Dequeue()
}

func (t Topics) BlockingDequeue(ctx context.Context, id string) ([]byte, error) {
	queue, ok := t[id]
	if !ok {
		// Create a new queue if it doesn't exist yet
		queue = q.NewQueue()
		t[id] = queue
	}
	return queue.BlockingDequeue(ctx), nil
}

func (t Topics) Len(id string) (int, error) {
	queue, ok := t[id]
	if !ok {
		// There is no queue with this id
		return 0, fmt.Errorf("queue does not exist")
	}
	return queue.Len(), nil
}
