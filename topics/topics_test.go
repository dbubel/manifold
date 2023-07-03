package topics

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestTopics_EnqueueDequeueSimple(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	topic := NewTopics()

	topic.Enqueue("test", []byte("hello test"))
	data := topic.Dequeue(ctx, "test")
	if string(data) != "hello test" {
		t.Errorf("expected 'hello test', got %s", string(data))
		return
	}
}

func TestTopics_EnqueueDequeueMultipleTopics(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	topic := NewTopics()

	topic.Enqueue("test", []byte("hello test"))
	data := topic.Dequeue(ctx, "test")
	if string(data) != "hello test" {
		t.Errorf("expected 'hello test', got %s", string(data))
		return
	}

	topic.Enqueue("test two", []byte("hello test two"))
	data = topic.Dequeue(ctx, "test two")
	if string(data) != "hello test two" {
		t.Errorf("expected 'hello test two', got %s", string(data))
		return
	}
}

func TestTopics_AsyncEnqueueDequeueMultipleTopics(t *testing.T) {
	topics := NewTopics()
	var wg sync.WaitGroup
	var results sync.Map

	for i := 0; i < 100; i++ {
		go func(a int) {
			val := topics.Dequeue(context.Background(), "topic_one")
			results.Store(string(val), a)
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func(a int) {
			val := topics.Dequeue(context.Background(), "topic_two")
			results.Store(string(val), a)
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func(a int) {
			wg.Add(1)
			topics.Enqueue("topic_one", []byte(fmt.Sprintf("hello world one %d", a)))
			wg.Done()
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func(a int) {
			wg.Add(1)
			topics.Enqueue("topic_two", []byte(fmt.Sprintf("hello world two %d", a)))
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

func BenchmarkTopics_AsyncEnqueueDequeue(b *testing.B) {
	topics := NewTopics()
	s := GenerateRandomString(1000)
	var n int

	for i := 0; i < b.N; i++ {
		n++
		go func(a int) {
			val := topics.Dequeue(context.Background(), testTopic)
			_ = val
		}(i)
	}

	for i := 0; i < b.N; i++ {
		go func(a int) {
			topics.Enqueue(testTopic, s)
		}(i)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(length int) []byte {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return []byte(string(s))
}

const testTopic = "test_topic"

func BenchmarkNewTopics(b *testing.B) {
	topics := NewTopics()
	s := GenerateRandomString(1000)
	for i := 0; i < b.N; i++ {
		topics.Enqueue(testTopic, s)
		topics.Dequeue(context.Background(), testTopic)
	}
}
