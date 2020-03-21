package stack

import "github.com/vpoletaev11/go-data-structures/list"

// ListStack - singly linked list based list of elements organized according to the LIFO principle
type ListStack struct {
	elements list.SinglyLinkedList
}

// Len returns length of the stack
func (s *ListStack) Len() int {
	return s.elements.Len()
}

// Push adds element into stack
func (s *ListStack) Push(val string) {
	s.elements.InsertHead(val)
}

// Pop obtains and deletes element from stack
func (s *ListStack) Pop() (val string, ok bool) {
	val, ok = s.elements.GetHead()
	return val, ok
}

// Peek obtains element from stack
func (s *ListStack) Peek() (val string, ok bool) {
	val, ok = s.elements.PeekHead()
	return val, ok
}
