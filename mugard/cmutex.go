package mugard

import (
	"reflect"
	"sync"
)

type Mugard struct {
	sync.Mutex
	val     interface{}
	counter uint
}

func Guard(val interface{}) (*Mugard, error) {
	newVal, err := getNewVal(val)
	if err != nil {
		return nil, err
	}

	mugard := &Mugard{
		val:     newVal,
		counter: 0,
		Mutex:   sync.Mutex{},
	}

	return mugard, nil
}

func getNewVal(val interface{}) (newVal interface{}, err error) {
	if val == nil {
		err = ErrNilValue
	}

	newVal = reflect.New(reflect.ValueOf(val).Elem().Type()).Interface()

	return
}

func (m *Mugard) Get() interface{} {
	m.Lock()
	mugard, _ := getNewVal(m.val)
	return mugard
}
