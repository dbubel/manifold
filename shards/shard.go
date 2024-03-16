package shards

// import (
// 	"context"
// 	"github.com/dbubel/manifold/pkg/logging"
// 	"github.com/dbubel/manifold/topics"
// 	"math/rand"
// 	"time"
// )

//var shardCount = runtime.NumCPU()
//
// type TopicShards struct {
// 	topicShards []*topics.Topics
// }
//
// func pickShard(numShards int) int {
// 	timestamp := time.Now().UnixNano()
// 	rand.Seed(timestamp)
// 	hashValue := rand.Int63()
// 	shardID := int(hashValue % int64(numShards))
// 	return shardID
// }
//
// func NewShards(n int, l *logging.Logger) *TopicShards {
// 	var topicShards []*topics.Topics
// 	for i := 0; i < n; i++ {
// 		topicShards = append(topicShards, topics.NewTopics(l))
// 	}
//
// 	return &TopicShards{topicShards: topicShards}
// }
//
// func (t *TopicShards) Enqueue(topicName string, data []byte) {
// 	shardId := pickShard(len(t.topicShards))
// 	queue := t.topicShards[shardId].GetOrCreateTopic(topicName)
// 	queue.Enqueue(data)
// }
//
// func (t *TopicShards) EnqueueHighPriority(topicName string, data []byte) {
// 	shardId := pickShard(len(t.topicShards))
// 	queue := t.topicShards[shardId].GetOrCreateTopic(topicName)
// 	queue.EnqueueHighPriority(data)
// }
//
// func (t *TopicShards) Len(topicName string) int {
// 	totals := 0
// 	for i := 0; i < len(t.topicShards); i++ {
// 		//fmt.Println("shard", i, "len", t.topicShards[i].Len(topicName))
// 		totals = totals + t.topicShards[i].Len(topicName)
// 	}
//
// 	return totals
// }

// func (t *TopicShards) Dequeue(ctx context.Context, topicName string) []byte {
// 	shardId := pickShard(len(t.topicShards))
// 	queue := t.topicShards[shardId].GetOrCreateTopic(topicName)
// 	return queue.BlockingDequeue(ctx)
// }
//
// func (t *TopicShards) DeleteTopic(topicName string) {
// 	for i := 0; i < len(t.topicShards); i++ {
// 		t.topicShards[i].DeleteTopic(topicName)
// 	}
// }
//
//
//import (
//	"fmt"
//	"github.com/dbubel/manifold/topics"
//	"math/rand"
//	"time"
//)
//
//type Shard struct {
//	topics *topics.Topics
//	ID     string
//}
//
//func NewShard() *Shard {
//	return &Shard{
//		topics: topics.New(),
//		ID:     fmt.Sprintf("shard-%d", randomNDigitString(3)),
//	}
//}
//
//func (s *Shard) Enqueue(topicName string, data []byte) {
//	s.topics.Enqueue(topicName, data)
//}
//
//// Function to generate a random n-digit string
//func randomNDigitString(n int) string {
//	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time
//
//	// Define the range of characters allowed in the string
//	charset := "0123456789"
//
//	// Generate the string by randomly selecting characters from the charset
//	result := make([]byte, n)
//	for i := range result {
//		result[i] = charset[rand.Intn(len(charset))]
//	}
//
//	return string(result)
//}
