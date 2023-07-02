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
	fmt.Println("got shardID", shardID)
	return d.shards[shardID], nil
}

func (d *ShardedTopics) Dequeue(ctx context.Context, topic string) []uint8 {
	var dataChan = make(chan []uint8, d.NumShards)
	//var wg sync.WaitGroup
	//wg.Add(int(d.NumShards))

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	for i := range d.shards {

		go func(dChan chan []uint8, shard *Shard) {
			//defer wg.Done()
			//rnd, _ := generateRandomBytes(10)
			//data, := d.shards[0].topics.Dequeue(ctx, topic)

			//if err != nil {
			//	return
			//}

			//ctxx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
			//data := shard.topics.Dequeue(ctx, topic)
			//if data != nil {
			//	cancel()
			//	dChan <- data
			//}
			fmt.Println(shard.ID)
			data := shard.topics.Dequeue(ctx, topic)
			if data == nil {
				cancel()
			}
			dChan <- data
			//cancel()

		}(dataChan, d.shards[i])
	}
	//wg.Wait()
	//close(dataChan)
	//fmt.Println("From ")
	//for i := range dataChan {
	//	if i != nil {
	//		//cancel()
	//		return i, nil
	//	}
	//}
	return <-dataChan

	//wg.Wait()
	//fmt.Println("From chans", <-dataChan)

	//for i := range dataChan {
	//	if i != nil {
	//		return i, nil
	//	}
	//}

	// we randomly generate a key to get a shard
	//x, err := generateRandomBytes(10)
	//shard, err := d.GetShard(x)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// from the randomly picked shard we attempt to dequeue an item from the topic
	//data := shard.topics.Dequeue(ctx, topic)
	//if data != nil {
	//	return data, nil
	//}
	//
	//// iterate over the shards via a for loop using an index
	////for i := 0; i < int(d.NumShards); i++ {
	////	x, err := generateRandomBytes(10)
	////	if err != nil {
	////		return nil, err
	////	}
	////
	////	shard, err := d.GetShard(x)
	////	if err != nil {
	////		return nil, err
	////	}
	////	tpc := shard.topics.GetTopic(topic)
	////	if tpc == nil {
	////		continue
	////	}
	////
	////	if tpc.Queue.Len(context.Background()) > 0 {
	////		data := tpc.Queue.Read(context.Background())
	////		fmt.Println("topic exists and has items", topic, tpc.Queue.Len(context.Background()), string(data))
	////		return data, nil
	////	}
	////}
	//
	//for _, v := range d.shards {
	//	ctx2, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//	defer cancel()
	//	// we check if the topic exists on this shard
	//	tpc := v.topics.GetTopic(topic)
	//	if tpc == nil {
	//		continue
	//	}
	//	//fmt.Println("topic exists", topic, tpc.Queue.Len(context.Background()))
	//
	//	// if the topic exists we attempt to dequeue an item
	//	if tpc.Queue.Len(ctx2) > 0 {
	//		data = tpc.Queue.Read(ctx2)
	//		//fmt.Println("topic exists and has items", topic, tpc.Queue.Len(context.Background()), string(data))
	//		return data, nil
	//
	//		//}
	//	}
	//}
	//
	//return nil, nil
}

func (d *ShardedTopics) Enqueue(id string, value []byte) error {
	fmt.Println("enqueue", id, string(value))
	rnd, _ := generateRandomBytes(200)
	shard, err := d.GetShard(rnd)
	if err != nil {
		fmt.Println("error getting shard", err)
		return err
	}

	shard.topics.Enqueue(id, value)
	return nil
}
