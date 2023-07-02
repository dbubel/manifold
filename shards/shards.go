package shards

import (
	"crypto/rand"
	"fmt"
	"github.com/dbubel/manifold/topics"
	"hash/fnv"
)

type ShardedTopics struct {
	shards    map[uint32]*topics.Topics
	input     chan Wrapper
	outputReq chan string
	output    chan []byte
	NumShards uint32
}

func NewShards(shardNum uint32) *ShardedTopics {
	sd := &ShardedTopics{
		NumShards: shardNum,
		input:     make(chan Wrapper),
		output:    make(chan []byte),
		outputReq: make(chan string),
		shards:    make(map[uint32]*topics.Topics),
	}

	var i uint32
	for i = 0; i < shardNum; i++ {
		sd.shards[i] = topics.New()
	}
	go sd.run()
	return sd
}

type Wrapper struct {
	TopicName string
	Data      []byte
}

func (t *ShardedTopics) run() {
	for {
		select {
		case val := <-t.input:
			b, _ := generateRandomBytes(20)
			n, _ := t.GetShardNum(b)
			fmt.Println("shard num in", n)
			t.shards[n].Enqueue(val.TopicName, val.Data)
		case topicName := <-t.outputReq:
			b, _ := generateRandomBytes(20)
			n, _ := t.GetShardNum(b)
			fmt.Println("shard num out", n)
			a := t.shards[n].Dequeue(topicName)
			fmt.Println("AAA", a)
			t.output <- a
		}
	}
}
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func (d *ShardedTopics) GetShardNum(key []byte) (uint32, error) {
	var shardNum uint32
	hasher := fnv.New32a()
	_, err := hasher.Write(key)
	if err != nil {
		return 0, err
	}

	shardNum = hasher.Sum32() % d.NumShards
	return shardNum, nil
}

func (s *ShardedTopics) Enqueue(topicName string, data []byte) {
	s.input <- Wrapper{
		TopicName: topicName,
		Data:      data,
	}
}

func (s *ShardedTopics) Dequeue(topicName string) []byte {
	s.outputReq <- topicName
	fmt.Println("here")
	return <-s.output
}

//
//func (d *ShardedTopics) Dequeue(ctx context.Context, topic string) []uint8 {
//	var dataChan = make(chan []uint8, d.NumShards)
//
//	ctx, cancel := context.WithTimeout(ctx, time.Second)
//	defer cancel()
//	//fmt.Println("lenShards", len(d.shards))
//	for i := range d.shards {
//		go func(dChan chan []uint8, shard *Shard) {
//			data := shard.topics.Dequeue(ctx, topic)
//			if data != nil {
//				//fmt.Println(shard.ID, "got data", string(data))
//				dChan <- data
//				//cancel()
//			}
//		}(dataChan, d.shards[i])
//	}
//	//time.Sleep(time.Millisecond * 100)
//
//	return <-dataChan
//}
//
//func (d *ShardedTopics) Enqueue(id string, value []byte) error {
//	//fmt.Println("enqueue", id, string(value))
//	rnd, _ := generateRandomBytes(200)
//	shard, err := d.GetShard(rnd)
//	if err != nil {
//		fmt.Println("error getting shard", err)
//		return err
//	}
//
//	shard.topics.Enqueue(id, value)
//	return nil
//}
