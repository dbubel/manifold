package buffer

import (
	"math/rand"
	"testing"
	"time"
)

func TestCircularBuffer_Read(t *testing.T) {
	cb := DoublyLinked()

	cb.Write(1)
	cb.Write(2)
	cb.Write(3)

	if cb.Read() != 1 {
		t.Errorf("Expected 1, but got %v", cb.Read())
	}

	if cb.Read() != 2 {
		t.Errorf("Expected 2, but got %v", cb.Read())
	}

	if cb.Read() != 3 {
		t.Errorf("Expected 3, but got %v", cb.Read())
	}
}

func TestCircularBuffer_Write(t *testing.T) {
	cb := DoublyLinked()

	cb.Write(1)
	cb.Write(2)
	cb.Write(3)

	if cb.head.data != 1 {
		t.Errorf("Expected head data to be 1, but got %v", cb.head.data)
	}

	if cb.tail.data != 3 {
		t.Errorf("Expected tail data to be 3, but got %v", cb.tail.data)
	}
}

func TestCircularBuffer_ReadAndWrite(t *testing.T) {
	cb := DoublyLinked()

	cb.Write(1)
	cb.Write(2)
	cb.Read()
	cb.Write(3)
	cb.Write(4)

	if cb.Read() != 2 {
		t.Errorf("Expected 2, but got %v", cb.Read())
	}

	if cb.Read() != 3 {
		t.Errorf("Expected 3, but got %v", cb.Read())
	}

	if cb.Read() != 4 {
		t.Errorf("Expected 4, but got %v", cb.Read())
	}
}

func TestCircularBuffer_ReadBeforeWrite(t *testing.T) {
	cb := DoublyLinked()

	go func() {
		time.Sleep(time.Millisecond * 200)
		cb.Write(1)
	}()

	if cb.Read() != 1 {
		t.Errorf("Expected 1, but got %v", cb.Read())
	}
}

func BenchmarkQueue(b *testing.B) {
	q := DoublyLinked()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Write random slices of bytes
	for i := 0; i < b.N; i++ {
		value := make([]uint8, rand.Intn(1000))
		rand.Read(value)
		q.Write(value)
	}

	// Dequeue all elements
	for i := 0; i < b.N; i++ {
		q.Read()
	}
}
