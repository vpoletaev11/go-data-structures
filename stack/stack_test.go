package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/stack"
)

// Testing slice based stack:
//
// test Len
func TestLenSliceStack(t *testing.T) {
	var s stack.SliceStack

	s.Push("one")
	s.Push("two")
	s.Push("three")

	assert.Equal(t, 3, s.Len())
}

// test Push
func TestPushSliceStack(t *testing.T) {
	var s stack.SliceStack

	s.Push("one")

	actual, ok := s.Peek()
	assert.Equal(t, "one", actual)
	assert.True(t, ok)
}

// test Pop
func TestPopSliceStackSuccess(t *testing.T) {
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

func TestPopEmptyStackSliceStack(t *testing.T) {
	var s stack.SliceStack

	elem, ok := s.Pop()

	assert.Equal(t, "", elem)
	assert.False(t, ok)
}

// test Peek
func TestPeekSliceStackSuccess(t *testing.T) {
	var s stack.SliceStack

	s.Push("one")
	elem, ok := s.Peek()

	assert.Equal(t, "one", elem)
	assert.True(t, ok)
}

func TestPeekEmptyStackSliceStack(t *testing.T) {
	var s stack.SliceStack

	elem, ok := s.Peek()

	assert.Equal(t, "", elem)
	assert.False(t, ok)
}

// Testing list based stack:
//
// test Len
func TestLenListStack(t *testing.T) {
	var s stack.ListStack

	s.Push("one")
	s.Push("two")
	s.Push("three")

	assert.Equal(t, 3, s.Len())
}

// test Push
func TestPushListStack(t *testing.T) {
	var s stack.ListStack

	s.Push("one")

	actual, ok := s.Peek()
	assert.Equal(t, "one", actual)
	assert.True(t, ok)
}

// test Pop
func TestPopListStackSuccess(t *testing.T) {
	var s stack.ListStack

	s.Push("one")
	s.Push("two")

	elem, ok := s.Pop()
	assert.Equal(t, "two", elem)
	assert.True(t, ok)

	elem, ok = s.Pop()
	assert.Equal(t, "one", elem)
	assert.True(t, ok)
}

func TestPopEmptyStackListStack(t *testing.T) {
	var s stack.ListStack

	elem, ok := s.Pop()

	assert.Equal(t, "", elem)
	assert.False(t, ok)
}

// test Peek
func TestPeekListStackSuccess(t *testing.T) {
	var s stack.ListStack

	s.Push("one")
	elem, ok := s.Peek()

	assert.Equal(t, "one", elem)
	assert.True(t, ok)
}

func TestPeekEmptyStackListStack(t *testing.T) {
	var s stack.ListStack

	elem, ok := s.Peek()

	assert.Equal(t, "", elem)
	assert.False(t, ok)
}
