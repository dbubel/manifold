package queue

//
//func TestQueue(t *testing.T) {
//	queue := NewQueue()
//
//	t.Run("Test enqueue and dequeue", func(t *testing.T) {
//		queue.Enqueue([]byte("Hello"))
//		value, err := queue.Dequeue()
//		if err != nil {
//			t.Error(err)
//		}
//		if !bytes.Equal(value, []byte("Hello")) {
//			t.Errorf("Expected '%v, got '%v'", []byte("Hello"), value)
//		}
//	})
//
//	t.Run("Test length", func(t *testing.T) {
//		queue.Enqueue([]byte("World"))
//		len := queue.Len()
//		if len != 1 {
//			t.Errorf("Expected 1, got %d", len)
//		}
//		queue.Dequeue() // empty the queue
//	})
//
//	t.Run("Test blocking dequeue", func(t *testing.T) {
//		go func() {
//			time.Sleep(time.Millisecond * 5)
//			queue.Enqueue([]byte("Go"))
//		}()
//		value := queue.BlockingDequeue(context.TODO())
//		if !bytes.Equal(value, []byte("Go")) {
//			t.Errorf("Expected 'Go', got '%s'", string(value))
//		}
//		queue.Dequeue() // empty the queue
//	})
//
//	t.Run("Test dequeue from empty queue", func(t *testing.T) {
//		_, err := queue.Dequeue()
//		if err == nil {
//			t.Error("Expected error, got nil")
//		}
//		if err.Error() != "queue is empty" {
//			t.Errorf("Expected 'queue is empty', got '%s'", err.Error())
//		}
//	})
//}
//
//func BenchmarkQueue(b *testing.B) {
//	q := NewQueue()
//
//	// Seed the random number generator
//	rand.Seed(time.Now().UnixNano())
//
//	// Enqueue random slices of bytes
//	for i := 0; i < b.N; i++ {
//		value := make([]uint8, rand.Intn(1000))
//		rand.Read(value)
//		q.Enqueue(value)
//	}
//
//	// Dequeue all elements
//	for i := 0; i < b.N; i++ {
//		q.Dequeue()
//	}
//}
