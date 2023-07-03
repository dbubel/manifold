package scrap

import (
	"context"
	"testing"
	"time"
)

func TestLinkedListSignal_Front(t *testing.T) {
	ll := NewLinkedList()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	node := ll.Front(ctx)
	t.Log(node)
}

func TestLinkedListBlocking(t *testing.T) {
	ll := NewLinkedList()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	var node *Node
	go func() {
		node = ll.Front(ctx)
	}()

	time.Sleep(time.Millisecond * 500) // Delay before push
	ll.PushBack(10)

	time.Sleep(time.Millisecond * 500) // Give time for operation to complete

	if node == nil {
		t.Error("Expected a node, got nil")
	} else if node.data != 10 {
		t.Error("Expected node value to be 10, got ", node.data)
	}
}

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()

	ll.PushBack(10)
	ll.PushBack(20)
	ll.PushBack(30)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	node := ll.Front(ctx)
	if node == nil || node.data != 10 {
		t.Error("Expected head node value to be 10, got ", node.data)
	}

	ll.Remove(20)

	// Iterate to find the removed element
	for n := ll.Front(ctx); n != nil; n = n.next {
		if n.data == 20 {
			t.Error("Expected element with value 20 to be removed")
		}
	}

	// Check if front returns nil after the context is cancelled
	ctx, cancel = context.WithCancel(context.Background())
	cancel()
	node = ll.Front(ctx)
	if node != nil {
		t.Error("Expected front method to return nil after context cancelation")
	}
}
