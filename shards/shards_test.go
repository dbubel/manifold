package shards

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	//for i:=0;i<10000;i++{
	//	pickShard(4)
	//}
	var m map[int64]int
	m = make(map[int64]int)
	i := 0
	timestamp := time.Now().UnixNano()
	rand.Seed(timestamp)

	for {
		i++
		if i > 10_000_000 {
			return
		}

		hashValue := rand.Int63()
		if _, ok := m[hashValue]; !ok {
			m[hashValue] = 1
			continue
		} else {
			t.Log("found collision", hashValue)
			t.Fail()
		}
	}
}

func TestRandomBuckets(t *testing.T) {

	var m map[int64]int
	m = make(map[int64]int)
	i := 0
	timestamp := time.Now().UnixNano()
	rand.Seed(timestamp)

	n := 1000000
	d := 4

	for {
		i++
		if i > n {
			break
		}

		hashValue := rand.Int63()
		hashValue = hashValue % int64(d)

		if v, ok := m[hashValue]; !ok {
			m[hashValue] = 1
			continue
		} else {
			m[hashValue] = v + 1
		}
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, 100*(float64(v)-float64(n/d))/float64(n/d))
	}
}

//
//func TestShardedTopics_Enqueue(t *testing.T) {
//	t.Run("test simple enqueue", func(t *testing.T) {
//		shards := NewShards(1)
//		shards.Enqueue("TopicName", []byte("hello shard"))
//	})
//}
//
//func TestShardedTopics_Dequeue(t *testing.T) {
//	t.Run("test simple enqueue", func(t *testing.T) {
//		shards := NewShards(1)
//		shards.Enqueue("TopicName", []byte("hello shard"))
//		shards.Dequeue("TopicName")
//	})
//}

//
//func TestShardedDataBasic(t *testing.T) {
//	data := NewShardedTopics(1)
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	defer cancel()
//
//	err := data.Enqueue("test", []byte("Hello World!"))
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//
//	dq := data.Dequeue(ctx, "test")
//
//	if string(dq) != "Hello World!" {
//		t.Errorf("Expected: %v, got: %v", "test", string(dq))
//	}
//}
//
//func TestShardedDataMultipleShards(t *testing.T) {
//	data := NewShardedTopics(10)
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//
//	err := data.Enqueue("test", []byte("Hello World!"))
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//	//time.Sleep(time.Millisecond * 100)
//
//	dq := data.Dequeue(ctx, "test")
//
//	if string(dq) != "Hello World!" {
//		t.Errorf("Expected: %v, got: %v", "test", string(dq))
//	}
//}
//
//func TestShardedDataBasicAsync(t *testing.T) {
//	data := NewShardedTopics(2)
//	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	//defer cancel()
//	//var m sync.RWMutex
//	for i := 0; i < 10; i++ {
//		go func(a int) {
//			data.Enqueue("test", []byte(fmt.Sprintf("Hello World! %d", a)))
//		}(i)
//	}
//	//time.Sleep(time.Millisecond * 100)
//	//results := make([][]byte, 10)
//	for i := 0; i < 10; i++ {
//		//ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
//		dq := data.Dequeue(context.Background(), "test")
//		_ = dq
//		fmt.Println(string(dq))
//		//cancel()
//		//if string(dq) != fmt.Sprintf("Hello World! %d", i) {
//		//	t.Errorf("Expected: %v, got: %v", "test", string(dq))
//		//}
//	}

//sort.Slice(results, func(i, j int) bool {
//	return string(results[i]) < string(results[j])
//})
//
//for i := 0; i < 10; i++ {
//	if string(results[i]) != fmt.Sprintf("Hello World! %d", i) {
//		t.Errorf("Expected: %v, got: %v", "test", string(results[i]))
//	}
//}

//}

//func TestShardedMultipleEnquqeDequeu(t *testing.T) {
//	shards := NewShardedTopics(2)
//	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
//	_ = ctx
//	defer cancel()
//
//	for i := 0; i < 10; i++ {
//		go func(a int) {
//			//time.Sleep(time.Millisecond * 100)
//			shards.Enqueue("test", []byte(fmt.Sprintf("Hello World! %d", a)))
//		}(i)
//	}
//	time.Sleep(time.Millisecond * 10000)
//
//	for i := 0; i < 10; i++ {
//
//		time.Sleep(time.Millisecond * 100)
//		dq, err := shards.Dequeue(ctx, "test")
//
//		if err != nil {
//			t.Errorf("Error: %v", err)
//		}
//
//		fmt.Println(string(dq))
//		//if string(dq) != fmt.Sprintf("Hello World! %d", a) {
//		//	t.Errorf("Expected: %v, got: %v", "test", string(dq))
//		//}
//
//	}
//}
