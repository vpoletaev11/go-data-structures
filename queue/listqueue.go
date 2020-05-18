package queue

import "github.com/vpoletaev11/go-data-structures/list"

/*
	LIST QUEUE:

	Len()                Enqueue()            Dequeue()            Peek()
	Best:    O(n)        Best:    O(n)        Best:    O(1)        Best:    O(1)
	Average: O(n)        Average: O(n)        Average: O(1)        Average: O(1)
	Worst:   O(n)        Worst:   O(n)        Worst:   O(1)        Worst:   O(1)
*/

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
