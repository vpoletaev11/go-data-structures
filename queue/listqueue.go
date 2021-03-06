package queue

/*
	LIST QUEUE:

	Len()                Enqueue()            Dequeue()            Peek()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(1)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(1)
	Worst:   O(1)        Worst:   O(1)        Worst:   O(1)        Worst:   O(1)

	
	goos: linux
	goarch: amd64
	BenchmarkEnqueueListQueue-8       	 1701579	       708 ns/op	     320 B/op	      10 allocs/op
	BenchmarkDequeueListQueue-8       	15704728	        76.5 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPeekListQueue-8          	1000000000	         0.297 ns/op	       0 B/op	       0 allocs/op
*/

type nodeLQ struct {
	data     string  // Stored data
	nextNode *nodeLQ // Pointer to next node
}

// ListQueue - singly linked list based list of elements organized according to the FIFO principle
type ListQueue struct {
	head *nodeLQ // First element in queue
	tail *nodeLQ // Last element in queue
	len  int     // Length of the queue
}

// Len returns length of the queue
func (l *ListQueue) Len() int {
	return l.len
}

// Enqueue adds element into queue
func (l *ListQueue) Enqueue(val string) {
	if l.tail == nil {
		l.tail = &nodeLQ{val, nil}
		l.head = l.tail
		l.len++
		return
	}

	l.tail.nextNode = &nodeLQ{val, nil}
	l.tail = l.tail.nextNode
	l.len++
}

// Dequeue obtains and deletes element from queue
func (l *ListQueue) Dequeue() (val string, ok bool) {
	if l.head == nil {
		return "", false
	}

	val = l.head.data
	l.head = l.head.nextNode
	l.len--
	return val, true
}

// Peek obtains element from queue
func (l *ListQueue) Peek() (val string, ok bool) {
	if l.head == nil {
		return "", false
	}

	return l.head.data, true
}
