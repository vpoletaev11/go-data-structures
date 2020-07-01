package queue

/*
	SLICE QUEUE:

	Len()                Enqueue()            Dequeue()            Peek()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(1)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(1)
	Worst:   O(1)        Worst:   O(n)        Worst:   O(1)        Worst:   O(1)

	
	goos: linux
	goarch: amd64
	BenchmarkEnqueueSliceQueue-8      	 1269865	       956 ns/op	     496 B/op	       5 allocs/op
	BenchmarkDequeueSliceQueue-8      	14281522	        84.0 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPeekSliceQueue-8         	1000000000	         0.288 ns/op	       0 B/op	       0 allocs/op
*/

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

	val = s.elements[0]         // Obtain the first inserted element.
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
