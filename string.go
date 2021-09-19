package mugard

import (
	"fmt"
	"reflect"
)

type stringGuard struct {
	*guard
	val string
}

func StringGuard(val string) (guard *stringGuard, err error) {
	v, err := clone(val, typeString)
	if err != nil {
		return
	}

	guard = &stringGuard{
		guard: newGuard(),
	}

	newVal, err := guard.toString(v)
	if err != nil {
		return
	}

	guard.val = newVal

	return
}

func (s *stringGuard) GetReadLock() (newVal string, err error) {
	v, err := s.guard.getReadLock(s.val, typeString)
	if err != nil {
		return
	}

	newVal, err = s.toString(v)

	return
}

func (s *stringGuard) GetRead() (newVal string, err error) {
	v, err := s.getRead(s.val, typeString)
	if err != nil {
		return
	}

	newVal, err = s.toString(v)

	return
}

func (s *stringGuard) GetWrite() (val *string, err error) {
	err = s.askWrite()
	if err != nil {
		return
	}
	val = &s.val
	return
}

func (i *stringGuard) toString(val interface{}) (newVal string, err error) {
	newVal, ok := val.(string)
	if !ok {
		err = fmt.Errorf("failed to cast %s into string", reflect.ValueOf(val).Kind())
	}

	return
}
