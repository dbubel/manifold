package server

import (
	"context"
	"fmt"

	proto "github.com/dbubel/manifold/proto_files"
	"github.com/google/uuid"
)

func (s *server) StreamDequeue(req *proto.DequeueMsg, stream proto.Manifold_StreamDequeueServer) error {
	defer s.l.WithFields(map[string]interface{}{"topic": req.GetTopicName()}).Info("dequeue stream ending")

	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			if req.GetTopicName() == "" {
				return fmt.Errorf("error topic name is required")
			}

			data := s.topics.Dequeue(stream.Context(), req.TopicName)

			res := &proto.DequeueAck{
				Data: data.Value,
			}

			if err := stream.Send(res); err != nil {
				// if we had an error sending then enqueue the data we tried to send
				s.topics.Enqueue(req.TopicName, data)
				return err
			}
		}
	}
}

func (s *server) Ack(_ context.Context, in *proto.Remove) (*proto.Empty, error) {
	if in.GetTopicName() == "" {
		return &proto.Empty{}, fmt.Errorf("error topic name is required")
	}

	id, err := uuid.Parse(in.Id)
	if err != nil {
		return &proto.Empty{}, err
	}
	s.topics.AckDequeue(in.GetTopicName(), id)
	return &proto.Empty{}, nil
}
