package queues

import "sync"

type BlockingQueue struct {
	m        sync.Mutex
	c        sync.Cond
	data     []interface{}
	capacity int
}

// Constructor
func NewBlockingQueue(capacity int) *BlockingQueue {
	q := new(BlockingQueue)
	q.c = sync.Cond{L: &q.m}
	q.capacity = capacity
	return q
}

func (q *BlockingQueue) isFull() bool {
	return len(q.data) == q.capacity
}

func (q *BlockingQueue) Enqueue(item interface{}) {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	for q.isFull() {
		q.c.Wait()
	}

	q.data = append(q.data, item)
	q.c.Signal()
}

func (q *BlockingQueue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *BlockingQueue) Dequeue() interface{} {

	q.c.L.Lock()
	defer q.c.L.Unlock()

	for q.isEmpty() {
		q.c.Wait()
	}

	result := q.data[0]

	// Remove the first element from FIFO queue
	q.data = q.data[1:len(q.data)]

	return result
}
