package queue

import (
	"context"
	"fmt"
	"sync"

	"github.com/dbubel/manifold/compression"
	"github.com/dbubel/manifold/linked_list"
)

type Queue struct {
	list        *linked_list.List
	lock        sync.Mutex
	Cond        *sync.Cond
	compression *compression.SnappyCompressor
}

func NewQueue() *Queue {
	q := &Queue{
		list:        linked_list.New(),
		compression: compression.NewSnappy(),
	}
	q.Cond = sync.NewCond(&q.lock)
	return q
}

func (q *Queue) Enqueue(value []uint8) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	cmpValue, err := q.compression.Compress(value)
	if err != nil {
		return err
	}
	q.list.PushBack(cmpValue)
	q.Cond.Signal()

	return nil
}

func (q *Queue) Dequeue() (*linked_list.Element, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.list.Len() == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	element := q.list.Front()
	q.list.Remove(element)

	err := q.compression.Decompress(element)
	if err != nil {
		return nil, err
	}

	return element, nil
}

func (q *Queue) BlockingDequeue(ctx context.Context) *linked_list.Element {
	q.lock.Lock()
	defer q.lock.Unlock()

	for q.list.Len() == 0 {
		q.Cond.Wait()
	}

	if ctx.Err() != nil {
		return nil
	}

	element := q.list.Front()
	err := q.compression.Decompress(element)
	if err != nil {
		return nil
	}
	q.list.Remove(element)

	return element
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.list.Len()
}
