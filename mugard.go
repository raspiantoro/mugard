package mugard

import (
	"errors"
	"sync"
)

var (
	ErrMultipleWrite = errors.New("another process has already acquired the write lock")
)

type readFunc[T any] func(val T)
type writeFunc[T any] func(val *T)

type Guard[T any] struct {
	mu  *sync.RWMutex
	val T
}

func NewGuard[T any](val T) *Guard[T] {
	return &Guard[T]{
		mu:  &sync.RWMutex{},
		val: val,
	}
}

func (g *Guard[T]) Read() T {
	return g.val
}

func (g *Guard[T]) TryReadLock(fn readFunc[T]) error {
	if !g.mu.TryRLock() {
		return ErrMultipleWrite
	}
	defer g.mu.RUnlock()

	fn(g.val)

	return nil
}

func (g *Guard[T]) ReadLock(fn readFunc[T]) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	fn(g.val)
}

func (g *Guard[T]) TryWrite(fn writeFunc[T]) error {
	if !g.mu.TryLock() {
		return ErrMultipleWrite
	}
	defer g.mu.Unlock()

	fn(&g.val)

	return nil
}

func (g *Guard[T]) Write(fn writeFunc[T]) {
	g.mu.Lock()
	defer g.mu.Unlock()
	fn(&g.val)
}
