package main

import (
	"fmt"

	"github.com/raspiantoro/mugard"
)

func main() {
	testString()
	testInt8()
}

func testString() {
	dataVal := "test data"
	data, err := mugard.StringGuard(dataVal)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Guard instance: %+v\n", data)

	readOne, err := data.GetRead()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read one: %+v\n", readOne)

	readOne = "Change the text"

	fmt.Printf("Read one: %+v\n", readOne)

	readTwo, err := data.GetReadLock()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read two: %+v\n", readTwo)

	readSix, err := data.GetReadLock()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read six: %+v\n", readSix)

	data.ReleaseRead()
	data.ReleaseRead()

	readThree, err := data.GetWrite()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read three: %s\n", *readThree)

	_, err = data.GetWrite()
	if err != nil {
		fmt.Println(err)
	}

	data.ReleaseWrite()

	readFive, err := data.GetWrite()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read five: %s\n", *readFive)

	*readFive = "change now"

	fmt.Printf("Read one: %+v\n", readOne)
	fmt.Printf("Read two: %+v\n", readTwo)
	fmt.Printf("Read three: %s\n", *readThree)
	fmt.Printf("Read five: %s\n", *readFive)

	data.ReleaseWrite()
}

func testInt8() {
	dataVal := int8(10)
	data, err := mugard.Int8Guard(dataVal)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Guard instance: %+v\n", data)

	readOne, err := data.GetRead()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read one: %+v\n", readOne)

	readOne = 5

	fmt.Printf("Read one: %+v\n", readOne)

	readTwo, err := data.GetRead()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read two: %+v\n", readTwo)

	readThree, err := data.GetWrite()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read three: %d\n", *readThree)

	_, err = data.GetWrite()
	if err != nil {
		fmt.Println(err)
	}

	data.ReleaseWrite()

	readFive, err := data.GetWrite()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read five: %d\n", *readFive)

	*readFive = 2

	fmt.Printf("Read one: %+v\n", readOne)
	fmt.Printf("Read two: %+v\n", readTwo)
	fmt.Printf("Read three: %d\n", *readThree)
	fmt.Printf("Read five: %d\n", *readFive)
}
