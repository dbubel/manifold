package server

import (
	"context"
	"fmt"
	proto "github.com/dbubel/manifold/proto_files"
)

func (s *server) Dequeue(ctx context.Context, in *proto.DequeueMsg) (*proto.DequeueAck, error) {
	if in.GetTopicName() == "" {
		s.l.Error("error dequeue empty topic")
		return &proto.DequeueAck{}, fmt.Errorf("error topic name is required")
	}

	data := s.topics.Dequeue(ctx, in.TopicName)

	s.l.WithFields(map[string]interface{}{"topic": in.GetTopicName()}).Debug("dequeue msg")

	return &proto.DequeueAck{Data: data}, nil
}

func (s *server) StreamDequeue(req *proto.DequeueMsg, stream proto.Manifold_StreamDequeueServer) error {
	defer s.l.WithFields(map[string]interface{}{"topic": req.GetTopicName()}).Info("dequeue stream ending")
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			data := s.topics.Dequeue(stream.Context(), req.TopicName)

			res := &proto.DequeueAck{
				Data: data,
			}

			if err := stream.Send(res); err != nil {
				return err
			}
		}
	}
}
