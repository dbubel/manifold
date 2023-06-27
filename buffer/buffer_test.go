package buffer

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCircularBuffer_Read(t *testing.T) {
	cb := NewBuffer()

	cb.Write([]uint8{1})
	cb.Write([]uint8{2})
	cb.Write([]uint8{3})

	val := cb.Read()
	if !bytes.Equal(val, []uint8{1}) {
		t.Errorf("Expected %d, but got %d", []uint8{1}, val)
	}

	val = cb.Read()
	if !bytes.Equal(val, []uint8{2}) {
		t.Errorf("Expected %d, but got %d", []uint8{2}, val)
	}

	val = cb.Read()
	if !bytes.Equal(val, []uint8{3}) {
		t.Errorf("Expected %d, but got %d", []uint8{3}, val)
	}
}

func TestCircularBuffer_Write(t *testing.T) {
	cb := NewBuffer()

	cb.Write([]byte{1})
	cb.Write([]byte{2})
	cb.Write([]byte{3})

	if !bytes.Equal(cb.head.data, []byte{1}) {
		t.Errorf("Expected head data to be 1, but got %v", cb.head.data)
	}

	if !bytes.Equal(cb.tail.data, []byte{3}) {
		t.Errorf("Expected tail data to be 3, but got %v", cb.tail.data)
	}
}

func TestCircularBuffer_ReadAndWrite(t *testing.T) {
	cb := NewBuffer()

	cb.Write([]byte{1})
	cb.Write([]byte{2})
	cb.Read()
	cb.Write([]byte{3})
	cb.Write([]byte{4})

	if !bytes.Equal(cb.Read(), []byte{2}) {
		t.Errorf("Expected 2, but got %v", cb.Read())
	}

	if !bytes.Equal(cb.Read(), []byte{3}) {
		t.Errorf("Expected 3, but got %v", cb.Read())
	}

	if !bytes.Equal(cb.Read(), []byte{4}) {
		t.Errorf("Expected 4, but got %v", cb.Read())
	}
}

func TestCircularBuffer_ReadBeforeWrite(t *testing.T) {
	cb := NewBuffer()

	go func() {
		time.Sleep(time.Millisecond * 200)
		cb.Write([]byte{1})
	}()

	if !bytes.Equal(cb.Read(), []byte{1}) {
		t.Errorf("Expected 1, but got %v", cb.Read())
	}
}

func BenchmarkQueue(b *testing.B) {
	q := NewBuffer()

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

func TestCircularBuffer_ConcurrentReadWrite(t *testing.T) {
	cb := NewBuffer()
	wg := sync.WaitGroup{}

	res := []uint8{}

	for i := 0; i < 100; i++ {
		go func(a int) {
			wg.Add(1)
			cb.Write([]uint8{uint8(a)})
		}(i)
	}

	for i := 0; i < 100; i++ {
		go func() {
			res = append(res, cb.Read()[0])
			wg.Done()
		}()
	}

	fmt.Println(verifySlice(res))
	//t.Log(res)

}

func verifySlice(numbers []uint8) bool {
	numMap := make(map[uint8]bool)

	for _, num := range numbers {
		numMap[num] = true
	}

	for i := 0; i < 100; i++ {
		if _, ok := numMap[uint8(i)]; !ok {
			return false
		}
	}

	return true
}
