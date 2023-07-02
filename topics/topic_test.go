package topics

import "testing"

func TestTopic_Enqueue(t *testing.T) {
	topic := newTopic("topic")
	topic.Enqueue([]byte("hello"))
	//data := topic.Dequeue()
	//t.Log(string(data))
}
