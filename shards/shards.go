package shards

import (
	"crypto/rand"
	"fmt"
	"github.com/dbubel/manifold/queue"
	"hash/fnv"
)

//const NumberOfShards = 1

type Shard struct {
	//sync.RWMutex
	queues queue.Queues
	//NumShards int
}

type ShardedData struct {
	shards    []*Shard
	NumShards int
}

func NewShardedQueues(shardNum int) *ShardedData {
	d := &ShardedData{
		NumShards: shardNum,
	}
	for i := 0; i < shardNum; i++ {
		d.shards = append(d.shards, &Shard{
			queues: make(queue.Queues),
		})
	}
	return d
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (d *ShardedData) GetShard(key []byte) (*Shard, error) {
	hasher := fnv.New32a()
	_, err := hasher.Write(key)
	if err != nil {
		return nil, err
	}
	shardID := hasher.Sum32() % uint32(d.NumShards)
	return d.shards[shardID], nil
}

func (d *ShardedData) GetShardList() {

}

func (d *ShardedData) Dequeue(topic string) ([]byte, error) {
	x, _ := generateRandomBytes(10)
	//fmt.Printf("%x", x)
	shard, _ := d.GetShard(x)

	bytes, err := shard.queues.Dequeue(topic)

	if err != nil && err.Error() == "queue is empty" {
		for _, s := range d.shards {
			fmt.Println(s.queues.Len(topic))
			if i, _ := s.queues.Len(topic); i > 0 {
				bytes, err = s.queues.Dequeue(topic)
				break
			}
		}
	}
	return bytes, err
}

func (d *ShardedData) Enqueue(id string, value []byte) error {
	rnd, _ := generateRandomBytes(20)
	shard, err := d.GetShard(rnd)
	if err != nil {
		return err
	}

	//shard.Lock()
	//defer shard.Unlock()
	z := shard.queues.Enqueue(id, value)
	return z
	//shard.queues[key] = value
}

//
//func main() {
//	queues := NewShardedQueues()
//
//	queues.Set("key1", "value1")
//	queues.Set("key2", "value2")
//	queues.Set("key3", "value3")
//
//	println(queues.Get("key1")) // prints "value1"
//	println(queues.Get("key2")) // prints "value2"
//	println(queues.Get("key3")) // prints "value3"
//}
