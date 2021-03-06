package main

import (
	"fmt"

	"github.com/raspiantoro/mugard"
)

func main() {
	guard := mugard.NewGuard(10)

	readResourceOne := guard.GetRead()

	fmt.Println("readResourceOne: ", readResourceOne)

	// only modify the value of readResourceOne, without
	// affecting the value inside the Guard
	readResourceOne = 20

	readResourceTwo := guard.GetReadLock()

	// should print 10
	fmt.Println("readResourceTwo: ", readResourceTwo)

	guard.ReleaseRead()

	writeResource := guard.GetWrite()

	// modifying value of writeResource also modifying
	// the value inside the guard
	*writeResource = 30

	// data inside the guard already modify by writeResource
	readResourceThree := guard.GetRead()

	// should print 30
	fmt.Println("readResourceThree: ", readResourceThree)

	// releasing write access, so other resources
	// can have the write access
	err := guard.ReleaseWrite(&writeResource)
	if err != nil {
		fmt.Println(err)
	}

	// writeResource is read only now,
	// modify it's value won't affect the data inside the guard
	*writeResource = 100

	readResourceFour := guard.GetRead()

	// should print 30
	fmt.Println("readResourceFour: ", readResourceFour)

	writeResource = guard.GetWrite()

	// won't block the process
	anotherWriteResource, err := guard.TryGetWrite()
	if err != nil {
		fmt.Println(err)
	}

	// should block the process until previous resource that hold the value
	// releasing their write access.
	// in this case it will cause a deadlock, since the previous process not
	// releasing the write access yet
	_ = guard.GetWrite()

	// won't reach
	fmt.Println(anotherWriteResource)
}
