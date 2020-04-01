package queue

import "github.com/vpoletaev11/go-data-structures/list"

// ListQueue - singly linked list based list of elements organized according to the FIFO principle
type ListQueue struct {
	elements list.SinglyLinkedList
}

// Len returns length of the queue
func (l *ListQueue) Len() int {
	return l.elements.Len()
}

// Enqueue adds element into queue
func (l *ListQueue) Enqueue(val string) {
	l.elements.InsertTail(val)
}

// Dequeue obtains and deletes element from queue
func (l *ListQueue) Dequeue() (val string, ok bool) {
	return l.elements.GetHead()
}

// Peek obtains element from queue
func (l *ListQueue) Peek() (val string, ok bool) {
	return l.elements.PeekHead()
}
