package queue

import (
	"time"

	"github.com/google/uuid"
)

func (q *Queue) RemoveDequed(e uuid.UUID) {
	if ele, exists := q.dequeued[e]; exists {
		ele.Complete <- struct{}{}
		delete(q.dequeued, e)
	}
}

func (q *Queue) AddDequeued(e *Element) {
	go func(e *Element) {
		TimeTicker := time.NewTicker(time.Second * 30)
		defer TimeTicker.Stop()
		select {
		case <-e.Complete:
			q.log.WithFields(map[string]interface{}{
				"id": e.ID.String(),
			}).Debug("messaged ack removing")
			return
		case <-TimeTicker.C:
			q.log.WithFields(map[string]interface{}{
				"id": e.ID,
			}).Debug("back on to the queue")
			q.Enqueue(e)
			return
		}
	}(e)
}
