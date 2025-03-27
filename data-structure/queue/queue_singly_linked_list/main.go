package main

import "fmt"

type node[T any] struct {
	value    T
	nextNode *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
	}
}

type myCircularQueue[T any] struct {
	capacity int
	count    int
	head     *node[T]
	tail     *node[T]
}

func New[T any](capacity int) *myCircularQueue[T] {
	return &myCircularQueue[T]{
		capacity: capacity,
	}
}

func (q *myCircularQueue[T]) EnQueue(value T) bool {
	if q.IsFull() {
		return false
	}
	n := newNode(value)
	if q.count == 0 {
		q.head = n
		q.tail = q.head
	} else {
		q.tail.nextNode = n
		q.tail = n
	}
	q.count++
	return true
}

func (q *myCircularQueue[T]) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.head = q.head.nextNode
	q.count--
	return true
}

func (q *myCircularQueue[T]) Front() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}
	return q.head.value, true
}

func (q *myCircularQueue[T]) Rear() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}
	return q.tail.value, true
}

func (q *myCircularQueue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *myCircularQueue[T]) IsFull() bool {
	return q.count == q.capacity
}

func main() {
	q := New[int](5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)

	fmt.Println(q.Front())
	fmt.Println(q.Rear())

	q.DeQueue()
	fmt.Println(q.Front())
}
