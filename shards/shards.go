package shards

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/dbubel/manifold/topics"
	"hash/fnv"
)

type Shard struct {
	topics *topics.Topics
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

func (d *ShardedTopics) Dequeue(ctx context.Context, topic string) ([]uint8, error) {
	// we randomly generate a key to get a shard

	//shard, err := d.GetShard(x)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// from the randomly picked shard we attempt to dequeue an item from the topic
	//fmt.Println("dequeue", topic)
	//data := shard.topics.Dequeue(ctx, topic)
	//fmt.Println("dequeue res", topic, string(data))
	// if the topic is empty we iterate over shards and return the first item we find
	//if data == nil {

	// iterate over the shards via a for loop using an index
	for i := 0; i < int(d.NumShards); i++ {
		x, err := generateRandomBytes(10)
		if err != nil {
			return nil, err
		}

		shard, err := d.GetShard(x)
		if err != nil {
			return nil, err
		}
		tpc := shard.topics.GetTopic(topic)
		if tpc == nil {
			continue
		}

		if tpc.Queue.Len(context.Background()) > 0 {
			data := tpc.Queue.Read(context.Background())
			fmt.Println("topic exists and has items", topic, tpc.Queue.Len(context.Background()), string(data))
			return data, nil
		}
	}

	//for _, v := range d.shards {
	//	//ctx2, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//	//defer cancel()
	//	// we check if the topic exists on this shard
	//	tpc := v.topics.GetTopic(topic)
	//	if tpc == nil {
	//		continue
	//	}
	//	fmt.Println("topic exists", topic, tpc.Queue.Len(context.Background()))
	//
	//	// if the topic exists we attempt to dequeue an item
	//	if tpc.Queue.Len(context.Background()) > 0 {
	//		data := tpc.Queue.Read(context.Background())
	//		fmt.Println("topic exists and has items", topic, tpc.Queue.Len(context.Background()), string(data))
	//		return data, nil
	//
	//		//}
	//	}
	//}

	return nil, nil
}

func (d *ShardedTopics) Enqueue(id string, value []byte) error {
	rnd, _ := generateRandomBytes(20)
	shard, err := d.GetShard(rnd)
	if err != nil {
		return err
	}

	shard.topics.Enqueue(id, value)
	return nil
}
