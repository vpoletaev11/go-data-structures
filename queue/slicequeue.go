package queue

// SliceQueue - slice based list of elements organized according to the FIFO principle
type SliceQueue struct {
	elements []string
}

// Len returns length of the queue
func (q *SliceQueue) Len() int {
	return len(q.elements)
}

// Enqueue adds element into queue
func (q *SliceQueue) Enqueue(element string) {
	q.elements = append(q.elements, element)
}

// Dequeue obtains and deletes element from queue
func (q *SliceQueue) Dequeue() (element string, ok bool) {
	if q.Len() == 0 {
		return "", false
	}

	element = q.elements[0]     // Obtain the first inserted element.
	q.elements = q.elements[1:] // Remove the first inserted element.
	return element, true
}

// Peek obtains element from queue
func (q *SliceQueue) Peek() (element string, ok bool) {
	if q.Len() == 0 {
		return "", false
	}

	return q.elements[0], true
}
