package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/queue"
)

// Testing slice based queue:
func TestSliceQueue(t *testing.T) {
	var s queue.SliceQueue

	// Enqueue
	s.Enqueue("one")
	s.Enqueue("two")
	s.Enqueue("three")

	// Len
	assert.Equal(t, 3, s.Len())

	// Success Peek and Dequeue
	val, ok := s.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = s.Dequeue()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = s.Peek()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.Dequeue()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.Peek()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.Dequeue()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	// Peek and Dequeue in empty queue
	val, ok = s.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Dequeue()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing singly linked list based queue:
func TestListQueue(t *testing.T) {
	var l queue.ListQueue

	// Enqueue
	l.Enqueue("one")
	l.Enqueue("two")
	l.Enqueue("three")

	// Len
	assert.Equal(t, 3, l.Len())

	// Success Peek and Dequeue
	val, ok := l.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.Dequeue()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.Peek()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Dequeue()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Peek()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.Dequeue()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	// Peek and Dequeue in empty queue
	val, ok = l.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Dequeue()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing priority queue:
func TestPriorityQueue(t *testing.T) {
	var p queue.PriorityQueue

	// Push
	p.Push("one", 1)
	p.Push("two", 2)
	p.Push("three", 3)
	p.Push("four", 4)
	p.Push("five", 5521)
	p.Push("six", 111)
	p.Push("seven", 99)
	p.Push("eight", 13)

	// Len
	assert.Equal(t, 8, p.Len())

	// Success Peek and Pop
	val, ok := p.Peek()
	assert.Equal(t, "five", val)
	assert.True(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "five", val)
	assert.True(t, ok)

	val, ok = p.Peek()
	assert.Equal(t, "six", val)
	assert.True(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "six", val)
	assert.True(t, ok)

	val, ok = p.Peek()
	assert.Equal(t, "seven", val)
	assert.True(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "seven", val)
	assert.True(t, ok)

	val, ok = p.Peek()
	assert.Equal(t, "eight", val)
	assert.True(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "eight", val)
	assert.True(t, ok)

	val, ok = p.Peek()
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = p.Peek()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = p.Peek()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = p.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// Peek and Pop in empty queue
	val, ok = p.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = p.Pop()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}
