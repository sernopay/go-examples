package main

import (
	"fmt"
	"time"
)

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
