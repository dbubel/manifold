package shards

import (
	"context"
	"crypto/rand"
	"github.com/dbubel/manifold/topics"
	"hash/fnv"
)

type Shard struct {
	queues topics.Topics
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
			queues: make(topics.Topics),
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
	return d.shards[shardID], nil
}

// Experimental
func (d *ShardedTopics) BlockingDequeue(ctx context.Context, topic string) ([]uint8, error) {
	x, err := generateRandomBytes(10)
	if err != nil {
		return nil, err
	}

	shard, err := d.GetShard(x)
	if err != nil {
		return nil, err
	}

	data, err := shard.queues.Dequeue(topic)

	if err != nil && err.Error() == "queue is empty" {
		for _, v := range d.shards {
			if i, _ := v.queues.Len(topic); i > 0 {
				data, err = v.queues.Dequeue(topic)
				break
			}
		}
	}

	data, err = shard.queues.BlockingDequeue(ctx, topic)

	return data, err
}

func (d *ShardedTopics) Dequeue(topic string) ([]uint8, error) {
	x, err := generateRandomBytes(10)
	if err != nil {
		return nil, err
	}

	shard, err := d.GetShard(x)
	if err != nil {
		return nil, err
	}

	data, err := shard.queues.Dequeue(topic)

	if err != nil && err.Error() == "queue is empty" {
		for _, v := range d.shards {
			if i, _ := v.queues.Len(topic); i > 0 {
				data, err = v.queues.Dequeue(topic)
				break
			}
		}
	}
	return data, err
}

func (d *ShardedTopics) Enqueue(id string, value []byte) error {
	rnd, _ := generateRandomBytes(20)
	shard, err := d.GetShard(rnd)
	if err != nil {
		return err
	}

	z := shard.queues.Enqueue(id, value)
	return z
}
