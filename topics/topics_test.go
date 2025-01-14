package topics

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/dbubel/manifold/pkg/logging"
	"github.com/dbubel/manifold/queue"
	"github.com/google/uuid"
)

func TestTopics_EnqueueDequeueSimple(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	id := uuid.New()
	element := queue.Element{
		Value:       []byte("dean"),
		ID:          id,
		EnqueueTime: time.Now(),
		Complete:    make(chan struct{}),
	}
	defer cancel()

	topic1 := NewTopics(logging.New(logging.DEBUG))

	topic1.Enqueue("test", &element)

	item := topic1.Dequeue(ctx, "test")
	topic1.AckDequeue("test", item.ID)
	fmt.Println(item)
}

func TestTopics_EnqueueDequeueMultipleTopics(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	topic := NewTopics(logging.New(logging.DEBUG))

	id := uuid.New()
	element := queue.Element{
		Value:       []byte("test"),
		ID:          id,
		EnqueueTime: time.Now(),
		Complete:    make(chan struct{}),
	}
	topic.Enqueue("test", &element)
	data := topic.Dequeue(ctx, "test")
	if !bytes.Equal(data.Value, element.Value) {
		t.Errorf("values not equal")
		return
	}

	element.Value = []byte("test2")
	topic.Enqueue("test2", &element)
	data = topic.Dequeue(ctx, "test2")
	if !bytes.Equal(data.Value, element.Value) {
		t.Errorf("value 2 not equal [%s] [%s]", data.Value, element.Value)
		return
	}
}

func TestTopics_AsyncEnqueueDequeueMultipleTopics(t *testing.T) {
	topics := NewTopics(logging.New(logging.DEBUG))
	var wg sync.WaitGroup
	var results sync.Map
	id := uuid.New()

	for i := 0; i < 100; i++ {
		go func(a int) {
			val := topics.Dequeue(context.Background(), "topic_one")
			results.Store(val, a)
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func(a int) {
			val := topics.Dequeue(context.Background(), "topic_two")
			results.Store(val, a)
		}(i)
	}

	for i := 0; i < 100; i++ {
		element := queue.Element{
			Value:       []byte("test"),
			ID:          id,
			EnqueueTime: time.Now(),
			Complete:    make(chan struct{}),
		}
		go func(a int) {
			wg.Add(1)
			topics.Enqueue("topic_one", &element)
			wg.Done()
		}(i)
	}

	for i := 0; i < 100; i++ {
		element := queue.Element{
			Value:       []byte("test"),
			ID:          id,
			EnqueueTime: time.Now(),
			Complete:    make(chan struct{}),
		}
		go func(a int) {
			wg.Add(1)
			topics.Enqueue("topic_two", &element)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Millisecond * 250)
	wg.Wait()

	i := 0
	results.Range(func(key, value any) bool {
		i++
		return true
	})

	if i != 200 {
		t.Errorf("expected 200 results, got %d", i)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(length int) []byte {
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return []byte(string(s))
}

//
// const testTopic = "test_topic"
//
// func TestTopics_DeleteTopic(t *testing.T) {
// 	topics := NewTopics(logging.New(logging.DEBUG))
// 	topics.Enqueue("test", []byte("hello test2"))
// 	topics.Enqueue("test", []byte("hello test1"))
// 	topics.DeleteTopic("test")
// 	x := topics.Len("test")
// 	if x != 0 {
// 		t.Errorf("topic len should be 0 got %d", x)
// 	}
//
// 	topics.Enqueue("test", []byte("new data"))
// 	x = topics.Len("test")
// 	if x != 1 {
// 		t.Errorf("topic len should be 1 got %d", x)
// 	}
// 	data := topics.Dequeue(context.Background(), "test")
//
// 	if string(data) != "new data" {
// 		t.Errorf("data invalid")
// 	}
// }
