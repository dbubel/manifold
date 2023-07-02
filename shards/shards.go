package shards

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/dbubel/manifold/topics"
	"hash/fnv"
	"time"
)

type Shard struct {
	topics *topics.Topics
	ID     string
}

type ShardedTopics struct {
	shards    map[uint32]*Shard
	NumShards uint32
}

func NewShardedTopics(shardNum uint32) *ShardedTopics {
	sd := &ShardedTopics{
		NumShards: shardNum,
		shards:    make(map[uint32]*Shard),
	}

	var i uint32
	for i = 0; i < shardNum; i++ {
		sd.shards[i] = &Shard{
			topics: topics.New(),
			ID:     fmt.Sprintf("shard-%d", i),
		}
	}
	return sd
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func (d *ShardedTopics) GetShard(key []byte) (*Shard, error) {
	hasher := fnv.New32a()
	_, err := hasher.Write(key)
	if err != nil {
		return nil, err
	}

	shardID := hasher.Sum32() % d.NumShards
	//fmt.Println("got shardID", shardID)
	return d.shards[shardID], nil
}

func (d *ShardedTopics) Dequeue(ctx context.Context, topic string) []uint8 {
	var dataChan = make(chan []uint8, d.NumShards)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	//fmt.Println("lenShards", len(d.shards))
	for i := range d.shards {
		go func(dChan chan []uint8, shard *Shard) {
			data := shard.topics.Dequeue(ctx, topic)
			if data != nil {
				//fmt.Println(shard.ID, "got data", string(data))
				dChan <- data
				//cancel()
			}
		}(dataChan, d.shards[i])
	}
	//time.Sleep(time.Millisecond * 100)

	return <-dataChan
}

func (d *ShardedTopics) Enqueue(id string, value []byte) error {
	//fmt.Println("enqueue", id, string(value))
	rnd, _ := generateRandomBytes(200)
	shard, err := d.GetShard(rnd)
	if err != nil {
		fmt.Println("error getting shard", err)
		return err
	}

	shard.topics.Enqueue(id, value)
	return nil
}
