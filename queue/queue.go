package queue

//
//type Queue struct {
//	list        *linked_list.List
//	lock        sync.Mutex
//	Cond        *sync.Cond
//	compression *compression.SnappyCompressor
//}
//
//func NewQueue() *Queue {
//	q := &Queue{
//		list:        linked_list.New(),
//		compression: compression.NewSnappy(),
//	}
//	q.Cond = sync.NewCond(&q.lock)
//	return q
//}
//
//func (q *Queue) Enqueue(value []uint8) error {
//	q.lock.Lock()
//	defer q.lock.Unlock()
//
//	cmpValue, err := q.compression.Compress(value)
//	if err != nil {
//		return err
//	}
//	q.list.PushBack(cmpValue)
//	q.Cond.Signal()
//
//	return nil
//}
//
//func (q *Queue) Dequeue() ([]byte, error) {
//	q.lock.Lock()
//	defer q.lock.Unlock()
//
//	if q.list.Len() == 0 {
//		return nil, fmt.Errorf("queue is empty")
//	}
//
//	element := q.list.Front()
//	q.list.Remove(element)
//
//	decompressed, err := q.compression.Decompress(element.Value)
//	if err != nil {
//		return nil, err
//	}
//
//	return decompressed, nil
//}
//
//func (q *Queue) BlockingDequeue(ctx context.Context) []byte {
//	q.lock.Lock()
//	defer q.lock.Unlock()
//
//	for q.list.Len() == 0 {
//		q.Cond.Wait()
//	}
//
//	if ctx.Err() != nil {
//		return nil
//	}
//
//	element := q.list.Front()
//	decompressed, err := q.compression.Decompress(element.Value)
//	if err != nil {
//		return nil
//	}
//	q.list.Remove(element)
//
//	return decompressed
//}
//
//func (q *Queue) Len() int {
//	q.lock.Lock()
//	defer q.lock.Unlock()
//
//	return q.list.Len()
//}
