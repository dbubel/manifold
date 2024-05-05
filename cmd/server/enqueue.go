package server

import (
	"context"
	"fmt"

	"github.com/dbubel/manifold/internal"
	proto "github.com/dbubel/manifold/proto_files"
	"github.com/dbubel/manifold/queue"
)

func (s *server) Enqueue(_ context.Context, in *proto.EnqueueMsg) (*proto.Empty, error) {
	if in.GetTopicName() == internal.EmptyString {
		s.l.Error("topic name is required")
		return &proto.Empty{}, fmt.Errorf("topic name is required")
	}

	switch in.Priority {
	case proto.Priority_NORMAL:
		// s.topics.Enqueue(in.GetTopicName(), in.GetData())
		s.topics.Enqueue(in.GetTopicName(), &queue.Element{Value: in.GetData()})
	case proto.Priority_HIGH:
		s.topics.EnqueueHighPriority(in.GetTopicName(), &queue.Element{Value: in.GetData()})
	}

	s.l.WithFields(map[string]interface{}{"priority": in.Priority.String(), "topic": in.GetTopicName(), "dataLen": len(in.GetData())}).Debug("enqueue ok")

	return &proto.Empty{}, nil
}
