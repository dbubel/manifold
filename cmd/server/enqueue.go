package server

import (
	"context"
	"fmt"
	"github.com/dbubel/manifold/proto_files"
)

func (s *server) Enqueue(_ context.Context, in *proto.EnqueueMsg) (*proto.EnqueueAck, error) {
	if in.GetTopicName() == "" {
		s.l.Error("error enqueue empty topic")
		return &proto.EnqueueAck{}, fmt.Errorf("topic name is required")
	}

	s.l.WithFields(map[string]interface{}{"topic": in.GetTopicName(), "dataLen": len(in.GetData())}).Debug("enqueue ok")

	s.q.Enqueue(in.GetTopicName(), in.GetData())
	return &proto.EnqueueAck{Data: "OK"}, nil
}
