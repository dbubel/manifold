package topics

import (
	"context"
	"github.com/dbubel/manifold/pkg/logging"
	"github.com/dbubel/manifold/queue"
	"sync"
)

type Topics struct {
	topics map[string]*queue.Queue
	lock   sync.RWMutex
	log    *logging.Logger
}

func NewTopics(l *logging.Logger) *Topics {
	return &Topics{
		topics: make(map[string]*queue.Queue),
		log:    l,
	}
}

func (t *Topics) Enqueue(topicName string, data []byte) {
	topic := t.GetOrCreateTopic(topicName)
	topic.Enqueue(data)
}

func (t *Topics) EnqueueHighPriority(topicName string, data []byte) {
	topic := t.GetOrCreateTopic(topicName)
	topic.EnqueueHighPriority(data)
}

func (t *Topics) Dequeue(ctx context.Context, topicName string) []byte {
	topic := t.GetOrCreateTopic(topicName)
	return topic.BlockingDequeue(ctx)
}

func (t *Topics) Len(topicName string) int {
	topic := t.GetOrCreateTopic(topicName)
	return topic.Len()
}

//func (t *Topics) ShutdownTopics() {
//	for k, v := range t.topics {
//		t.log.WithFields(map[string]interface{}{"topic": k}).Info("topic shutting down")
//		v.Shutdown()
//	}
//}

//func (t *Topics) ShutdownTopic(topic string) {
//	if v, ok := t.topics[topic]; ok {
//		t.log.WithFields(map[string]interface{}{"topic": topic}).Info("topic shutting down")
//		v.Shutdown()
//	}
//}

func (t *Topics) ListTopics() {

}

func (t *Topics) DeleteTopic(topicName string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	//t.ShutdownTopic(topicName)
	delete(t.topics, topicName)
}

func (t *Topics) GetOrCreateTopic(topicName string) *queue.Queue {
	t.lock.Lock()
	q, ok := t.topics[topicName]
	if !ok {
		q = queue.NewQueue(t.log)
		t.topics[topicName] = q
	}
	t.lock.Unlock()
	return q
}

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
//			t.GetOrCreateTopic(req.topicName).Enqueue(req.data)
//		case req := <-t.dequeue:
//			queue := t.GetOrCreateTopic(req.topicName)
//			req.responseChan <- queue.BlockingDequeue(req.ctx)
//		case req := <-t.lenReq:
//			queue := t.GetOrCreateTopic(req.topicName)
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
//func (t *Topics) ShutdownTopics() {
//	close(t.shutdown)
//}
//
//func (t *Topics) GetOrCreateTopic(topicName string) *buffer.Queue {
//	queue, ok := t.topics[topicName]
//	if !ok {
//		queue = buffer.NewQueue()
//		t.topics[topicName] = queue
//	}
//	return queue
//}
