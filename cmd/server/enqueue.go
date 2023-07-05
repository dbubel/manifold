package server

import (
	"context"
	"fmt"
	"github.com/dbubel/manifold/internal"
	proto "github.com/dbubel/manifold/proto_files"
)

func (s *server) Enqueue(_ context.Context, in *proto.EnqueueMsg) (*proto.EnqueueAck, error) {
	if in.GetTopicName() == internal.EmptyString {
		s.l.Error("topic name is required")
		return &proto.EnqueueAck{}, fmt.Errorf("topic name is required")
	}

	s.l.WithFields(map[string]interface{}{"topic": in.GetTopicName(), "dataLen": len(in.GetData())}).Debug("enqueue ok")

	s.t.Enqueue(in.GetTopicName(), in.GetData())

	return &proto.EnqueueAck{Data: "OK"}, nil
}
