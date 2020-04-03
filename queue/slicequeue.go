package queue

// SliceQueue - slice based list of elements organized according to the FIFO principle
type SliceQueue struct {
	elements []string
}

// Len returns length of the queue
func (s *SliceQueue) Len() int {
	return len(s.elements)
}

// Enqueue adds element into queue
func (s *SliceQueue) Enqueue(val string) {
	s.elements = append(s.elements, val)
}

// Dequeue obtains and deletes element from queue
func (s *SliceQueue) Dequeue() (val string, ok bool) {
	if s.Len() == 0 {
		return "", false
	}

	val = s.elements[0]     // Obtain the first inserted element.
	s.elements = s.elements[1:] // Remove the first inserted element.
	return val, true
}

// Peek obtains element from queue
func (s *SliceQueue) Peek() (val string, ok bool) {
	if s.Len() == 0 {
		return "", false
	}

	return s.elements[0], true
}
