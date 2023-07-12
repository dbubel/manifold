package store

import (
	"context"
	"github.com/dbubel/manifold/internal"
)

// noopStore is the default implementation of store if none is provided
// on creation. Just provides a pass through if you dont wish to have
// persistent storage of unprocessed messages
type noopStore struct{}

func (n noopStore) HealthCheck(ctx context.Context, fn context.CancelFunc) error {
	return nil
}
func (s noopStore) Insert(shards []string) (string, error) {
	return internal.EmptyString, nil
}
