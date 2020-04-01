package stack

import "github.com/vpoletaev11/go-data-structures/list"

// ListStack - singly linked list based list of elements organized according to the LIFO principle
type ListStack struct {
	elements list.SinglyLinkedList
}

// Len returns length of the stack
func (l *ListStack) Len() int {
	return l.elements.Len()
}

// Push adds element into stack
func (l *ListStack) Push(val string) {
	l.elements.InsertHead(val)
}

// Pop obtains and deletes element from stack
func (l *ListStack) Pop() (val string, ok bool) {
	return l.elements.GetHead()
}

// Peek obtains element from stack
func (l *ListStack) Peek() (val string, ok bool) {
	return l.elements.PeekHead()
}
