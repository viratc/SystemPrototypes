package queues

import "sync"

type BlockingQueue struct {
	m        sync.Mutex
	c        sync.Cond
	data     []interface{}
	capacity int
}

func new_blocking_queue(capacity int) *BlockingQueue {
	q := new(BlockingQueue)
	q.c = sync.Cond{L: &q.m}
	q.capacity = capacity
	return q
}
