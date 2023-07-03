package topics

import (
	"bytes"
	"context"
	"testing"
	"time"
)

func TestTopics(t *testing.T) {
	topics := New()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	t.Run("Test enqueue and dequeue", func(t *testing.T) {
		topics.Enqueue("queue1", []byte("Hello"))
		value := topics.Dequeue(ctx, "queue1")
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
		value := topics.Dequeue(ctx, "queue1")

		if value == nil {
			t.Error("Expected 'Hello', got nil")
			return
		}
		if !bytes.Equal(value, []byte("Hello")) {
			t.Errorf("Expected 'Hello', got '%v'", string(value))
		}

		topics.Enqueue("queue2", []byte("Hello"))
		value = topics.Dequeue(ctx, "queue2")
		if value == nil {
			t.Error("Expected 'Hello', got nil")
			return
		}
		if !bytes.Equal(value, []byte("Hello")) {
			t.Errorf("Expected 'Hello', got '%v'", string(value))
		}
	})
}
func TestTopicsAsync(t *testing.T) {
	topics := New()
	t.Run("async read write", func(t *testing.T) {
		for i := 0; i < 1; i++ {
			go func(a int) {
				topics.Enqueue("queue1", []byte("Hello"))
			}(i)
		}
	
		for i := 0; i < 1; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
			_ = cancel
			value := topics.Dequeue(ctx, "queue1")
			t.Log(string(value))
			//cancel()
		}
	})
}
