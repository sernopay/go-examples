package main

import "fmt"

type MyCircularQueue struct {
	queue     []int
	headIndex int
	count     int
	capacity  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	q := make([]int, k)
	return MyCircularQueue{
		queue:    q,
		capacity: k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[(this.headIndex+this.count)%this.capacity] = value
	this.count++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.headIndex = (this.headIndex + 1) % this.capacity
	this.count--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.headIndex]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	tailIndex := (this.headIndex + this.count - 1) % this.capacity
	return this.queue[tailIndex]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.capacity
}

func main() {
	q := Constructor(3)
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
