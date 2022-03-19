# Mugard

Mugard is a simple extension for Mutex in golang. Inspired by Mutex in Rust Language, Mugard is protecting the data (resources) instead of protecting the code, read this [blog posts](https://medium.com/@deckarep/paradigms-of-rust-for-the-go-developer-210f67cd6a29#:~:text=Paradigm%20shift%3A%20Lock%20data%20not%20code) for more informations.

## Requirements
```
Go 1.18
```
Since mugard using Generic Type. It's required to be used with **Go 1.18**

## How to install
```bash
go get github.com/raspiantoro/mugard
```

## Example usage
```go
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
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
