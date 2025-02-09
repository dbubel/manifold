package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dbubel/manifold/config"
	"github.com/dbubel/manifold/pkg/logging"
	proto "github.com/dbubel/manifold/proto_files"
	"github.com/dbubel/manifold/topics"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type ManifoldServerCmd struct{}

func (c *ManifoldServerCmd) Help() string {
	return ""
}

func (c *ManifoldServerCmd) Synopsis() string {
	return ""
}

type server struct {
	proto.ManifoldServer
	// topics *shards.TopicShards
	topics *topics.Topics
	l      *logging.Logger
}

func (c *ManifoldServerCmd) Run(args []string) int {
	l := logging.New(logging.DEBUG)

	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		l.Error(err.Error())
		return 0
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(mwLogger(l)), grpc.ConnectionTimeout(time.Second))
	defer l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server stopped")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		l.Error(err.Error())
		return 0
	}
	//
	// y := &server{
	// 	topics: shards.NewShards(runtime.NumCPU(), l),
	// 	l:      l,
	// }

	y := &server{
		topics: topics.NewTopics(l),
		l:      l,
	}

	proto.RegisterManifoldServer(grpcServer, y)

	l.WithFields(map[string]interface{}{"port": ":50052"}).Info("server started")

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			l.Error(err.Error())
		}
	}()

	y.waitForShutdown(grpcServer)

	return 0
}

func (s *server) waitForShutdown(server *grpc.Server) {
	// Create a channel to receive the interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	// Block until a signal is received
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = ctx
	s.l.Info("server shutting down...")
	server.Stop()
	// server.GracefulStop()

	// s.topics.ShutdownTopics()
}

// UnaryInterceptor is a gRPC middleware that logs the duration of each unary RPC call.
func mwLogger(_ *logging.Logger) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// start := time.Now()

		// Call the handler to process the RPC request
		resp, err := handler(ctx, req)

		// duration := time.Since(start)
		// l.WithFields(map[string]interface{}{"method": info.FullMethod, "duration": duration.String()}).Debug("handled request")

		return resp, err
	}
}
