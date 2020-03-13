package stack

// Stack - list of elements organized according to the LIFO principle
type Stack struct {
	elements []string
}

// Len returns length of the stack
func (s *Stack) Len() int {
	return len(s.elements)
}

// Push adds element into stack
func (s *Stack) Push(element string) {
	s.elements = append(s.elements, element)
}

// Pop obtains and deletes element from stack
func (s *Stack) Pop() (element string, status bool) {
	if s.Len() == 0 {
		return "", false
	}

	index := s.Len() - 1            // Get the index of the last inserted element.
	element = s.elements[index]     // Obtain the last inserted element.
	s.elements = s.elements[:index] // Remove the last inserted element.
	return element, true
}

// Peek obtains element from stack
func (s *Stack) Peek() (element string, status bool) {
	if s.Len() == 0 {
		return "", false
	}

	return s.elements[s.Len()-1], true
}
