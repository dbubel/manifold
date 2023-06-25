package shards

import (
	"context"
	"testing"
	"time"
)

func TestShardedDataBasic(t *testing.T) {
	data := NewShardedTopics(1)
	err := data.Enqueue("test", []byte("Hello World!"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	dq, err := data.Dequeue("test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(dq) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(dq))
	}
}

func TestShardedDataTwoShards(t *testing.T) {
	data := NewShardedTopics(20)
	err := data.Enqueue("test", []byte("Hello World!"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	dq, err := data.Dequeue("test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(dq) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(dq))
	}
}

func TestShardedBlockingDequeue(t *testing.T) {
	data := NewShardedTopics(1)
	go func() {
		time.Sleep(time.Second)
		err := data.Enqueue("test", []byte("Hello World!"))
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	}()

	d, err := data.BlockingDequeue(context.Background(), "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(d) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(d))
	}
}

func TestShardedBlockingDequeueMultipleShards(t *testing.T) {
	data := NewShardedTopics(5)
	go func() {
		time.Sleep(time.Second)
		err := data.Enqueue("test", []byte("Hello World!"))
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	}()

	d, err := data.BlockingDequeue(context.Background(), "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(d) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(d))
	}
}
