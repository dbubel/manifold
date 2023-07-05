//go:build exclude
// +build exclude

package client

import (
	"context"
	"github.com/dbubel/manifold/proto_files"
	"testing"
)

func TestNewManifoldClient_Enqueue(t *testing.T) {
	client, err := NewManifoldClient("localhost:50051")
	if err != nil {
		t.Error("error making client", err.Error())
		return
	}

	ack, err := client.Enqueue(context.Background(), &proto.EnqueueMsg{
		TopicName: "test_topic",
		Data:      []byte("hello world"),
	})

	if err != nil {
		t.Error("error enqueueing", err.Error())
	}

	if ack.GetData() != "OK" {
		t.Errorf("expected 'OK' got [%s]", err.Error())
	}
}

func TestNewManifoldClient_EnqueueDequeue(t *testing.T) {
	client, err := NewManifoldClient("")
	if err != nil {
		t.Error("error making client", err.Error())
		return
	}

	ack, err := client.Enqueue(context.Background(), &proto.EnqueueMsg{
		TopicName: "test_topic",
		Data:      []byte("hello world"),
	})

	if err != nil {
		t.Error("error enqueueing", err.Error())
	}

	if ack.GetData() != "OK" {
		t.Errorf("expected 'OK' got [%s]", err.Error())
	}

	dack, err := client.Dequeue(context.Background(), &proto.DequeueMsg{
		TopicName: "test_topic",
	})

	if err != nil {
		t.Error("error dequeueing", err.Error())
	}

	if string(dack.GetData()) != "hello world" {
		t.Errorf("expected 'hello world' got [%s]", string(dack.GetData()))
	}
}
