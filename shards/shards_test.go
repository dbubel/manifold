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

	dq := data.Dequeue(ctx, "test")

	if string(dq) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(dq))
	}
}

func TestShardedDataMultipleShards(t *testing.T) {
	data := NewShardedTopics(10)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := data.Enqueue("test", []byte("Hello World!"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	//time.Sleep(time.Millisecond * 100)

	dq := data.Dequeue(ctx, "test")

	if string(dq) != "Hello World!" {
		t.Errorf("Expected: %v, got: %v", "test", string(dq))
	}
}

//func TestShardedDataBasicAsync(t *testing.T) {
//	data := NewShardedTopics(1)
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	defer cancel()
//
//	go func() {
//		time.Sleep(time.Millisecond * 100)
//		err := data.Enqueue("test", []byte("Hello World!"))
//		if err != nil {
//			t.Errorf("Error: %v", err)
//		}
//	}()
//
//	dq, err := data.Dequeue(ctx, "test")
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//	if string(dq) != "Hello World!" {
//		t.Errorf("Expected: %v, got: %v", "test", string(dq))
//	}
//}

//func TestShardedMultipleEnquqeDequeu(t *testing.T) {
//	shards := NewShardedTopics(2)
//	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
//	_ = ctx
//	defer cancel()
//
//	for i := 0; i < 10; i++ {
//		go func(a int) {
//			//time.Sleep(time.Millisecond * 100)
//			shards.Enqueue("test", []byte(fmt.Sprintf("Hello World! %d", a)))
//		}(i)
//	}
//	time.Sleep(time.Millisecond * 10000)
//
//	for i := 0; i < 10; i++ {
//
//		time.Sleep(time.Millisecond * 100)
//		dq, err := shards.Dequeue(ctx, "test")
//
//		if err != nil {
//			t.Errorf("Error: %v", err)
//		}
//
//		fmt.Println(string(dq))
//		//if string(dq) != fmt.Sprintf("Hello World! %d", a) {
//		//	t.Errorf("Expected: %v, got: %v", "test", string(dq))
//		//}
//
//	}
//}
