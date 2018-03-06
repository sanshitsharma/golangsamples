package queue

import (
	"errors"
)

// Queue represent the queue DS
type Queue struct {
	elems []interface{} // Size of the queue is 10
	front int
	rear  int
}

// NewQueue creates a new queue of size
func NewQueue(size int) *Queue {
	return &Queue{
		elems: make([]interface{}, size),
		front: -1,
		rear:  0,
	}
}

// IsFull checks is the queue is full
func (q *Queue) IsFull() bool {
	size := len(q.elems)
	if q.rear%size == q.front {
		return true
	}

	return false
}

// IsEmpty check if queue is empty
func (q *Queue) IsEmpty() bool {
	if q.front == -1 || q.front == q.rear {
		return true
	}

	return false
}

// Enqueue adds an item to the rear of the queue if there is space
// else returns an error
func (q *Queue) Enqueue(elem interface{}) error {
	if q.IsFull() {
		return errors.New("OVERFLOW.. Queue already full")
	}

	q.elems[q.rear] = elem
	q.rear++

	if q.front == -1 {
		q.front = 0
	}

	return nil
}

// Dequeue removes an element from the front of queue
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("UNDERFLOW.. Queue is empty")
	}

	elem := q.elems[q.front]
	q.front++

	return elem, nil
}
