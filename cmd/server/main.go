package server

import (
	"fmt"
	"github.com/dbubel/manifold/logging"
	"github.com/dbubel/manifold/proto_files"
	"github.com/dbubel/manifold/topics"
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
	//q *shards.ShardedTopics
	l *logging.Logger
}

func (c *ManifoldServerCmd) Run(args []string) int {
	t := topics.New()
	go func() {
		time.Sleep(time.Second)
		for i := 0; i < 10; i++ {
			//fmt.Println("sending")
			t.Enqueue("topicOne", []byte(fmt.Sprintf("hello world %d", i)))
			//time.Sleep(time.Millisecond * 1)
		}

		//time.Sleep(time.Second)
		//for i := 0; i < 10; i++ {
		//	//fmt.Println("sending")
		//	t.Enqueue("topicOne", []byte(fmt.Sprintf("hello world %d", i)))
		//	//time.Sleep(time.Millisecond * 1)
		//}
	}()
	//time.Sleep(time.Second)

	for {
		fmt.Println(string(t.Dequeue("topicOne")))
	}
	//topics.Dequeue(topicOne)
	//grpcServer := grpc.NewServer()
	//
	//l := logging.New(logging.INFO)
	//defer l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server stopped")
	//
	//lis, err := net.Listen("tcp", ":50051")
	//if err != nil {
	//	l.Error(err.Error())
	//}
	//
	//proto.RegisterManifoldServer(grpcServer, &server{
	//	q: shards.NewShardedTopics(10),
	//	l: l,
	//})
	//
	//l.WithFields(map[string]interface{}{"port": ":50051"}).Info("server started")
	//
	//if err := grpcServer.Serve(lis); err != nil {
	//	l.Error(err.Error())
	//}
	return 0
}
