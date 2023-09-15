package server

import (
	"context"
	"fmt"
	"github.com/dbubel/manifold/internal"
	proto "github.com/dbubel/manifold/proto_files"
	"math/rand"
	"runtime"
	"time"
)

func (s *server) Enqueue(_ context.Context, in *proto.EnqueueMsg) (*proto.EnqueueAck, error) {
	if in.GetTopicName() == internal.EmptyString {
		s.l.Error("topic name is required")
		return &proto.EnqueueAck{}, fmt.Errorf("topic name is required")
	}

	switch in.Priority {
	case proto.Priority_NORMAL:
		s.topics[pickShard(runtime.NumCPU())].Enqueue(in.GetTopicName(), in.GetData())
	case proto.Priority_HIGH:
		s.topics.EnqueueHighPriority(in.GetTopicName(), in.GetData())
	}

	s.l.WithFields(map[string]interface{}{"priority": in.Priority.String(), "topic": in.GetTopicName(), "dataLen": len(in.GetData())}).Debug("enqueue ok")

	return &proto.EnqueueAck{Data: "OK"}, nil
}

func pickShard(numShards int) int {
	// Get the current timestamp
	timestamp := time.Now().UnixNano()

	// Seed the random number generator with the timestamp
	rand.Seed(timestamp)

	// Generate a random hash value using the seeded random number generator
	hashValue := rand.Int63()

	// Pick a shard based on the hash value
	shardID := int(hashValue % int64(numShards))
	return shardID
}
