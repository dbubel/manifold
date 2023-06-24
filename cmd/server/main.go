package server

import (
	"github.com/dbubel/manifold/logging"
	"github.com/dbubel/manifold/proto_files"
	"github.com/dbubel/manifold/queue"
	"google.golang.org/grpc"
	"net"
)

type ManifoldServerCmd struct {
}

func (c *ManifoldServerCmd) Help() string {
	return ""
}

func (c *ManifoldServerCmd) Synopsis() string {
	return ""
}

type server struct {
	proto.ManifoldServer
	q queue.Queues
	l *logging.Logger
}

func (c *ManifoldServerCmd) Run(args []string) int {
	grpcServer := grpc.NewServer()

	l := logging.New(logging.INFO)
	defer l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server stopped")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		l.Error(err.Error())
	}

	proto.RegisterManifoldServer(grpcServer, &server{
		q: make(queue.Queues),
		l: l,
	})

	l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server started")

	if err := grpcServer.Serve(lis); err != nil {
		l.Error(err.Error())
	}
	return 0
}
