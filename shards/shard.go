package shards

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
