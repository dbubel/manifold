package queue

import (
	"context"

	// list "github.com/dbubel/manifold/mempool"
	"github.com/dbubel/manifold/pkg/logging"
	"github.com/google/uuid"
)

// type ByteArray []byte
//
// func (array ByteArray) Erase() error {
// 	for i := 0; i < len(array); i++ {
// 		array[i] = 0
// 	}
//
// 	return nil
// }

type Queue struct {
	list                *List
	dequeued            map[uuid.UUID]*Element
	enqueue             chan *Element
	enqueueHighPriority chan *Element
	dequeue             chan chan *Element
	lenReq              chan chan int
	cancelFunc          context.CancelFunc
	ctx                 context.Context
	log                 *logging.Logger
}

func NewQueue(l *logging.Logger) *Queue {
	ctx, cancel := context.WithCancel(context.Background())

	q := &Queue{
		list:                New(),
		dequeued:            make(map[uuid.UUID]*Element),
		enqueue:             make(chan *Element),
		enqueueHighPriority: make(chan *Element),
		dequeue:             make(chan chan *Element),
		lenReq:              make(chan chan int),
		cancelFunc:          cancel,
		log:                 l,
	}
	go q.start(ctx)
	return q
}

func (q *Queue) start(ctx context.Context) {
	defer q.log.Info("queue stopped")
	for {
		select {
		case value := <-q.enqueueHighPriority:
			q.list.PushFront(value)
		case value := <-q.enqueue:
			q.list.PushBack(value)
		case responseChan := <-q.dequeue:
			if q.list.Len() == 0 {
				// Wait for an enqueue operation if the list is empty
				value := <-q.enqueue
				responseChan <- value
			} else {
				element := q.list.Front()
				q.AddDequeued(element)
				q.dequeued[element.ID] = element
				val := element
				q.list.Remove(element)
				responseChan <- val
			}
		case responseChan := <-q.lenReq:
			responseChan <- q.list.Len()
		case <-ctx.Done():
			q.log.Info("queue is shutting down...")
			close(q.enqueue)
			close(q.dequeue)
			close(q.lenReq)
			close(q.enqueueHighPriority)
			return
		}
	}
}

// BatchedDequeue TODO:(dean) implement ability to dequeue batches of messages
func (q *Queue) BatchedDequeue(ctx context.Context, n int) [][]uint8 {
	return [][]uint8{}
}

func (q *Queue) Enqueue(value *Element) {
	q.enqueue <- value
}

// func (q *Queue) EnqueueElement(value *Element) {
// 	q.log.Debug("in enqueueElement func")
// 	q.enqueueElement <- value
// }

func (q *Queue) EnqueueHighPriority(value *Element) {
	q.enqueueHighPriority <- value
}

func (q *Queue) BlockingDequeue(ctx context.Context) *Element {
	responseChan := make(chan *Element)
	select {
	case q.dequeue <- responseChan:
		return <-responseChan
	case <-ctx.Done():
		return nil
	}
}

func (q *Queue) Len() int {
	return q.list.len
}

//func (q *Queue) Shutdown() {
//	q.cancelFunc()
//	<-q.shutdown
//}

//package buffer
//
//import (
//	"context"
//	"sync"
//)
//
//type Queue struct {
//	list *List
//	lock sync.Mutex
//	Cond *sync.Cond
//}
//
//func NewQueue() *Queue {
//	q := &Queue{
//		list: New(),
//	}
//	q.Cond = sync.NewCond(&q.lock)
//	return q
//}
//
//func (q *Queue) Enqueue(value []uint8) {
//	q.lock.Lock()
//	defer q.lock.Unlock()
//
//	q.list.PushBack(value)
//	q.Cond.Signal()
//}
//
////func (q *Queue) Dequeue() []uint8 {
////	q.lock.Lock()
////	defer q.lock.Unlock()
////
////	if q.list.Len() == 0 {
////		return nil
////	}
////
////	element := q.list.Front()
////	q.list.Remove(element)
////
////	return element.Value
////}
//
//func (q *Queue) BlockingDequeue(ctx context.Context) []uint8 {
//	q.lock.Lock()
//	defer q.lock.Unlock()
//
//	for q.list.Len() == 0 {
//		q.Cond.Wait()
//	}
//
//	if ctx.Err() != nil {
//		return nil
//	}
//
//	element := q.list.Front()
//	val := element.Value
//	q.list.Remove(element)
//
//	return val
//}
//
//func (q *Queue) Len() int {
//	q.lock.Lock()
//	defer q.lock.Unlock()
//
//	return q.list.Len()
//}
