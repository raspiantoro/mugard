package cmutex

import (
	"reflect"
	"sync"
)

type Cmutex struct {
	sync.Mutex
	val     interface{}
	counter uint
}

func Guard(val interface{}) *Cmutex {
	newVal := reflect.New(reflect.ValueOf(val).Elem().Type()).Interface()
	return &Cmutex{
		val:     newVal,
		counter: 0,
		Mutex:   sync.Mutex{},
	}
}

func getNewVal(val interface{}) (newVal interface{}, err error) {
	if val == nil {

	}
}

func (c *Cmutex) Get() interface{} {
	c.Lock()
	return c.val
}
