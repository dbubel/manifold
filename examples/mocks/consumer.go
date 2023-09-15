package mocks

import (
	"context"
	"fmt"
	"io"
	"log"

	echo "github.com/dbubel/manifold/proto_files"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConsumeCommand struct {
}

func (c *ConsumeCommand) Help() string {
	return ""
}

func (c *ConsumeCommand) Synopsis() string {
	return ""
}

func (c *ConsumeCommand) Run(args []string) int {

	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials())) //grpc.WithBlock()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	x := echo.NewManifoldClient(conn)

	streamer, err := x.StreamDequeue(context.Background(), &echo.DequeueMsg{TopicName: "hello23"})
	if err != nil {
		log.Fatalf("%v.MyStreamingMethod(_) = _, %v", c, err)
	}

	// Listen on the stream
	fmt.Println("Listening on the stream")
	for {
		response, err := streamer.Recv()
		if err == io.EOF {
			// If the stream is closed, exit the loop
			break
		}

		if err != nil {
			log.Fatalf("%v.MyStreamingMethod(_) = _, %v", c, err)
		}
		//fmt.Println(string(response.Data))
		_ = response
	}

	return 0
}
