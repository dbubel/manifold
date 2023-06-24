package mocks

import (
	"context"
	"fmt"
	"github.com/dbubel/manifold/internal"
	"github.com/dbubel/manifold/logging"
	"github.com/dbubel/manifold/proto_files"
	"github.com/dbubel/manifold/queue"
	"google.golang.org/grpc"
	"net"
)

type ServeCommand struct {
}

func (c *ServeCommand) Help() string {
	return ""
}

func (c *ServeCommand) Synopsis() string {
	return "Runs the cohesion content API server"
}

func (c *ServeCommand) Run(args []string) int {
	fmt.Println("Starting server...")
	s := grpc.NewServer()

	l := logging.New(logging.INFO)
	defer l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server stopped")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		l.Error(err.Error())
	}

	serv := &server{
		q: make(queue.Queues),
		l: l,
	}

	proto.RegisterManifoldServer(s, serv)

	l.WithFields(logging.Fields{"port": ":50051"}).Info("server started")

	if err := s.Serve(lis); err != nil {
		l.Error(err.Error())
	}
	return 0
}

type server struct {
	proto.ManifoldServer
	q queue.Queues
	l *logging.Logger
}

func (s *server) ListTopics(_ context.Context, _ *proto.Empty) (*proto.StringList, error) {
	return &proto.StringList{MyMap: s.q.List()}, nil
}

func (s *server) Enqueue(_ context.Context, in *proto.EnqueueMsg) (*proto.EnqueueAck, error) {
	if in.GetTopicName() == internal.EmptyString {
		s.l.Error("error enqueue empty topic")
		return &proto.EnqueueAck{}, fmt.Errorf("topic name is required")
	}

	s.l.WithFields(logging.Fields{"topic": in.GetTopicName(), "dataLen": len(in.GetData())}).Debug("enqueue ok")

	s.q.Enqueue(in.GetTopicName(), in.GetData())
	return &proto.EnqueueAck{Data: "OK"}, nil
}

func (s *server) Dequeue(_ context.Context, in *proto.DequeueMsg) (*proto.DequeueAck, error) {
	if in.GetTopicName() == internal.EmptyString {
		s.l.Error("error dequeue empty topic")
		return &proto.DequeueAck{}, fmt.Errorf("topic name is required")
	}

	data, err := s.q.Dequeue(in.TopicName)
	if err != nil {
		return &proto.DequeueAck{}, err
	}

	s.l.WithFields(logging.Fields{"topic": in.GetTopicName()}).Debug("dequeue ok")

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
				s.l.WithFields(logging.Fields{
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
