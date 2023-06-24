package rank_one

import (
	"hash/fnv"
	"sync"
)

const NumberOfShards = 10

type Shard struct {
	sync.RWMutex
	data map[string]string
}

type ShardedData struct {
	shards [NumberOfShards]*Shard
}

func NewShardedData() *ShardedData {
	d := &ShardedData{}
	for i := 0; i < NumberOfShards; i++ {
		d.shards[i] = &Shard{
			data: make(map[string]string),
		}
	}
	return d
}

func (d *ShardedData) GetShard(key string) *Shard {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	shardID := hasher.Sum32() % NumberOfShards
	return d.shards[shardID]
}

func (d *ShardedData) Get(key string) (value string, ok bool) {
	shard := d.GetShard(key)
	shard.RLock()
	defer shard.RUnlock()

	value, ok = shard.data[key]
	return
}

func (d *ShardedData) Set(key string, value string) {
	shard := d.GetShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.data[key] = value
}

func main() {
	data := NewShardedData()

	data.Set("key1", "value1")
	data.Set("key2", "value2")
	data.Set("key3", "value3")

	println(data.Get("key1")) // prints "value1"
	println(data.Get("key2")) // prints "value2"
	println(data.Get("key3")) // prints "value3"
}
