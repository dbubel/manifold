package server

import (
	"context"
	"fmt"
	"github.com/dbubel/manifold/proto_files"
)

func (s *server) Dequeue(_ context.Context, in *proto.DequeueMsg) (*proto.DequeueAck, error) {
	if in.GetTopicName() == "" {
		s.l.Error("error dequeue empty topic")
		return &proto.DequeueAck{}, fmt.Errorf("topic name is required")
	}

	data, err := s.q.Dequeue(in.TopicName)
	if err != nil {
		return &proto.DequeueAck{}, err
	}

	s.l.WithFields(map[string]interface{}{"topic": in.GetTopicName()}).Debug("dequeue ok")

	return &proto.DequeueAck{Data: data}, nil
}

func (s *server) StreamDequeue(req *proto.DequeueMsg, stream proto.Manifold_StreamDequeueServer) error {
	// This is a loop that continues to stream messages to the client
	// TODO: add select for cancelling when the app is shut down
	ctx, cancel := context.WithCancel(stream.Context())
	defer cancel()
	defer fmt.Println("stream dequeue ended")

	for {
		select {
		case <-stream.Context().Done():
			fmt.Println("stream dequeue cancelled")
			return nil
		default:
			// TODO: return an object that we can finalize deque if the send fails
			result, err := s.q.BlockingDequeue(ctx, req.TopicName)
			if err != nil {
				s.l.WithFields(map[string]interface{}{
					"err": err.Error(),
				}).Error("error in blocking dequeue")
				return err
			}

			// Create a new StreamResponse message
			res := &proto.DequeueAck{
				Data: result,
			}

			// Send the response to the client
			if err = stream.Send(res); err != nil {
				return err
			}
		}
	}
}
