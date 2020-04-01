package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/stack"
)

// Testing slice based stack:
func TestSliceStack(t *testing.T) {
	var s stack.SliceStack

	// Push
	s.Push("one")
	s.Push("two")
	s.Push("three")

	// Len
	assert.Equal(t, 3, s.Len())

	// Success Peek and Pop
	val, ok := s.Peek()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.Pop()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.Peek()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.Pop()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = s.Pop()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// Peek and Pop in empty stack
	val, ok = s.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Pop()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing singly linked list based stack:
func TestListStack(t *testing.T) {
	var l stack.ListStack

	// Push
	l.Push("one")
	l.Push("two")
	l.Push("three")

	// Len
	assert.Equal(t, 3, l.Len())

	// Success Peek and Pop
	val, ok := l.Peek()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.Pop()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.Peek()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Pop()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.Pop()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// Peek and Pop in empty stack
	val, ok = l.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Pop()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}
