package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/queue"
)

// Testing slice based queue:
func TestSliceQueue(t *testing.T) {
	var q queue.SliceQueue

	// Enqueue
	q.Enqueue("one")
	q.Enqueue("two")
	q.Enqueue("three")

	// Len
	assert.Equal(t, 3, q.Len())

	// Success Peek and Dequeue
	val, ok := q.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = q.Dequeue()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = q.Dequeue()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = q.Dequeue()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	// Peek and Dequeue in empty queue
	val, ok = q.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = q.Dequeue()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing singly linked list based queue:
func TestListQueue(t *testing.T) {
	var q queue.ListQueue

	// Enqueue
	q.Enqueue("one")
	q.Enqueue("two")
	q.Enqueue("three")

	// Len
	assert.Equal(t, 3, q.Len())

	// Success Peek and Dequeue
	val, ok := q.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = q.Dequeue()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = q.Dequeue()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = q.Dequeue()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	// Peek and Dequeue in empty queue
	val, ok = q.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = q.Dequeue()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing priority queue:
func TestPriorityQueue(t *testing.T) {
	var q queue.PriorityQueue

	// Push
	q.Push("one", 1)
	q.Push("two", 2)
	q.Push("three", 3)
	q.Push("four", 4)
	q.Push("five", 5521)
	q.Push("six", 111)
	q.Push("seven", 99)
	q.Push("eight", 13)

	// Len
	assert.Equal(t, 8, q.Len())

	// Success Peek and Pop
	val, ok := q.Peek()
	assert.Equal(t, "five", val)
	assert.True(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "five", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "six", val)
	assert.True(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "six", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "seven", val)
	assert.True(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "seven", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "eight", val)
	assert.True(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "eight", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = q.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// Peek and Pop in empty queue
	val, ok = q.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = q.Pop()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}
