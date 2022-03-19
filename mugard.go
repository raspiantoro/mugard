package mugard

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type Guard[T any] struct {
	*sync.RWMutex
	counter uint64
	val     T
}

func NewGuard[T any](val T) *Guard[T] {
	return &Guard[T]{
		RWMutex: &sync.RWMutex{},
		val:     val,
		counter: 0,
	}
}

func (g *Guard[T]) GetRead() T {
	return g.val
}

func (g *Guard[T]) GetReadLock() T {
	g.RWMutex.RLock()
	return g.val
}

func (g *Guard[T]) ReleaseRead() {
	g.RWMutex.RUnlock()
}

func (g *Guard[T]) TryGetWrite() (*T, error) {
	if g.counter > 0 {
		return nil, ErrMultipleWrite
	}

	g.RWMutex.Lock()
	atomic.AddUint64(&g.counter, 1)
	return &g.val, nil
}

func (g *Guard[T]) GetWrite() *T {
	g.RWMutex.Lock()
	atomic.AddUint64(&g.counter, 1)
	return &g.val
}

func (g *Guard[T]) ReleaseWrite(holders **T) error {
	if g.counter <= 0 {
		return ErrNoLockedResources
	}

	// if the holders keep the write access, it's should be
	// pointing to g.val address
	if unsafe.Pointer(*holders) != unsafe.Pointer(&g.val) {
		return ErrNotHoldingWriteAccess
	}

	// put temp on the heap, so it's still there
	// when this function goes out of scope
	temp := new(T)
	*temp = g.val

	// make holders as read only by changing the pointing address
	*holders = temp

	g.RWMutex.Unlock()
	atomic.AddUint64(&g.counter, ^uint64(0))

	return nil
}
