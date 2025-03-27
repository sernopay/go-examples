package main

import "fmt"

type MyQueue[T any] struct {
	// store elements
	data []T
	// a pointer to indicate the start position
	pStart int
}

// insert an element into the queue. Return true if the operational is successful
func (q *MyQueue[T]) EnQueue(element T) bool {
	q.data = append(q.data, element)
	return true
}

// delete an element from the queue. return true if the operational is successful
func (q *MyQueue[T]) DeQueue() bool {
	if len(q.data) == 0 {
		return false
	}
	q.pStart++
	return true
}

// Get the front item from the queue
func (q *MyQueue[T]) Front() T {
	return q.data[q.pStart]
}

// Check whether the queue is empty or not
func (q *MyQueue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func main() {
	q := &MyQueue[int]{}

	q.EnQueue(1)
	q.EnQueue(2)
	fmt.Println(q.Front())

	q.DeQueue()
	fmt.Println(q.Front())
}
