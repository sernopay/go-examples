# Buffered Channel Example

This example demonstrates how to use buffered channels in Go.

## About Buffered Channels
Buffered channels are a type of channel that can store a fixed number of elements. When the buffer is full, sending a value to the channel will block until there is space available in the buffer. Similarly, when the buffer is empty, receiving a value from the channel will block until there is a value available in the buffer.

## Code Explanation
The following Go code creates a buffered channel, launches a goroutine to send data into the channel, and then receives the data from the main function:
```go
func main() {
    ch := make(chan int, 2)

    go func() {
        for i := 0; i < 3; i++ {
            ch <- i
            fmt.Println(i, " sent")
        }
    }()

    time.Sleep(100 * time.Millisecond)
    fmt.Println(<-ch, " receive 0")
    time.Sleep(500 * time.Millisecond)
    fmt.Println(<-ch, " receive 1")
    fmt.Println(<-ch, " receive 2")
}
```
* We create a buffered channel `ch` with a capacity of `2`. This means the channel can hold up to 2 integers before blocking.
* We launch a goroutine that sends 3 integers to the channel. Since the channel has a buffer size of 2, the first two values are sent without blocking. The third value will block until the buffer has space.
* The first receive operation will unblock immediately since there is a value in the buffer.
* We add a delay of 100 milliseconds to allow the goroutine to send the first and second values so we can see it blocked on the third value.
* We add another delay of 500 milliseconds to allow the goroutine to send the third value and the main function to receive the remaining values.

## Running the Code
To run the code, use the `go run` command:
```bash
go run main.go
```
You should see the following output:
```bash
0  sent
1  sent
0  receive 0
2  sent
1  receive 1
2  receive 2
```

## Summary
Buffered channels allow you to send and receive data without immediate synchronization, providing a way to control the flow of data between concurrent operations. By specifying a buffer size, you can control how many elements can be stored in the channel before blocking occurs. This can be useful for scenarios where you want to decouple the sending and receiving operations or when you need to limit the number of elements in transit.
