package queue

/*
	LIST QUEUE:

	Len()                Enqueue()            Dequeue()            Peek()
	Best:    O(n)        Best:    O(1)        Best:    O(1)        Best:    O(1)
	Average: O(n)        Average: O(1)        Average: O(1)        Average: O(1)
	Worst:   O(n)        Worst:   O(1)        Worst:   O(1)        Worst:   O(1)
*/

type nodeLQ struct {
	data     string  // Stored data
	nextNode *nodeLQ // Pointer to next node
}

// ListQueue - singly linked list based list of elements organized according to the FIFO principle
type ListQueue struct {
	head *nodeLQ // First element in queue
	tail *nodeLQ // Last element in queue
}

// Len returns length of the queue
func (l *ListQueue) Len() int {
	node := l.head
	len := 1
	for {
		if node.nextNode == nil {
			return len
		}
		node = node.nextNode
		len++
	}
}

// Enqueue adds element into queue
func (l *ListQueue) Enqueue(val string) {
	if l.tail == nil {
		l.tail = &nodeLQ{val, nil}
		l.head = l.tail
		return
	}

	l.tail.nextNode = &nodeLQ{val, nil}
	l.tail = l.tail.nextNode
}

// Dequeue obtains and deletes element from queue
func (l *ListQueue) Dequeue() (val string, ok bool) {
	if l.head == nil {
		return "", false
	}

	val = l.head.data
	l.head = l.head.nextNode
	return val, true
}

// Peek obtains element from queue
func (l *ListQueue) Peek() (val string, ok bool) {
	if l.head == nil {
		return "", false
	}

	return l.head.data, true
}
