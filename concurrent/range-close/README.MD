# Range Close Example

This example demonstrates how to use the `range` function to iterate over a channel until it is closed.

## About Range and Close in Go Channels
We can use the `range` keyword to iterate over a channel. The `range` keyword will continue to iterate over the channel until the channel is closed. Sending a value to a closed channel will panic, but receiving from a closed channel will return the zero value immediately. Receiver can use the second value returned by the receive operation to check if the channel is closed.
```go
v, ok := <-ch
```
Note: Only the sender should close a channel, never the receiver. Channel are not like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

## Code Explanation
The following Go code defines a function to fill a channel with integers, launches a goroutine to execute this function, and then receives the data from the channel in the main function

### fillChannel Function:
```go
func fillChannel(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
```
- The `fillChannel` function takes a channel of integers as an argument.
- Inside the function, we use a `for` loop to send integers from 0 to 9 into the channel using `c <- i`.
- After sending all the integers, we close the channel using `close(c)`. Closing the channel is important to signal that no more data will be sent.

### Main Function:
```go
func main() {
	c := make(chan int)
	go fillChannel(c)
	for val := range c {
		fmt.Println(val)
	}
}
```
- In the main function, we create a channel of integers using `make(chan int)`.
- We launch a goroutine to execute the `fillChannel` function by calling `go fillChannel(c)`.
- We use a `range` loop to iterate over the channel `c`. The loop will continue until the channel is closed.
- Inside the loop, we receive the values from the channel and print them using `fmt.Println(val)`.

## Running the Code
To run the code, use the `go run` command followed by the filename.
```bash
go run main.go
```
The output will be:
```bash
0
1
2
3
4
5
6
7
8
9
```

## Summary
This example demonstrates how to use the `range` function to iterate over a channel until it is closed. The sender should close the channel to signal that no more data will be sent. 