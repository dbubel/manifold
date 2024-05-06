package server

import (
	"context"
	"fmt"
	"time"

	"github.com/dbubel/manifold/internal"
	proto "github.com/dbubel/manifold/proto_files"
	"github.com/dbubel/manifold/queue"
	"github.com/google/uuid"
)

func (s *server) Enqueue(_ context.Context, in *proto.EnqueueMsg) (*proto.Empty, error) {
	if in.GetTopicName() == internal.EmptyString {
		s.l.Error("topic name is required")
		return &proto.Empty{}, fmt.Errorf("topic name is required")
	}

	element := queue.Element{
		Value:       in.GetData(),
		ID:          uuid.New(),
		EnqueueTime: time.Now(),
		Complete:    make(chan struct{}),
	}

	switch in.Priority {
	case proto.Priority_NORMAL:
		s.topics.Enqueue(in.GetTopicName(), &element)
	case proto.Priority_HIGH:
		s.topics.EnqueueHighPriority(in.GetTopicName(), &element)
	}

	s.l.WithFields(map[string]interface{}{
		"priority": in.Priority.String(),
		"topic":    in.GetTopicName(),
		"dataLen":  len(in.GetData()),
	}).Debug("enqueue ok")

	return &proto.Empty{}, nil
}
