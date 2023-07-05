package server

import (
	"context"
	"fmt"
	proto "github.com/dbubel/manifold/proto_files"
)

func (s *server) Dequeue(ctx context.Context, in *proto.DequeueMsg) (*proto.DequeueAck, error) {
	if in.GetTopicName() == "" {
		s.l.Error("error dequeue empty topic")
		return &proto.DequeueAck{}, fmt.Errorf("topic name is required")
	}

	data := s.t.Dequeue(ctx, in.TopicName)

	s.l.WithFields(map[string]interface{}{"topic": in.GetTopicName()}).Debug("dequeued msg")

	return &proto.DequeueAck{Data: data}, nil
}

func (s *server) StreamDequeue(req *proto.DequeueMsg, stream proto.Manifold_StreamDequeueServer) error {
	// This is a loop that continues to stream messages to the client
	// TODO: add select for cancelling when the app is shut down

	for {
		select {
		case <-stream.Context().Done():
			fmt.Println("stream dequeue cancelled")
			return nil
		default:
			// TODO: return an object that we can finalize deque if the send fails
			data := s.t.Dequeue(stream.Context(), req.TopicName)

			// Create a new StreamResponse message
			res := &proto.DequeueAck{
				Data: data,
			}

			// Send the response to the client
			if err := stream.Send(res); err != nil {
				return err
			}
		}
	}
}
