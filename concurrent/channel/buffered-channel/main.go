package main

import (
	"fmt"
	"time"
)

func main() {

	// unbuffered channel
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

	fmt.Println()

	// buffered channel
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
