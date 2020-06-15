package stack

import "github.com/vpoletaev11/go-data-structures/list"

/*
	LIST STACK:

	Len()                Push()               Pop()                Peek()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(1)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(1)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(1)


	goos: linux
	goarch: amd64
	BenchmarkPushListStack-8         1000000               563 ns/op             320 B/op         10 allocs/op
	BenchmarkPopListStack-8          1000000                74.2 ns/op             0 B/op          0 allocs/op
	BenchmarkPeekListStack-8         1000000                 0.647 ns/op           0 B/op          0 allocs/op
*/

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
