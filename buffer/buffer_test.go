package buffer

import (
	"bytes"
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	val := cb.Read(ctx)
	if !bytes.Equal(val, []uint8{1}) {
		t.Errorf("Expected %d, but got %d", []uint8{1}, val)
	}

	val = cb.Read(ctx)
	if !bytes.Equal(val, []uint8{2}) {
		t.Errorf("Expected %d, but got %d", []uint8{2}, val)
	}

	val = cb.Read(ctx)
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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cb.Write([]byte{1})
	cb.Write([]byte{2})
	cb.Read(ctx)
	cb.Write([]byte{3})
	cb.Write([]byte{4})

	val := cb.Read(ctx)
	if !bytes.Equal(val, []byte{2}) {
		t.Errorf("Expected 2, but got %v", val)
	}

	val = cb.Read(ctx)
	if !bytes.Equal(val, []byte{3}) {
		t.Errorf("Expected 3, but got %v", val)
	}

	val = cb.Read(ctx)
	if !bytes.Equal(val, []byte{4}) {
		t.Errorf("Expected 4, but got %v", val)
	}
}

func TestCircularBuffer_ReadBeforeWrite(t *testing.T) {
	cb := NewBuffer()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cb.Write([]byte{1})
	}()

	val := cb.Read(ctx)
	if !bytes.Equal(val, []byte{1}) {
		t.Errorf("Expected 1, but got %v", val)
	}
}

//func TestCircularBuffer_Len(t *testing.T) {
//	cb := NewBuffer()
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	defer cancel()
//
//	cb.Write([]byte{1})
//	cb.Write([]byte{1})
//	cb.Write([]byte{1})
//	cb.Read(ctx)
//	t.Log(cb.Len(ctx))
//	t.Log(cb.Len(ctx))
//	t.Log(cb.Len(ctx))
//	t.Log(cb.Len(ctx))
//
//	cb.Read(ctx)
//	cb.Read(ctx)
//	cb.Read(ctx)
//	cb.Read(ctx)
//	t.Log(cb.Len(ctx))
//
//	//cb.Read(ctx)
//	//t.Log(cb.Length())
//	//cb.Read(ctx)
//	//t.Log(cb.Length())
//	//cb.Read(ctx)
//	//time.Sleep(time.Millisecond * 1000)
//
//}

//func TestCircularBuffer_LenEmpty(t *testing.T) {
//	cb := NewBuffer()
//	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
//	defer cancel()
//
//	//cb.Write([]byte{1})
//	//cb.Write([]byte{1})
//	//cb.Write([]byte{1})
//	//cb.Write([]byte{1})
//	//cb.Read(ctx)
//	t.Log(cb.Len(ctx))
//}

func TestCircularBuffer_ReadTimeout(t *testing.T) {
	cb := NewBuffer()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cb.Write([]byte{1})
	cb.Write([]byte{2})
	cb.Read(ctx)
	cb.Read(ctx)

	val := cb.Read(ctx)
	if val != nil {
		t.Errorf("Expected nil, but got %v", val)
	}
}

func BenchmarkQueue(b *testing.B) {
	q := NewBuffer()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
		q.Read(ctx)
	}
}

func TestCircularBuffer_ConcurrentReadWrite(t *testing.T) {
	cb := NewBuffer()
	t.Run("concurrent read write, write first", func(t *testing.T) {
		wg := sync.WaitGroup{}
		var mu sync.Mutex

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

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
				res = append(res, cb.Read(ctx)...)
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

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		var res []uint8
		for i := 0; i < 100; i++ {
			go func() {
				mu.Lock()
				defer mu.Unlock()
				res = append(res, cb.Read(ctx)...)
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
