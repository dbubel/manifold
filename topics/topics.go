package topics

type Topics struct {
	addTopic      chan string
	input         chan TopicEnqueueWrapper
	output        chan []byte
	outputRequest chan string
	//listTopicsReq chan struct{}
	//listTopicsRes chan map[string]int
	Topics map[string]*Topic
}

func New() *Topics {
	t := Topics{
		Topics:        make(map[string]*Topic),
		addTopic:      make(chan string),
		input:         make(chan TopicEnqueueWrapper),
		output:        make(chan []byte),
		outputRequest: make(chan string),
		//listTopicsReq: make(chan struct{}),
		//listTopicsRes: make(chan map[string]int),
	}

	go t.run()
	return &t
}

type TopicEnqueueWrapper struct {
	TopicName string
	Data      []byte
}

func (t *Topics) run() {
	for {
		select {
		//case <-t.listTopicsReq:
		//topicsList := make(map[string]int)
		//for k,v:=range t.Topics{
		//
		//}
		//t.listTopicsRes <- make(map[string]int)
		case topicName := <-t.addTopic:
			if _, ok := t.Topics[topicName]; !ok {
				t.Topics[topicName] = newTopic(topicName)
			}
		case val := <-t.input:
			if _, ok := t.Topics[val.TopicName]; !ok {
				t.Topics[val.TopicName] = newTopic(val.TopicName)
			}
			t.Topics[val.TopicName].Enqueue(val.Data)
		case topicName := <-t.outputRequest:
			//if _, ok := t.Topics[topicName]; !ok {
			//	t.addTopic <- topicName
			//}
			x := t.Topics[topicName].Dequeue()
			t.output <- x
		}
	}
}

//func (t *Topics) ListTopics() map[string]int {
//	t.listTopicsReq <- struct{}{}
//	return <-t.listTopicsRes
//}

func (t *Topics) AddTopic(topicName string) {
	t.addTopic <- topicName
}

func (t *Topics) Enqueue(topicName string, data []byte) {
	t.input <- TopicEnqueueWrapper{
		TopicName: topicName,
		Data:      data,
	}
}

func (t *Topics) Dequeue(topicName string) []byte {
	t.outputRequest <- topicName
	return <-t.output
}

//
//func (t *Topics) AddTopic(name string) {
//	t.m.Lock()
//	defer t.m.Unlock()
//	t.Topics[name] = newTopic(name)
//}

//func (t *Topics) GetTopic(name string) *Topic {
//	t.m.RLock()
//	defer t.m.RUnlock()
//	if _, exists := t.Topics[name]; exists {
//		return t.Topics[name]
//	}
//	return nil
//}
//
//func (t *Topics) Enqueue(topicName string, data []byte) {
//	if _, exists := t.Topics[topicName]; !exists {
//		fmt.Println("create", topicName)
//		t.AddTopic(topicName)
//	}
//	t.Topics[topicName].Queue.Write(data)
//}
//
//func (t *Topics) Dequeue(ctx context.Context, topicName string) []byte {
//	//if _, exists := t.Topics[topicName]; !exists {
//	//	t.AddTopic(topicName)
//	//}
//	return t.Topics[topicName].Queue.Read(ctx)
//}

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
