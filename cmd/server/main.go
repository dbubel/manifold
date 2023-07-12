package server

import (
	"context"
	"github.com/dbubel/manifold/config"
	"github.com/dbubel/manifold/pkg/logging"
	"github.com/dbubel/manifold/proto_files"
	"github.com/dbubel/manifold/topics"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
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

	y := &server{
		t: topics.NewTopics(l),
		l: l,
	}
	proto.RegisterManifoldServer(grpcServer, y)

	l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server started")

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			l.Error(err.Error())
		}

	}()
	y.waitForShutdown(grpcServer)
	//time.Sleep(time.Second)
	return 0

}

func (s *server) waitForShutdown(server *grpc.Server) {
	// Create a channel to receive the interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = ctx

	server.GracefulStop()

	s.t.Shutdown()

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
