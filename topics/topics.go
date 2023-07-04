//package topics
//
//import (
//	"context"
//	"github.com/dbubel/manifold/buffer"
//	"sync"
//)
//
//type Topics struct {
//	topics   map[string]*buffer.Queue
//	enqueue  chan enqueueRequest
//	dequeue  chan dequeueRequest
//	lenReq   chan lenRequest
//	shutdown chan struct{}
//}
//
//type enqueueRequest struct {
//	topicName string
//	data      []byte
//}
//
//type dequeueRequest struct {
//	topicName    string
//	responseChan chan []byte
//	ctx          context.Context
//}
//
//type lenRequest struct {
//	topicName    string
//	responseChan chan int
//}
//
//func NewTopics() *Topics {
//	t := &Topics{
//		topics:   make(map[string]*buffer.Queue),
//		enqueue:  make(chan enqueueRequest),
//		dequeue:  make(chan dequeueRequest),
//		lenReq:   make(chan lenRequest),
//		shutdown: make(chan struct{}),
//	}
//	go t.start()
//	return t
//}
//
//func (t *Topics) start() {
//	for {
//		select {
//		case req := <-t.enqueue:
//			t.getOrCreateQueue(req.topicName).Enqueue(req.data)
//		case req := <-t.dequeue:
//			queue := t.getOrCreateQueue(req.topicName)
//			req.responseChan <- queue.BlockingDequeue(req.ctx)
//		case req := <-t.lenReq:
//			queue := t.getOrCreateQueue(req.topicName)
//			req.responseChan <- queue.Len()
//		case <-t.shutdown:
//			close(t.enqueue)
//			close(t.dequeue)
//			close(t.lenReq)
//			return
//		}
//	}
//}
//
//func (t *Topics) Enqueue(topicName string, data []byte) {
//	t.enqueue <- enqueueRequest{
//		topicName: topicName,
//		data:      data,
//	}
//}
//
//func (t *Topics) Dequeue(ctx context.Context, topicName string) []byte {
//	responseChan := make(chan []byte)
//	t.dequeue <- dequeueRequest{
//		topicName:    topicName,
//		responseChan: responseChan,
//		ctx:          ctx,
//	}
//	select {
//	case data := <-responseChan:
//		return data
//	case <-ctx.Done():
//		return nil
//	}
//}
//
//func (t *Topics) Len(topicName string) int {
//	responseChan := make(chan int)
//	t.lenReq <- lenRequest{
//		topicName:    topicName,
//		responseChan: responseChan,
//	}
//	return <-responseChan
//}
//
//func (t *Topics) Shutdown() {
//	close(t.shutdown)
//}
//
//func (t *Topics) getOrCreateQueue(topicName string) *buffer.Queue {
//	queue, ok := t.topics[topicName]
//	if !ok {
//		queue = buffer.NewQueue()
//		t.topics[topicName] = queue
//	}
//	return queue
//}

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
