package shards

import (
	"context"
	"testing"
	"time"
)

func TestShardedDataBasic(t *testing.T) {
	data := NewShardedTopics(1)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := data.Enqueue("test", []byte("Hello World!"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	dq, err := data.Dequeue(ctx, "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(dq) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(dq))
	}
}

func TestShardedDataMultipleShards(t *testing.T) {
	data := NewShardedTopics(2)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cancel()

	err := data.Enqueue("test", []byte("Hello World!"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	//time.Sleep(time.Millisecond * 100)
	dq, err := data.Dequeue(ctx, "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(dq) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(dq))
	}
}

func TestShardedDataBasicAsync(t *testing.T) {
	data := NewShardedTopics(1)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := data.Enqueue("test", []byte("Hello World!"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	dq, err := data.Dequeue(ctx, "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(dq) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(dq))
	}
}
