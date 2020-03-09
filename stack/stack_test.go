package stack

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	var stack Stack

	stack.Push("one")
	stack.Push("two")
	stack.Push("three")

	if stack.Len() != 3 {
		t.Error("Pushed 3 elements, but length of stack ==", stack.Len())
	}
}

func TestPush(t *testing.T) {
	var stack Stack

	stack.Push("one")

	if stack.items[0] != "one" {
		t.Error("Pushed element \"one\", but stack item ==", stack.items[0])
	}
}

func TestPopSuccess(t *testing.T) {
	var stack Stack

	stack.items = append(stack.items, "one")
	stack.items = append(stack.items, "two")

	elem, ok := stack.Pop()
	if ok != true {
		t.Error("Pop status should be true, but returns false")
	}
	if elem != "two" {
		t.Error("Pushed value \"two\", but popped:", elem)
	}

	elem, ok = stack.Pop()
	if ok != true {
		t.Error("Pop status should be true, but returns false")
	}
	if elem != "one" {
		t.Error("Pushed value \"one\", but popped:", elem)
	}
}

func TestPopEmptyStack(t *testing.T) {
	var stack Stack

	elem, ok := stack.Pop()
	if elem != "" {
		t.Error("Nothing pushed, but popped value ==", elem)
	}
	if ok != false {
		t.Error("Nothing pushed, but pop status == true")
	}
}

func TestPeekSuccess(t *testing.T) {
	var stack Stack

	stack.items = append(stack.items, "one")

	elem, ok := stack.Peek()
	if ok != true {
		t.Error("Peek should return true, but returns false")
	}
	if elem != "one" {
		t.Error("Pushed value \"one\", but peeked:", elem)
	}
}

func TestPeekEmptyStack(t *testing.T) {
	var stack Stack

	elem, ok := stack.Peek()
	if elem != "" {
		t.Error("Nothing pushed, but peeked value ==", elem)
	}
	if ok != false {
		t.Error("Nothing pushed, but peek status == true")
	}
}

// TestStackDemo displays example of using the stack data structure
func TestStackDemo(t *testing.T) {
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

	for stack.Len() > 0 {
		elem, ok := stack.Pop()
		if ok {
			fmt.Println("Popped:", elem)
		}
	}
}
