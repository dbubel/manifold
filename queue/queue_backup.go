package queue

import "time"

func (q *Queue) RemoveDequed(e *Element) {
	e.Complete <- struct{}{}
	delete(q.dequeued, e.ID)
}

func (q *Queue) AddDequeued(e *Element) {
	q.log.WithFields(map[string]interface{}{
		"elementID": e.ID.String(),
	}).Debug("adding to backup")

	go func(e *Element) {
		TimeTicker := time.NewTicker(time.Second * 3)
		defer TimeTicker.Stop()
		select {
		case <-e.Complete:
			q.log.Debug("ackknowledged remove")
			return
		case <-TimeTicker.C:
			q.log.WithFields(map[string]interface{}{
				"uid": e.ID,
			}).Debug("ticked, back on to the queue")

			q.Enqueue(e)
			return
		}
	}(e)
}
