package queue

// Queue - list of items organized according to the FIFO principle
type Queue struct {
	items []string
}

// Len returns length of the queue
func (q *Queue) Len() int {
	return len(q.items)
}

// Enqueue adds element into queue
func (q *Queue) Enqueue(elem string) {
	q.items = append(q.items, elem)
}

// Dequeue obtains and deletes element from queue
func (q *Queue) Dequeue() (string, bool) {
	if q.Len() == 0 {
		return "", false
	}

	elem := q.items[0]
	q.items = q.items[1:]
	return elem, true
}

// Peek obtains element from queue
func (q *Queue) Peek() (string, bool) {
	if q.Len() == 0 {
		return "", false
	}

	return q.items[0], true
}
