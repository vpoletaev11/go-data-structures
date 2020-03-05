package queue

import "fmt"

// Queue - list of elements organized according to the FIFO principle
type Queue []string

// Enqueue adds element into queue
func (q *Queue) Enqueue(elem string) {
	*q = append(*q, elem)
}

// Dequeue obtains and deletes element from queue
func (q *Queue) Dequeue() (string, bool) {
	if len(*q) == 0 {
		return "", false
	}

	elem := (*q)[0]
	*q = (*q)[1:]
	return elem, true
}

// Peek obtains element from queue
func (q *Queue) Peek() (string, bool) {
	if len(*q) == 0 {
		return "", false
	}

	return (*q)[0], true
}

// Demo displays example of using the queue data structure
func Demo() {
	fmt.Println("Demonstration of using the queue data structure:")
	var queue Queue

	fmt.Println("Enqueued: one")
	queue.Enqueue("one")

	fmt.Println("Enqueued: two")
	queue.Enqueue("two")

	fmt.Println("Enqueued: three")
	queue.Enqueue("three")
	fmt.Print("\n")

	elem, ok := queue.Peek()
	if ok {
		fmt.Println("Peeked:", elem)
	}

	for len(queue) > 0 {
		elem, ok := queue.Dequeue()
		if ok {
			fmt.Println("Dequeued:", elem)
		}
	}
}
