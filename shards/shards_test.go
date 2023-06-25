package shards

import (
	"testing"
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
