# Unbuffered Channel VS Buffered Channel

## Learning Objectives
- Understand the difference between unbuffered and buffered channels in Go.
- Learn how to create unbuffered and buffered channels.
- When to use unbuffered channels and when to use buffered channels.

## Unbuffered Channel
Channels in Go are used to communicate between goroutines. They can be either buffered or unbuffered. By default, if you create a channel without specifying a buffer size, it is an unbuffered channel. For example:
```go
ch := make(chan int) // unbuffered channel
```
An unbuffered channel means that the sender and receiver must be synchronized. It means that the sender will block sending data until the receiver is ready to receive it. Similarly, the receiver will block until there is data to receive. This is useful when you want to ensure that the sender and receiver are synchronized.

Let's look at an example of an unbuffered channel:
```go
func main() {
	uChan := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i, "sending data to unbuffered channel")
			uChan <- i
		}
	}()

	time.Sleep(50 * time.Millisecond)
	fmt.Println(<-uChan, "received data from unbuffered channel")
	time.Sleep(50 * time.Millisecond)
	fmt.Println(<-uChan, "received data from unbuffered channel")
	time.Sleep(50 * time.Millisecond)
	fmt.Println(<-uChan, "received data from unbuffered channel")
}
```
In this example, we create an unbuffered channel `uChan` and start a goroutine that sends data to the channel. I set a sleep of 50 milliseconds before receiving data from the channel to simulate the blocking behavior of the unbuffered channel. The output will be:
```bash
0 sending data to unbuffered channel
0 received data from unbuffered channel
1 sending data to unbuffered channel
1 received data from unbuffered channel
2 sending data to unbuffered channel
2 received data from unbuffered channel
```
As we can see, the sender and receiver are synchronized. After sending data to the channel, the sender blocks until the receiver is ready to receive it. We can see from the output that the sender is not able to send data until the receiver receives it. Similarly, the receiver blocks until there is data to receive. This is the behavior of an unbuffered channel.

## Buffered Channel
Different from unbuffered channels, buffered channels allow you to send data without waiting for the receiver to be ready, as long as the buffer is not full. You can create a buffered channel by specifying a buffer size when creating the channel. For example:
```go
ch := make(chan int, 2)
```
From the above example, we create a buffered channel with a buffer size of 2. This means that the channel can hold up to 2 values before it blocks the sender. If the buffer is full, the sender will block until there is space in the buffer. This is useful when you want to allow the sender to send data without waiting for the receiver to be ready, as long as the buffer is not full.

Let's look at an example of a buffered channel:
```go
func main() {
    ch := make(chan int, 2)

    go func() {
        for i := 0; i < 3; i++ {
            ch <- i
            fmt.Println(i, "sending data to buffered channel")
        }
    }()

    time.Sleep(50 * time.Millisecond)
    fmt.Println(<-ch, "received data from buffered channel")
    time.Sleep(50 * time.Millisecond)
    fmt.Println(<-ch, "received data from buffered channel")
    time.Sleep(50 * time.Millisecond)
    fmt.Println(<-ch, "received data from buffered channel")
}
```
We use the same example as the unbuffered channel, but this time we create a buffered channel with a buffer size of 2. The output will be:
```bash
0 sending data to buffered channel
1 sending data to buffered channel
0 received data from buffered channel
2 sending data to buffered channel
1 received data from buffered channel
2 received data from buffered channel
```
As we can see, the sender is able to send data to the channel with maximum 2 values in the buffer without waiting for the receiver to be ready. The sender blocks only when the buffer is full. Once the receiver retrieves data from the channel, the sender can send more data. This is the behavior of a buffered channel. It is different from the unbuffered channel where the sender blocks until the receiver is ready to receive it. If a buffered channel has a buffer size of 0, it behaves like an unbuffered channel.

## When to use Unbuffered Channels and When to use Buffered Channels
Unbuffered channels are useful when you want to ensure that the sender and receiver are synchronized. They are best suited for coordination and signaling between goroutines. For example, goroutine synchronization, pipelining/staging tasks, etc. \
Buffered channels are useful when you want to allow the sender to send data without waiting for the receiver to be ready, as long as the buffer is not full. They are best suited for scenarios where you want to decouple the sender and receiver. For example, Decouple Producer and Consumer, Job Queue/Task Buffering, Rate Limiting, etc.

## Summary
Unbuffered channels are used for synchronization between goroutines, while buffered channels are used for decoupling the sender and receiver. Unbuffered channels block the sender until the receiver is ready, while buffered channels allow the sender to send data without waiting for the receiver to be ready, as long as the buffer is not full. Understanding the difference between unbuffered and buffered channels helps you choose the right channel type for your use case.