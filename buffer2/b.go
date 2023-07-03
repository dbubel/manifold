package buffer2

import (
	"context"
	"github.com/dbubel/manifold/list"
)

type Queue struct {
	items    *list.List
	PushBack chan []uint8
	FrontReq chan struct{}
	FrontRes chan *list.Element
}

// NewQueue creates a new instance of Queue.
func NewQueue() *Queue {
	q := Queue{
		items:    list.New(),
		PushBack: make(chan []uint8),
		FrontReq: make(chan struct{}),
		FrontRes: make(chan *list.Element),
	}
	go q.run()
	return &q
}

func (cb *Queue) run() {
	for {
		select {
		case val := <-cb.PushBack:
			cb.items.PushBack(val)
		case <-cb.FrontReq:
			front := cb.items.Front()
			value := front
			cb.items.Remove(front)
			cb.FrontRes <- value
		}
	}
}

func (cb *Queue) Enqueue(data []uint8) {
	cb.PushBack <- data
}

func (cb *Queue) Dequeue(ctx context.Context) *list.Element {
	cb.FrontReq <- struct{}{}
	select {
	case val := <-cb.FrontRes:
		return val
	case <-ctx.Done():
		return nil
	}
}
