package mugard

import (
	"fmt"
	"reflect"
)

type int8Guard struct {
	*guard
	val int8
}

func Int8Guard(val int8) (guard *int8Guard, err error) {
	v, err := clone(val, typeInt8)
	if err != nil {
		return
	}

	guard = &int8Guard{
		guard: newGuard(),
	}

	newVal, err := guard.toInt8(v)
	if err != nil {
		return
	}

	guard.val = newVal

	return
}

func (i *int8Guard) GetReadLock() (newVal int8, err error) {
	v, err := i.guard.getReadLock(i.val, typeInt8)
	if err != nil {
		return
	}

	newVal, err = i.toInt8(v)

	return
}

func (i *int8Guard) GetRead() (newVal int8, err error) {
	v, err := i.getRead(i.val, typeInt8)
	if err != nil {
		return
	}

	newVal, err = i.toInt8(v)

	return
}

func (i *int8Guard) GetWrite() (val *int8, err error) {
	err = i.askWrite()
	if err != nil {
		return
	}
	val = &i.val
	return
}

func (i *int8Guard) toInt8(val interface{}) (newVal int8, err error) {
	newVal, ok := val.(int8)
	if !ok {
		err = fmt.Errorf("failed to cast %s into int8", reflect.ValueOf(val).Kind())
	}

	return
}
