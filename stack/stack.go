package stack

import "fmt"

// Stack - list of elements organized according to the LIFO principle
type Stack []string

// Push adds element into stack
func (s *Stack) Push(val string) {
	*s = append(*s, val)
}

// Pop obtains and deletes element from stack
func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}

	index := len(*s) - 1   // Get the index of the last inserted element.
	element := (*s)[index] // Obtain the last inserted element.
	*s = (*s)[:index]      // Remove the last inserted element.
	return element, true
}

// Peek obtains element from stack
func (s *Stack) Peek() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}

	return (*s)[len(*s)-1], true
}

// Demo displays example of using the stack data structure
func Demo() {
	fmt.Println("Demonstration of using the stack data structure:")
	var stack Stack

	stack.Push("one")
	fmt.Println("Pushed: one")

	stack.Push("two")
	fmt.Println("Pushed: two")

	stack.Push("three")
	fmt.Println("Pushed: three")
	fmt.Print("\n")

	elem, ok := stack.Peek()
	if ok {
		fmt.Println("Peeked:", elem)
	}

	for len(stack) > 0 {
		elem, ok := stack.Pop()
		if ok {
			fmt.Println("Popped:", elem)
		}
	}
}
