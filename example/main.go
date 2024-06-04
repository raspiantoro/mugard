package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/raspiantoro/mugard"
)

func main() {
	guard := mugard.NewGuard(10)

	readResourceOne := guard.Read()

	fmt.Println("[Read] readResourceOne: ", readResourceOne)

	// only modify the value of readResourceOne, without
	// affecting the value inside the Guard
	readResourceOne = 20

	guard.ReadLock(func(val int) {
		// Perform your read operation inside this closure while the read lock is held.
		readResources := val

		// should print 10
		fmt.Println("[ReadLock] readResource: ", readResources)
	})

	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)

		// Able to run multiple ReadLocks concurrently.
		go func() {
			defer wg.Done()
			guard.ReadLock(func(val int) {
				readResources := val

				// should print 10
				fmt.Println("[ReadLock] readResource: ", readResources)
				time.Sleep(500 * time.Millisecond)
			})
		}()
	}

	wg.Wait()
	wg.Add(2)

	go func() {
		defer wg.Done()
		guard.Write(func(val *int) {
			// Perform your write operation inside this closure.
			*val = *val + 5

			fmt.Println("[Write] readResource: ", *val)

			time.Sleep(1 * time.Second)
		})
	}()

	go func() {
		defer wg.Done()

		time.Sleep(500 * time.Millisecond)

		// this will hold until TryWrite finish
		guard.ReadLock(func(val int) {
			readResources := val

			// should print 15
			fmt.Println("[ReadLock] readResource: ", readResources)

		})
	}()

	wg.Wait()
	wg.Add(1)

	errChan := make(chan error)

	go func() {
		defer wg.Done()
		guard.Write(func(val *int) {
			// Perform your write operation inside this closure.
			*val = *val + 5

			fmt.Println("[Write] readResource: ", *val)

			time.Sleep(1 * time.Second)
		})
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)

		// It should return an error when there is another Write,
		// but it does not block the current goroutine because the ReadLock is never called.
		errChan <- guard.TryReadLock(func(val int) {
			// Perform your write operation inside this closure.
			resources := val

			fmt.Println("[TryReadLock] readResource: ", resources)

			time.Sleep(1 * time.Second)
		})
	}()

	err := <-errChan
	if err != nil {
		// Error should be printed
		fmt.Println("[TryReadLock] error: ", err)
	}

	wg.Wait()
	wg.Add(1)

	go func() {
		defer wg.Done()
		guard.Write(func(val *int) {
			// Perform your write operation inside this closure.
			*val = *val + 5

			fmt.Println("[Write] readResource: ", *val)

			time.Sleep(1 * time.Second)
		})
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)

		// It should return an error when there is another Write,
		// but it does not block the current goroutine because the lock is never called.
		errChan <- guard.TryWrite(func(val *int) {
			// Perform your write operation inside this closure.
			*val = *val + 5

			fmt.Println("[Write] readResource: ", *val)

			time.Sleep(1 * time.Second)
		})
	}()

	err = <-errChan
	if err != nil {
		// Error should be printed
		fmt.Println("[TryWrite] error: ", err)
	}

	wg.Wait()
	close(errChan)
}
