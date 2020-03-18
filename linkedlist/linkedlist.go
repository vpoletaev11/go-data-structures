package linkedlist

import "fmt"

// node of linked list
type node struct {
	data     string // Stored data
	nextNode *node  // Pointer to next node
}

// LinkedList -  data structure consisting of a collection of nodes,
// where every node contain pointer to next node, which together represent a sequence.
type LinkedList struct {
	head *node // First element in linked list
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
// If on index already exists value, it and all subsequent values will be moved by 1, and in its place will be inserted new value
func (l *LinkedList) Insert(index int, val string) (err error) {
	if index < 0 {
		return fmt.Errorf("Cannot use negative index: %v", index)
	}

	if index == 0 {
		l.InsertHead(val)
		return nil
	}

	prevNode := l.nodeByIndex(index - 1) // Find node before of candidate node to be inserted
	if prevNode == nil {
		return fmt.Errorf("Index out of range")
	}
	prevNode.nextNode = &node{val, prevNode.nextNode} // prevNode now pointing on inserted node which pointing to node who was here before

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
func (l *LinkedList) Get(index int) (val string, err error) {
	if l.head == nil {
		return "", fmt.Errorf("Empty list")
	}

	if index < 0 {
		return "", fmt.Errorf("Cannot use negative index: %v", index)
	}

	if index == 0 {
		// Similar to GetHead, but without check for empty of list
		val = l.head.data
		l.head = l.head.nextNode
		return val, nil
	}

	prevNode := l.nodeByIndex(index - 1) // Find node before of candidate node to be obtained and deleted
	if prevNode == nil || prevNode.nextNode == nil {
		return "", fmt.Errorf("Index out of range")
	}
	val = prevNode.nextNode.data                   // Obtain node data
	prevNode.nextNode = prevNode.nextNode.nextNode // prevNode now pointing to next node after deleted node.

	return val, nil
}

// GetTail obtains and deletes element from end of list
func (l *LinkedList) GetTail() (val string, status bool) {
	if l.head == nil {
		return "", false
	}

	p1, p2 := l.head, l.head // Initialize variables that stores pointers to the penultimate and last nodes respectively
	for {
		if p2.nextNode == nil {
			break
		}

		p1 = p2
		p2 = p2.nextNode
	}
	val = p2.data
	p1.nextNode = nil // Penultimate node points to nil i.e now it last node

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
func (l LinkedList) Peek(index int) (val string, err error) {
	if l.head == nil {
		return "", fmt.Errorf("Empty list")
	}

	node := l.nodeByIndex(index)
	if node == nil {
		return "", fmt.Errorf("Index out of range")
	}

	return node.data, nil
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
		if pointer.nextNode == nil { // Check if last node founded before of index node
			return nil
		}
		pointer = pointer.nextNode
	}
	return pointer
}
