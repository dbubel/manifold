package rank_one

import (
	"context"
	"fmt"
)

// Queues In Go, a map is a reference type, so you don't need to use pointer semantics to access the map.
type Queues map[string]*Queue

func (q Queues) Enqueue(id string, value []uint8) {
	queue, ok := q[id]
	if !ok {
		// Create a new queue if it doesn't exist yet
		queue = NewQueue()
		q[id] = queue
	}
	queue.Enqueue(value)
}

func (q Queues) List() map[string]int32 {
	var m map[string]int32
	m = make(map[string]int32)
	for k, v := range q {
		m[k] = int32(v.Len())
	}
	return m
}

func (q Queues) Dequeue(id string) ([]uint8, error) {
	queue, ok := q[id]
	if !ok {
		// Create a new queue if it doesn't exist yet
		queue = NewQueue()
		q[id] = queue
	}
	return queue.Dequeue()
}

func (q Queues) BlockingDequeue(ctx context.Context, id string) ([]uint8, error) {
	queue, ok := q[id]
	if !ok {
		// Create a new queue if it doesn't exist yet
		queue = NewQueue()
		q[id] = queue
	}
	return queue.BlockingDequeue(ctx), nil
}

func (q Queues) Len(id string) (int, error) {
	queue, ok := q[id]
	if !ok {
		// There is no queue with this id
		return 0, fmt.Errorf("queue does not exist")
	}
	return queue.Len(), nil
}
