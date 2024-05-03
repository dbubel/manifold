package topics

import (
	"context"
	"sync"
	"testing"

	"github.com/dbubel/manifold/pkg/logging"
)

func BenchmarkTopics_AsyncEnqueueDequeue(b *testing.B) {
	topics1 := NewTopics(logging.New(logging.DEBUG))
	// topics2 := NewTopics(logging.New(logging.DEBUG))
	// topics3 := NewTopics(logging.New(logging.DEBUG))
	// topics4 := NewTopics(logging.New(logging.DEBUG))

	s := GenerateRandomString(1000)
	var n int
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		n++
		go func(a int) {
			wg.Add(1)
			topics1.Dequeue(context.Background(), testTopic)
			wg.Done()
		}(i)
		//go func(a int) {
		//	wg.Add(1)
		//	topics2.Dequeue(context.Background(), testTopic)
		//	wg.Done()
		//}(i)
		//go func(a int) {
		//	wg.Add(1)
		//	topics3.Dequeue(context.Background(), testTopic)
		//	wg.Done()
		//}(i)
		//go func(a int) {
		//	wg.Add(1)
		//	topics4.Dequeue(context.Background(), testTopic)
		//	wg.Done()
		//}(i)
	}

	for i := 0; i < b.N; i++ {
		go func(a int) {
			wg.Add(1)
			topics1.Enqueue(testTopic, s)
			wg.Done()
		}(i)
		//go func(a int) {
		//	wg.Add(1)
		//	topics2.Enqueue(testTopic, s)
		//	wg.Done()
		//}(i)
		//go func(a int) {
		//	wg.Add(1)
		//	topics3.Enqueue(testTopic, s)
		//	wg.Done()
		//}(i)
		//go func(a int) {
		//	wg.Add(1)
		//	topics4.Enqueue(testTopic, s)
		//	wg.Done()
		//}(i)
	}
	wg.Wait()
}
func BenchmarkNewTopics(b *testing.B) {
	topics := NewTopics(logging.New(logging.DEBUG))
	s := GenerateRandomString(1000)
	for i := 0; i < b.N; i++ {
		topics.Enqueue(testTopic, s)
		topics.Dequeue(context.Background(), testTopic)
	}
}
