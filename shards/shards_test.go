package shards

import "testing"

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
	data := NewShardedTopics(10)
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

//
//func TestShardedBlockingDequeueSingleShard(t *testing.T) {
//	data := NewShardedTopics(1)
//	go func() {
//		time.Sleep(time.Second)
//		err := data.Enqueue("test", []byte("Hello World!"))
//		if err != nil {
//			t.Errorf("Error: %v", err)
//		}
//	}()
//
//	d, err := data.BlockingDequeue(context.Background(), "test")
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//
//	if string(d) != "Hello World!" {
//		t.Errorf("Expected: %v, got: %v", "test", string(d))
//	}
//}
//
//// this will pass and fail until i fix the blocking dequeue
//// to where the lock / cond is passed in.
//func TestShardedBlockingDequeueMultipleShards(t *testing.T) {
//	data := NewShardedTopics(5)
//	go func() {
//		time.Sleep(time.Second)
//		err := data.Enqueue("test", []byte("Hello World!"))
//		if err != nil {
//			t.Errorf("Error: %v", err)
//		}
//	}()
//
//	d, err := data.BlockingDequeue(context.Background(), "test")
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//
//	if string(d) != "Hello World!" {
//		t.Errorf("Expected: %v, got: %v", "test", string(d))
//	}
//}
//
//func TestList(t *testing.T) {
//	data := NewShardedTopics(2)
//
//	err := data.Enqueue("topic_one", []byte("Hello World! one"))
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//	data.List()
//	t.Log("----")
//
//	err = data.Enqueue("topic_two", []byte("Hello World! two"))
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//	data.List()
//}
