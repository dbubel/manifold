package topics

import "github.com/dbubel/manifold/buffer"

type Topic struct {
	Name  string
	Queue *buffer.DoublyLinkedd
}

func newTopic(name string) *Topic {
	return &Topic{
		Name:  name,
		Queue: buffer.NewBuffer(),
	}
}
