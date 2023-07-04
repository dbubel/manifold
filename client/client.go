package client

import (
	"github.com/dbubel/manifold/proto_files"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MClient struct {
	proto.ManifoldClient
}

func NewManifoldClient(host string) (*MClient, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) //grpc.WithBlock()
	if err != nil {
		return nil, err
	}

	x := MClient{proto.NewManifoldClient(conn)}

	return &x, nil
}
