# Mugard

Mugard is a simple extension for Mutex in Go. Inspired by Rust's Mutex, Mugard protects data (resources) rather than code. Read this [blog post](https://medium.com/@deckarep/paradigms-of-rust-for-the-go-developer-210f67cd6a29#:~:text=Paradigm%20shift%3A%20Lock%20data%20not%20code)  for more information.

## Requirements
```
Go 1.18
```
Since mugard uses type parameters, it is required to be used with **Go 1.18** or **newer**.

## How to install
```bash
go get github.com/raspiantoro/mugard
```

## Example usage
```go
func main() {
	guard := mugard.NewGuard(10)

	resources := guard.GetRead()

	fmt.Println("readResources: ", resources)

	// only modify the value of readResourceOne, without
	// affecting the value inside the Guard
	resources = 20

	guard.ReadLock(func(val int) {
		// Perform your read operation inside this closure while the read lock is held.
		resources := val

		// should print 10
		fmt.Println("[ReadLock] readResource: ", resources)
	})

	guard.Write(func(val *int) {
		// Perform your write operation inside this closure.
		*val = *val + 5

		fmt.Println("[Write] readResource: ", *val)

		time.Sleep(1 * time.Second)
	})
}
```

For more examples, refer to this [link](example/main.go)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
