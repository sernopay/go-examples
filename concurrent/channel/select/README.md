# Select Example
This example demonstrates how to use the `select` statement to receive data from multiple channels.

## About Select in Go Channels
The `select` statement lets a goroutine wait on multiple communication operations. A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. If no case is ready, it blocks until at least one case is ready. This pattern is useful for :
    - multiplexing input from multiple resources
    - implementing timeouts
    - handling cancellation signals in go programs

## Code Explanation
The example code shows how to launch multiple goroutines that send data into separate channels and use a `select` statement to receive and print the data from these channels.

### group1 Function
```go
func group1(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
}
```
- The `group1` function takes a channel of integers as an argument.
- Inside the function, we use a `for` loop to send integers from 1 to 5 into the channel using `c <- i`.
- After sending all the integers, we close the channel using `close(c)`. Closing the channel is important to signal that no more data will be sent.

### group2 Function
```go
func group2(c chan int) {
	for i := 6; i <= 10; i++ {
		c <- i
	}
	close(c)
}

```
- The `group2` function takes a channel of integers as an argument.
- Inside the function, we use a `for` loop to send integers from 6 to 10 into the channel using `c <- i`.
- After sending all the integers, we close the channel using `close(c)`. Closing the channel is important to signal that no more data will be sent.

### Main Function
```go
func main() {
	cg1 := make(chan int)
	cg2 := make(chan int)

	go group1(cg1)
	go group2(cg2)

	for {
		select {

		case val, ok := <-cg1:
			if ok {
				fmt.Println("print from cg1 ", val)
			} else {
				cg1 = nil
			}

		case val, ok := <-cg2:
			if ok {
				fmt.Println("print from cg2 ", val)
			} else {
				cg2 = nil
			}
		}

		if cg1 == nil && cg2 == nil {
			break
		}
	}
}
```
- In the main function, we create two channels of integers using `make(chan int)`.
- We launch two goroutines to execute the `group1` and `group2` functions by calling `go group1(cg1)` and `go group2(cg2)`.
- We use a `for` loop to continuously select and receive data from the channels `cg1` and `cg2`.
- Inside the `select` statement, we have two cases:
    - The first case receives values from the channel `cg1`. if the channel is open, it prints the value received from `cg1`. If the channel is closed, it sets the channel to `nil`.
    - The second case receives values from the channel `cg2`. if the channel is open, it prints the value received from `cg2`. If the channel is closed, it sets the channel to `nil`.
- The loop continues until both channels are nil, indicating that all data has been received and both channels are closed.

## Running the Code
To run the code, use the `go run` command followed by the filename.
```bash
go run main.go
```
The output will be:
```bash
print from cg1 1
print from cg2 6
print from cg1 2
print from cg2 7
print from cg1 3
print from cg2 8
print from cg1 4
print from cg2 9
print from cg1 5
print from cg2 10
```

## Summary
This example demonstrates how to use the `select` statement to receive data from multiple channels concurrently. The `select` statement allows you to wait on multiple communication operations and execute the case that is ready to run.