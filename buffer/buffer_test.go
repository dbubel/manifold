package buffer

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQueue_BlockingDequeue(t *testing.T) {
	q := NewQueue()

	go func() {
		time.Sleep(time.Millisecond * 50)
		q.Enqueue([]byte("hello"))
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	t.Log(q.BlockingDequeue(ctx))
}

func TestNewQueue(t *testing.T) {
	t.Run("test async enqueue async dequeue", func(t *testing.T) {
		topics := NewQueue()
		//topics.AddTopic(topicOne)
		var wg sync.WaitGroup
		var results sync.Map

		for i := 0; i < 100; i++ {
			go func(a int) {
				val := topics.BlockingDequeue(context.Background())
				results.Store(a, string(val))
				wg.Done()
			}(i)
		}

		for i := 0; i < 100; i++ {
			go func(a int) {
				wg.Add(1)
				topics.Enqueue([]byte(fmt.Sprintf("hello world %d", a)))
			}(i)
		}

		wg.Wait()

		//time.Sleep(time.Millisecond * 100)

		allPresent := true

		for i := 0; i < 100; i++ {
			_ = fmt.Sprintf("hello world %d", i)
			val, found := results.Load(i)
			_ = val
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
