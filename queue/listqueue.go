package queue

import "github.com/vpoletaev11/go-data-structures/list"

// ListQueue - singly linked list based list of elements organized according to the FIFO principle
type ListQueue struct {
	elements list.SinglyLinkedList
}

// Len returns length of the queue
func (q *ListQueue) Len() int {
	return q.elements.Len()
}

// Enqueue adds element into queue
func (q *ListQueue) Enqueue(val string) {
	q.elements.InsertTail(val)
}

// Dequeue obtains and deletes element from queue
func (q *ListQueue) Dequeue() (val string, ok bool) {
	val, ok = q.elements.GetHead()
	return val, ok
}

// Peek obtains element from queue
func (q *ListQueue) Peek() (val string, ok bool) {
	val, ok = q.elements.PeekHead()

	return val, ok
}
