package stack

// Stack - list of items organized according to the LIFO principle
type Stack struct {
	items []string
}

// Len returns length of the stack
func (s *Stack) Len() int {
	return len(s.items)
}

// Push adds element into stack
func (s *Stack) Push(val string) {
	s.items = append(s.items, val)
}

// Pop obtains and deletes element from stack
func (s *Stack) Pop() (string, bool) {
	if s.Len() == 0 {
		return "", false
	}

	index := s.Len() - 1      // Get the index of the last inserted element.
	element := s.items[index] // Obtain the last inserted element.
	s.items = s.items[:index] // Remove the last inserted element.
	return element, true
}

// Peek obtains element from stack
func (s *Stack) Peek() (string, bool) {
	if s.Len() == 0 {
		return "", false
	}

	return s.items[s.Len()-1], true
}
