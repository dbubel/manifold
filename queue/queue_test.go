package queue

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/dbubel/manifold/pkg/logging"
	"github.com/google/uuid"
)

func TestQueue_BlockingDequeue(t *testing.T) {
	id := uuid.New()
	element := Element{
		Value:       []byte("dean"),
		ID:          id,
		EnqueueTime: time.Now(),
		Complete:    make(chan struct{}),
	}
	q := NewQueue(logging.New(logging.DEBUG))

	go func() {
		time.Sleep(time.Millisecond * 50)
		q.Enqueue(&element)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	item := q.BlockingDequeue(ctx)
	if item.ID != id {
		t.Error("id does not match", id, item.ID)
		t.FailNow()
	}
	if !bytes.Equal(item.Value, element.Value) {
		t.Error("values do not match", id, item.ID)
		t.FailNow()
	}
	if item.EnqueueTime.Unix() != element.EnqueueTime.Unix() {
		t.Error("timestamps do not match", id, item.ID)
		t.FailNow()
	}
}

func TestNewQueue(t *testing.T) {
	t.Run("test async enqueue async dequeue", func(t *testing.T) {
		q := NewQueue(logging.New(logging.DEBUG))

		id := uuid.New()
		element := Element{
			Value:       []byte("dean"),
			ID:          id,
			EnqueueTime: time.Now(),
			Complete:    make(chan struct{}),
		}
		var wg sync.WaitGroup
		var results sync.Map

		for i := 0; i < 100; i++ {
			go func(a int) {
				defer wg.Done()
				val := q.BlockingDequeue(context.Background())
				results.Store(a, val)
			}(i)
		}

		for i := 0; i < 100; i++ {
			go func(a int) {
				wg.Add(1)
				q.Enqueue(&element)
			}(i)
		}

		wg.Wait()

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
