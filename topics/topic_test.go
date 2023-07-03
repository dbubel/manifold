package topics

import "testing"

func TestTopic_Enqueue(t *testing.T) {
	topic := newTopic("topic")
	topic.inputChannel <- []byte("hello")
	topic.Enqueue([]byte("hello"))
	data := topic.Dequeue()
	t.Log(string(data))
}

func TestTopic_Len(t *testing.T) {
	topic := newTopic("topic")
	topic.Enqueue([]byte("hello"))
	t.Log(topic.Len())
}
