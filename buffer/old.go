package buffer

import (
	"context"
	"github.com/dbubel/manifold/linked_list"
	"sync"
)

type Queue struct {
	list *linked_list.List
	lock sync.Mutex
	Cond *sync.Cond
	//compression *compression.SnappyCompressor
}

func NewQueue() *Queue {
	q := &Queue{
		list: linked_list.New(),
		//compression: compression.NewSnappy(),
	}
	q.Cond = sync.NewCond(&q.lock)
	return q
}

func (q *Queue) Enqueue(value []uint8) {
	q.lock.Lock()
	defer q.lock.Unlock()

	//cmpValue, err := q.compression.CompressIOPool(value)

	q.list.PushBack(value)
	q.Cond.Signal()
	//q.Cond.Broadcast()
}

func (q *Queue) Dequeue() []uint8 {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.list.Len() == 0 {
		return nil
	}

	element := q.list.Front()
	q.list.Remove(element)

	return element.Value
}

func (q *Queue) BlockingDequeue(ctx context.Context) []uint8 {
	q.lock.Lock()
	defer q.lock.Unlock()

	for q.list.Len() == 0 {
		q.Cond.Wait()
	}

	if ctx.Err() != nil {
		return nil
	}

	element := q.list.Front()
	val := element.Value
	//val, _ := q.compression.DecompressIOPool(element.Value)
	q.list.Remove(element)

	return val
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.list.Len()
}
