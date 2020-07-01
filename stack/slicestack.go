package stack

/*
	SLICE STACK:

	Len()                Push()               Pop()                Peek()
	Best:    O(1)        Best:    O(1)        Best:    O(1)        Best:    O(1)
	Average: O(1)        Average: O(1)        Average: O(1)        Average: O(1)
	Worst:   O(1)        Worst:   O(n)        Worst:   O(1)        Worst:   O(1)


	goos: linux
	goarch: amd64
	BenchmarkPushSliceStack-8   	 1271370	       935 ns/op	     496 B/op	       5 allocs/op
	BenchmarkPopSliceStack-8    	18954626	        64.8 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPeekSliceStack-8   	1000000000	         0.293 ns/op	       0 B/op	       0 allocs/op
*/


// SliceStack - slice based list of elements organized according to the LIFO principle
type SliceStack struct {
	elements []string
}

// Len returns length of the stack
func (s *SliceStack) Len() int {
	return len(s.elements)
}

// Push adds element into stack
func (s *SliceStack) Push(val string) {
	s.elements = append(s.elements, val)
}

// Pop obtains and deletes element from stack
func (s *SliceStack) Pop() (val string, status bool) {
	if s.Len() == 0 {
		return "", false
	}

	index := s.Len() - 1            // Get the index of the last inserted element.
	val = s.elements[index]         // Obtain the last inserted element.
	s.elements = s.elements[:index] // Remove the last inserted element.
	return val, true
}

// Peek obtains element from stack
func (s *SliceStack) Peek() (val string, status bool) {
	if s.Len() == 0 {
		return "", false
	}

	return s.elements[s.Len()-1], true
}
