package topics

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const topicOne = "TOPIC_ONE"

func TestTopics_AddTopic(t *testing.T) {
	topics := New()
	topics.AddTopic(topicOne)
}

func TestTopics_Enqueue(t *testing.T) {
	topics := New()
	topics.AddTopic(topicOne)
	topics.Enqueue(topicOne, []byte("hello world"))
}

func TestTopics_EnqueueDequeue(t *testing.T) {
	t.Run("test simple enqueue dequeue", func(t *testing.T) {
		topics := New()
		topics.AddTopic(topicOne)
		topics.Enqueue(topicOne, []byte("hello world"))
		val := topics.Dequeue(topicOne)
		if string(val) != "hello world" {
			t.Errorf("Expected: %v, got: %v", "hello world", string(val))
		}
	})

	//t.Run("test blocking dequeue", func(t *testing.T) {
	//	topics := New()
	//	topics.AddTopic(topicOne)
	//	go func() {
	//		time.Sleep(time.Millisecond * 100)
	//		topics.Enqueue(topicOne, []byte("hello world"))
	//	}()
	//
	//	val := topics.Dequeue(topicOne)
	//	if string(val) != "hello world" {
	//		t.Errorf("Expected: %v, got: %v", "hello world", string(val))
	//	}
	//})

	t.Run("test multiple enqueue dequeue", func(t *testing.T) {
		topics := New()
		topics.AddTopic(topicOne)
		for i := 0; i < 100; i++ {
			topics.Enqueue(topicOne, []byte(fmt.Sprintf("hello world %d", i)))
		}

		for i := 0; i < 100; i++ {
			val := topics.Dequeue(topicOne)
			t.Log(string(val))
		}
	})

	t.Run("test async enqueue serial dequeue", func(t *testing.T) {
		topics := New()
		topics.AddTopic(topicOne)
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			go func(a int) {
				wg.Add(1)
				topics.Enqueue(topicOne, []byte(fmt.Sprintf("hello world %d", a)))
				wg.Done()
			}(i)
		}
		wg.Wait()

		for i := 0; i < 100; i++ {
			val := topics.Dequeue(topicOne)
			t.Log(string(val))
		}
	})

	t.Run("test async enqueue async dequeue", func(t *testing.T) {
		topics := New()
		topics.AddTopic(topicOne)
		var wg sync.WaitGroup

		for i := 0; i < 100; i++ {
			go func(a int) {
				wg.Add(1)
				topics.Enqueue(topicOne, []byte(fmt.Sprintf("hello world %d", a)))
				wg.Done()
			}(i)
		}

		wg.Wait()

		var results sync.Map

		for i := 0; i < 100; i++ {
			go func(a int) {
				val := topics.Dequeue(topicOne)
				results.Store(a, string(val))
			}(i)
		}

		time.Sleep(time.Second)
		results.Range(func(key, value any) bool {
			//fmt.Println(key, value)
			return true
		})

		// Check if all expected strings are present
		allPresent := true

		for i := 0; i < 100; i++ {
			_ = fmt.Sprintf("hello world %d", i)
			_, found := results.Load(i)
			if !found {
				allPresent = false
				break
			}
		}

		if !allPresent {
			t.Error("Some strings are missing from the sync map")
			return
		}
	})
}
