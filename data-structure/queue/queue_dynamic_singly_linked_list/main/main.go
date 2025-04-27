package main

import (
	"fmt"

	queue "github.com/sernopay/go-examples/data-structure/queue/queue_dynamic_singly_linked_list/queue"
)

func main() {
	q := queue.New[int]()
	q.Add(1)
	q.Add(2)
	q.Add(3)

	fmt.Println(q)

	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())

	fmt.Println(q.Poll())
}
