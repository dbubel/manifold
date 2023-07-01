package topics

import "context"

type Topics struct {
	Topics map[string]*Topic
}

func New() *Topics {
	return &Topics{
		Topics: make(map[string]*Topic),
	}
}

func (t *Topics) AddTopic(name string) {
	t.Topics[name] = newTopic(name)
}

func (t *Topics) GetTopic(name string) *Topic {
	if _, exists := t.Topics[name]; exists {
		return t.Topics[name]
	}
	return nil
}

// Enqueue adds a value to the queue with the given id.
func (t *Topics) Enqueue(id string, value []uint8) {
	topic := t.GetTopic(id)
	if topic == nil {
		// Create a new queue if it doesn't exist yet
		t.AddTopic(id)
	}
	topic = t.GetTopic(id)
	topic.Queue.Write(value)
}

func (t *Topics) Dequeue(ctx context.Context, id string) []uint8 {
	topic := t.GetTopic(id)
	if topic == nil {
		// Create a new queue if it doesn't exist yet
		t.AddTopic(id)
	}
	topic = t.GetTopic(id)
	return topic.Queue.Read(ctx)
}

func (t Topics) List(ctx context.Context) map[string]int32 {
	var result = make(map[string]int32)
	for k, v := range t.Topics {
		result[k] = int32(v.Queue.Len(ctx))
	}
	return result
}

//
//func (t Topics) Dequeue(id string) ([]byte, error) {
//	queue, ok := t[id]
//	if !ok {
//		// Create a new queue if it doesn't exist yet
//		queue = q.NewQueue()
//		t[id] = queue
//	}
//	return queue.Dequeue()
//}
//
//func (t Topics) BlockingDequeue(ctx context.Context, id string) ([]byte, error) {
//	queue, ok := t[id]
//	if !ok {
//		// Create a new queue if it doesn't exist yet
//		queue = q.NewQueue()
//		t[id] = queue
//	}
//	return queue.BlockingDequeue(ctx), nil
//}
//
//func (t Topics) Len(id string) (int, error) {
//	queue, ok := t[id]
//	if !ok {
//		// There is no queue with this id
//		return 0, fmt.Errorf("queue does not exist")
//	}
//	return queue.Len(), nil
//}
