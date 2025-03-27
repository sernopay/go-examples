package main

import (
	"fmt"
	"sync"
)

type myCircularQueue[T any] struct {
	queue     []T
	headIndex int
	count     int
	capacity  int
	mu        sync.Mutex
}

func New[T any](c int) *myCircularQueue[T] {
	queue := make([]T, c)
	return &myCircularQueue[T]{
		queue:    queue,
		capacity: c,
	}
}

func (this *myCircularQueue[T]) EnQueue(value T) bool {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.IsFull() {
		return false
	}
	this.queue[(this.headIndex+this.count)%this.capacity] = value
	this.count++
	return true
}

func (this *myCircularQueue[T]) DeQueue() bool {
	this.mu.Lock()
	defer this.mu.Unlock()

	if this.IsEmpty() {
		return false
	}

	this.headIndex = (this.headIndex + 1) % this.capacity
	this.count--
	return true
}

func (this *myCircularQueue[T]) Front() (T, bool) {
	if this.IsEmpty() {
		return *new(T), false
	}
	return this.queue[this.headIndex], true
}

func (this *myCircularQueue[T]) Rear() (T, bool) {
	if this.IsEmpty() {
		return *new(T), false
	}
	tailIndex := (this.headIndex + this.count - 1) % this.capacity
	return this.queue[tailIndex], true
}

func (this *myCircularQueue[T]) IsEmpty() bool {
	return this.count == 0
}

func (this *myCircularQueue[T]) IsFull() bool {
	return this.count == this.capacity
}

func main() {
	q := New[int](3)
	fmt.Println(q.EnQueue(1)) // true
	fmt.Println(q.EnQueue(2)) // true
	fmt.Println(q.EnQueue(3)) // true
	fmt.Println(q.EnQueue(4)) // false
	fmt.Println(q.Rear())     // 3
	fmt.Println(q.IsFull())   // true
	fmt.Println(q.DeQueue())  // true
	fmt.Println(q.Front())    // 2
	fmt.Println(q.EnQueue(4)) // true
	q.DeQueue()
	q.DeQueue()
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	q.DeQueue()
	q.EnQueue(80)
	fmt.Println(q)
}
