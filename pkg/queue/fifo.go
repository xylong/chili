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

// Pop 出队
func (f *FIFO[T]) Pop() (T, bool) {
	f.lock.Lock()
	defer f.lock.Unlock()

	var t T
	if len([]T(f.data)) == 0 {
		return t, false
	}

	pop := f.data[0]
	f.data = f.data[1:]
	return pop, true
}

// Get 获取指定位置
func (f *FIFO[T]) Get(index int) (T, bool) {
	f.lock.RLock()
	defer f.lock.RUnlock()

	var t T
	if index < 0 || index >= len([]T(f.data)) {
		return t, false
	}

	return f.data[index], true
}

func (f *FIFO[T]) Length() int {
	return len([]T(f.data))
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
