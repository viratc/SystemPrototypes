package stack

import (
	"sync/atomic"
	"unsafe"
)

type Node struct {
	Val  interface{}
	Next *Node
}

type LockFreeStack struct {
	head unsafe.Pointer
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(value interface{}) {
	newNode := &Node{Val: value}

	for {
		oldHead := atomic.LoadPointer(&s.head)
		newNode.Next = (*Node)(oldHead)

		if atomic.CompareAndSwapPointer(&s.head, oldHead, unsafe.Pointer(newNode)) {
			return
		}
	}
}

func (s *LockFreeStack) Pop() interface{} {
	for {
		oldHead := atomic.LoadPointer(&s.head)
		if oldHead == nil {
			return nil
		}

		oldHeadNode := (*Node)(oldHead)
		newHead := unsafe.Pointer(oldHeadNode.Next)

		if atomic.CompareAndSwapPointer(&s.head, oldHead, newHead) {
			return oldHeadNode.Val
		}
	}
}

func (s *LockFreeStack) Peek() interface{} {
	head := atomic.LoadPointer(&s.head)

	if head == nil {
		return nil
	}

	return (*Node)(head).Val
}

func (s *LockFreeStack) IsEmpty() bool {
	head := atomic.LoadPointer(&s.head)
	return head == nil
}
