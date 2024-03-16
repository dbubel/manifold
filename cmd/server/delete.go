package server

import (
	"context"
	"fmt"
	proto "github.com/dbubel/manifold/proto_files"
)

func (s *server) DeleteTopic(_ context.Context, in *proto.DeleteTopicMsg) (*proto.Empty, error) {
	if in.GetTopicName() == "" {
		s.l.Error("error delete empty topic")
		return &proto.Empty{}, fmt.Errorf("error topic name is required")
	}

	s.topics.DeleteTopic(in.GetTopicName())
	return &proto.Empty{}, nil
}
