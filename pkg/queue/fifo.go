package queue

import (
	"fmt"
	"sync"
)

// FIFO 先入先出队列
type FIFO[T any] struct {
	data []T
	lock sync.RWMutex
}

func NewFIFO[T any]() *FIFO[T] {
	return &FIFO[T]{data: make([]T, 0)}
}

// Push 入队
func (f *FIFO[T]) Push(v T) {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.data = append([]T(f.data), v)
}

// List 队列
func (f *FIFO[T]) List() []T {
	f.lock.RLock()
	defer f.lock.RUnlock()

	list := make([]T, 0, len([]T(f.data)))
	for _, t := range f.data {
		list = append(list, t)
	}

	return list
}

func (f *FIFO[T]) String() string {
	return fmt.Sprintf("%v", f.List())
}
