package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/stack"
)

func TestLen(t *testing.T) {
	var s stack.SliceStack

	s.Push("one")
	s.Push("two")
	s.Push("three")

	assert.Equal(t, 3, s.Len())
}

func TestPush(t *testing.T) {
	var s stack.SliceStack

	s.Push("one")

	actual, ok := s.Peek()
	assert.Equal(t, "one", actual)
	assert.True(t, ok)
}

func TestPopSuccess(t *testing.T) {
	var s stack.SliceStack

	s.Push("one")
	s.Push("two")

	elem, ok := s.Pop()
	assert.Equal(t, "two", elem)
	assert.True(t, ok)

	elem, ok = s.Pop()
	assert.Equal(t, "one", elem)
	assert.True(t, ok)
}

func TestPopEmptyStack(t *testing.T) {
	var s stack.SliceStack

	elem, ok := s.Pop()

	assert.Equal(t, "", elem)
	assert.False(t, ok)
}

func TestPeekSuccess(t *testing.T) {
	var s stack.SliceStack

	s.Push("one")
	elem, ok := s.Peek()

	assert.Equal(t, "one", elem)
	assert.True(t, ok)
}

func TestPeekEmptyStack(t *testing.T) {
	var s stack.SliceStack

	elem, ok := s.Peek()

	assert.Equal(t, "", elem)
	assert.False(t, ok)
}
