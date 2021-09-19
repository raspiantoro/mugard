package mugard

import (
	"reflect"
	"sync"
	"sync/atomic"
)

type valueType uint64

const (
	typeString valueType = iota
	typeInt8
	typeInt16
	typeInt32
	typeInt64
	typeUint8
	typeUint16
	typeUint32
	typeUint64
	typeSlice
	typeInterface
)

func clone(val interface{}, valType valueType) (newVal interface{}, err error) {
	if val == nil {
		err = ErrNilValue
		return
	}

	switch valType {
	case typeString:
		newVal = reflect.Indirect(reflect.ValueOf(val)).Interface().(string)
		return
	case typeInt8:
		newVal = reflect.Indirect(reflect.ValueOf(val)).Interface().(int8)
		return
	}

	return
}

type guard struct {
	sync.RWMutex
	val     interface{}
	counter uint64
}

func newGuard() *guard {
	return &guard{
		counter: 0,
		RWMutex: sync.RWMutex{},
	}
}

func (g *guard) getReadLock(val interface{}, valType valueType) (newVal interface{}, err error) {
	g.RLock()
	newVal, err = g.getRead(val, valType)
	return
}

func (g *guard) getRead(val interface{}, valType valueType) (newVal interface{}, err error) {
	newVal, err = clone(val, valType)
	return
}

func (g *guard) askWrite() (err error) {
	if g.counter > 0 {
		err = ErrMultipleWrite
		return
	}

	g.Lock()
	atomic.AddUint64(&g.counter, 1)
	return
}

func (g *guard) ReleaseWrite() {
	g.Unlock()
	atomic.AddUint64(&g.counter, ^uint64(0))
}

func (g *guard) ReleaseRead() {
	g.RUnlock()
}
