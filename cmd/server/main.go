package server

import (
	"context"
	"github.com/dbubel/manifold/config"
	"github.com/dbubel/manifold/logging"
	"github.com/dbubel/manifold/proto_files"
	"github.com/dbubel/manifold/topics"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"net"
	"time"
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
	t *topics.Topics
	l *logging.Logger
}

func (c *ManifoldServerCmd) Run(args []string) int {
	l := logging.New(logging.DEBUG)

	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		l.Error(err.Error())
		return 0
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(xxx(l)))
	defer l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server stopped")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		l.Error(err.Error())
	}

	proto.RegisterManifoldServer(grpcServer, &server{
		t: topics.NewTopics(),
		l: l,
	})

	l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server started")

	if err := grpcServer.Serve(lis); err != nil {
		l.Error(err.Error())
	}
	return 0
}

// UnaryInterceptor is a gRPC middleware that logs the duration of each unary RPC call.
func xxx(l *logging.Logger) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		// Call the handler to process the RPC request
		resp, err := handler(ctx, req)

		duration := time.Since(start)
		l.WithFields(map[string]interface{}{"method": info.FullMethod, "duration": duration.String()}).Debug("handled request")

		return resp, err
	}
}

// z := func a(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//	start := time.Now()
//
//	// Call the handler to process the RPC request
//	resp, err := handler(ctx, req)
//
//	duration := time.Since(start)
//	log.Printf("Unary RPC call %s took %s", info.FullMethod, duration)
//
//	return resp, err
//}
