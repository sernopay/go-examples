package queue

type node[T any] struct {
	value T
	next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
	}
}

type queue[T any] struct {
	count int
	head  *node[T]
	tail  *node[T]
}

func New[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Add(value T) {
	n := newNode(value)
	if q.count == 0 {
		q.head = n
		q.tail = q.head
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
}

func (q *queue[T]) Poll() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}
	n := q.head
	q.head = q.head.next
	q.count--

	// reset tail if empty
	if q.IsEmpty() {
		q.tail = q.head
	}

	return n.value, true
}

func (q *queue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *queue[T]) Size() int {
	return q.count
}
