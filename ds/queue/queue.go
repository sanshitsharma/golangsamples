package queue

import (
	"errors"
	list "github.com/sanshitsharma/golangsamples/ds/linked_list"
	"fmt"
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

// Print displays the Queue
func (q *Queue) Print() {
	for _, item := range(q.elems) {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}

// LinkQueue creates a queue using linkedlist
type LinkQueue struct {
	front *list.Node
	rear *list.Node
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		front: nil,
		rear: nil,
	}
}

// Enqueue adds a value to rear of LinkQueue
func (q *LinkQueue) Enqueue(val interface{}) {
	if q.rear == nil {
		q.rear = &list.Node{Data: val, Next: nil}
		q.front = q.rear
		return
	}

	q.rear.Next = &list.Node{Data: val, Next: nil}
	q.rear = q.rear.Next
}

// Dequeue removes a from front of LinkQueue. Returns nil if list is empty
func (q *LinkQueue) Dequeue() interface{} {
	if q.front == nil {
		return nil
	}

	rv := q.front.Data
	q.front = q.front.Next
	return rv
}

// Print displays the LinkQueue
func (q *LinkQueue) Print() {
	curr := q.front
	for curr != nil {
		fmt.Printf("%v ", curr.Data)
		curr = curr.Next
	}

	fmt.Println()
}