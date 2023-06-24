package main

import (
	"context"
	"fmt"
	"github.com/dbubel/manifold"
	"github.com/dbubel/manifold/logging"
	sq "github.com/dbubel/manifold/proto_files"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	sq.ManifoldServer
	q manifold.Queues
	l *logging.Logger
}

func main() {
	fmt.Println("Starting server...")
	s := grpc.NewServer()

	l := logging.New(logging.INFO)
	defer l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server stopped")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		l.Error(err.Error())
	}

	serv := &server{
		q: make(manifold.Queues),
		l: l,
	}

	//go serv.sleepRandomTime()
	sq.RegisterManifoldServer(s, serv)

	l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server started")

	if err := s.Serve(lis); err != nil {
		l.Error(err.Error())
	}

}

func (s *server) ListTopics(_ context.Context, _ *sq.Empty) (*sq.StringList, error) {
	return &sq.StringList{MyMap: s.q.List()}, nil
}

func (s *server) Enqueue(_ context.Context, in *sq.EnqueueMsg) (*sq.EnqueueAck, error) {
	if in.GetTopicName() == "" {
		s.l.Error("error enqueue empty topic")
		return &sq.EnqueueAck{}, fmt.Errorf("topic name is required")
	}

	s.l.WithFields(map[string]interface{}{"topic": in.GetTopicName(), "dataLen": len(in.GetData())}).Debug("enqueue ok")

	s.q.Enqueue(in.GetTopicName(), in.GetData())
	return &sq.EnqueueAck{Data: "OK"}, nil
}

func (s *server) Dequeue(_ context.Context, in *sq.DequeueMsg) (*sq.DequeueAck, error) {
	if in.GetTopicName() == "" {
		s.l.Error("error dequeue empty topic")
		return &sq.DequeueAck{}, fmt.Errorf("topic name is required")
	}

	data, err := s.q.Dequeue(in.TopicName)
	if err != nil {
		return &sq.DequeueAck{}, err
	}

	s.l.WithFields(map[string]interface{}{"topic": in.GetTopicName()}).Debug("dequeue ok")

	return &sq.DequeueAck{Data: data}, nil
}

func (s *server) StreamDequeue(req *sq.DequeueMsg, stream sq.Manifold_StreamDequeueServer) error {
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
			res := &sq.DequeueAck{
				Data: result,
			}

			// Send the response to the client
			if err = stream.Send(res); err != nil {
				return err
			}
		}
	}
}
