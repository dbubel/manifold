package buffer2

import (
	"context"
	"testing"
	"time"
)

func tobyte(s string) []uint8 {
	return []byte(s)
}
func TestQueue_Enqueue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	q := NewQueue()
	q.Enqueue(tobyte("hello"))
	val := q.Dequeue(ctx)
	t.Log(string(val.Value))

	val = q.Dequeue(ctx)
	t.Log(string(val.Value))

	//val = q.Dequeue()
	//t.Log(val.(*list.Element))
}
