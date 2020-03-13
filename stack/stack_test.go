package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	var stack Stack

	stack.Push("one")
	stack.Push("two")
	stack.Push("three")

	assert.Equal(t, 3, stack.Len())
}

func TestPush(t *testing.T) {
	var stack Stack

	stack.Push("one")

	assert.Equal(t, "one", stack.elements[0])
}

func TestPopSuccess(t *testing.T) {
	var stack Stack

	stack.Push("one")
	stack.Push("two")

	elem, ok := stack.Pop()
	assert.Equal(t, "two", elem)
	assert.True(t, ok)

	elem, ok = stack.Pop()
	assert.Equal(t, "one", elem)
	assert.True(t, ok)
}

func TestPopEmptyStack(t *testing.T) {
	var stack Stack

	elem, ok := stack.Pop()

	assert.Equal(t, "", elem)
	assert.False(t, ok)
}

func TestPeekSuccess(t *testing.T) {
	var stack Stack

	stack.Push("one")
	elem, ok := stack.Peek()

	assert.Equal(t, "one", elem)
	assert.True(t, ok)
}

func TestPeekEmptyStack(t *testing.T) {
	var stack Stack

	elem, ok := stack.Peek()

	assert.Equal(t, "", elem)
	assert.False(t, ok)
}
