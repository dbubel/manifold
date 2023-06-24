package server

import (
	"context"
	"github.com/dbubel/manifold/proto_files"
)

func (s *server) ListTopics(_ context.Context, _ *proto.Empty) (*proto.StringList, error) {
	return &proto.StringList{MyMap: s.q.List()}, nil
}
