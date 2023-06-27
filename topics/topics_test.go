package topics

import (
	"bytes"
	"testing"
)

func TestQueues(t *testing.T) {
	topics := New()

	t.Run("Test enqueue and dequeue", func(t *testing.T) {
		topics.Enqueue("queue1", []byte("Hello"))
		value := topics.Dequeue("queue1")
		//fmt.Println(value)
		if value == nil {
			t.Error("Expected 'Hello', got nil")
			return
		}
		if !bytes.Equal(value, []byte("Hello")) {
			t.Errorf("Expected 'Hello', got '%v'", string(value))
		}
	})

	t.Run("Test enqueue and dequeue multiple topics", func(t *testing.T) {
		topics.Enqueue("queue1", []byte("Hello"))
		topics.Enqueue("queue1", []byte("World"))
		value := topics.Dequeue("queue1")

		if value == nil {
			t.Error("Expected 'Hello', got nil")
			return
		}
		if !bytes.Equal(value, []byte("Hello")) {
			t.Errorf("Expected 'Hello', got '%v'", string(value))
		}

		topics.Enqueue("queue2", []byte("Hello"))
		value = topics.Dequeue("queue2")
		if value == nil {
			t.Error("Expected 'Hello', got nil")
			return
		}
		if !bytes.Equal(value, []byte("Hello")) {
			t.Errorf("Expected 'Hello', got '%v'", string(value))
		}
	})
}
