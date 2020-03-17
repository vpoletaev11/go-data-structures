package linkedlist

import "fmt"

// node of linked list
type node struct {
	data     string // stored data
	nextNode *node  // pointer to next node
}

// LinkedList -  data structure consisting of a collection of nodes,
// where every node contain pointer to next node, which together represent a sequence
type LinkedList struct {
	head *node // first element in linked list
}

// InsertHead adds value in begin of list
func (l *LinkedList) InsertHead(val string) {
	if l.head == nil {
		l.head = &node{val, nil}
		return
	}

	l.head = &node{val, l.head}
}

// Insert adds value in list by index.
// If on index already exists value, it and all subsequent will be moved by 1, and in his place will be inserted value
func (l *LinkedList) Insert(index int, val string) (err error) {
	if index < 0 {
		return fmt.Errorf("Cannot use negative index: %v", index)
	}

	if index == 0 {
		if l.head == nil {
			l.head = &node{val, nil}
		}

		l.head = &node{val, l.head}
	}

	prevNode := l.nodeByIndex(index - 1)
	if prevNode == nil || prevNode.nextNode == nil {
		return fmt.Errorf("Index out of range")
	}

	prevNode.nextNode = &node{val, prevNode.nextNode}

	return nil
}

// InsertTail adds value in end of list
func (l *LinkedList) InsertTail(val string) {
	if l.head == nil {
		l.head = &node{val, nil}
		return
	}

	lastNode := l.findLastNode()
	lastNode.nextNode = &node{val, nil}
}

// GetHead obtains and deletes element from begin of list
func (l *LinkedList) GetHead() (val string, status bool) {
	if l.head == nil {
		return "", false
	}

	val = l.head.data
	l.head = l.head.nextNode
	return val, true
}

// Get obtains and deletes element by index from list
func (l *LinkedList) Get(index int) (val string, status bool) {
	if l.head == nil {
		return "", false
	}

	if index < 0 {
		return "", false
	}

	if index == 0 {
		val = l.head.data
		l.head = l.head.nextNode
		return val, true
	}

	prevNode := l.nodeByIndex(index - 1)
	if prevNode == nil || prevNode.nextNode == nil {
		return "", false
	}

	val = prevNode.nextNode.data

	prevNode.nextNode = prevNode.nextNode.nextNode

	return val, true
}

// GetTail obtains and deletes element from end of list
func (l *LinkedList) GetTail() (val string, status bool) {
	if l.head == nil {
		return "", false
	}

	p1, p2 := l.head, l.head
	for {
		if p2.nextNode == nil {
			break
		}

		p1 = p2
		p2 = p2.nextNode
	}

	val = p2.data
	p1.nextNode = nil

	return val, true
}

// PeekHead obtains element from begin of list
func (l LinkedList) PeekHead() (val string, status bool) {
	if l.head == nil {
		return "", false
	}

	return l.head.data, true
}

// Peek obtains element by index from list
func (l LinkedList) Peek(index int) (val string, status bool) {
	if l.head == nil {
		return "", false
	}

	node := l.nodeByIndex(index)
	if node == nil {
		return "", false
	}

	return node.data, true
}

// PeekTail obtains element from end of list
func (l LinkedList) PeekTail() (val string, status bool) {
	if l.head == nil {
		return "", false
	}

	lastNode := l.findLastNode()

	return lastNode.data, true
}

// findLastNode finds last element in list
func (l *LinkedList) findLastNode() (pointer *node) {
	pointer = l.head
	for {
		if pointer.nextNode == nil {
			return pointer
		}
		pointer = pointer.nextNode
	}
}

// nodeByIndex returns node from list by index
func (l *LinkedList) nodeByIndex(index int) (pointer *node) {
	pointer = l.head
	if pointer == nil {
		return nil
	}

	for i := 0; i < index; i++ {
		if pointer.nextNode == nil {
			return nil
		}
		pointer = pointer.nextNode
	}
	return pointer
}
