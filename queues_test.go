package rank_one

import (
	"bytes"
	"context"
	"testing"
	"time"
)

func TestQueues(t *testing.T) {
	queues := make(Queues)

	t.Run("Test enqueue and dequeue", func(t *testing.T) {
		queues.Enqueue("queue1", []byte("Hello"))
		value, err := queues.Dequeue("queue1")
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(value, []byte("Hello")) {
			t.Errorf("Expected 'Hello', got '%v'", string(value))
		}
	})

	t.Run("Test length", func(t *testing.T) {
		queues.Enqueue("queue1", []byte("World"))
		len, err := queues.Len("queue1")
		if err != nil {
			t.Error(err)
		}

		if len != 1 {
			t.Errorf("Expected 1, got %d", len)
		}
	})

	t.Run("Test blocking dequeue", func(t *testing.T) {
		queues.Dequeue("queue1") // empty the queue
		go func() {
			time.Sleep(time.Millisecond * 5)
			queues.Enqueue("queue1", []byte("Go"))
		}()

		value, err := queues.BlockingDequeue(context.TODO(), "queue1")
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(value, []byte("Go")) {
			t.Errorf("Expected 'Go', got '%v'", value)
		}
		queues.Dequeue("queue1") // empty the queue
	})

	t.Run("Test dequeue from empty queue", func(t *testing.T) {
		_, err := queues.Dequeue("queue1")
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if err.Error() != "queue is empty" {
			t.Errorf("Expected 'queue is empty', got '%s'", err.Error())
		}
	})

	t.Run("Test length of non-existent queue", func(t *testing.T) {
		_, err := queues.Len("queue2")
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if err.Error() != "queue does not exist" {
			t.Errorf("Expected 'queue does not exist', got '%s'", err.Error())
		}
	})
}
