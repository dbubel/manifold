package buffer

import (
	"bytes"
	"math/rand"
	"sort"
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

func TestCircularBuffer_Len(t *testing.T) {
	cb := NewBuffer()

	cb.Write([]byte{1})
	cb.Write([]byte{1})
	cb.Write([]byte{1})
	cb.Read()
	t.Log(cb.Length())
	t.Log(cb.Length())
	t.Log(cb.Length())
	t.Log(cb.Length())

	cb.Read()
	cb.Read()
	cb.Read()
	cb.Read()
	t.Log(cb.Length())

	//cb.Read()
	//t.Log(cb.Length())
	//cb.Read()
	//t.Log(cb.Length())
	//cb.Read()
	//time.Sleep(time.Millisecond * 1000)

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
	t.Run("concurrent read write, write first", func(t *testing.T) {
		wg := sync.WaitGroup{}
		var mu sync.Mutex

		for i := 0; i < 100; i++ {
			go func(a int) {
				wg.Add(1)
				cb.Write([]uint8{uint8(a)})
			}(i)
		}

		var res []uint8
		for i := 0; i < 100; i++ {
			go func() {
				mu.Lock()
				defer mu.Unlock()
				res = append(res, cb.Read()...)
				wg.Done()
			}()
		}

		wg.Wait()

		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})

		// verify that the list res contains all the numbers from 0 to 99
		for i := 0; i < 100; i++ {
			if res[i] != uint8(i) {
				t.Errorf("Expected %d, but got %d", i, res[i])
			}
		}
	})

	t.Run("concurrent read write, read first", func(t *testing.T) {
		wg := sync.WaitGroup{}
		var mu sync.Mutex

		var res []uint8
		for i := 0; i < 100; i++ {
			go func() {
				mu.Lock()
				defer mu.Unlock()
				res = append(res, cb.Read()...)
				wg.Done()
			}()
		}

		for i := 0; i < 100; i++ {
			go func(a int) {
				wg.Add(1)
				cb.Write([]uint8{uint8(a)})
			}(i)
		}

		wg.Wait()

		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})

		// verify that the list res contains all the numbers from 0 to 99
		for i := 0; i < 100; i++ {
			if res[i] != uint8(i) {
				t.Errorf("Expected %d, but got %d", i, res[i])
			}
		}
	})
}
