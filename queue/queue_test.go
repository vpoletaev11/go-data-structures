package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/queue"
)

// Testing slice based queue:
//
// test Len
func TestLenSliceQueue(t *testing.T) {
	var q queue.SliceQueue

	q.Enqueue("one")
	q.Enqueue("two")
	q.Enqueue("three")

	assert.Equal(t, 3, q.Len())
}

// test Enqueue
func TestEnqueueSliceQueue(t *testing.T) {
	var q queue.SliceQueue

	q.Enqueue("one")

	element, ok := q.Peek()
	assert.Equal(t, "one", element)
	assert.True(t, ok)
}

// test Dequeue
func TestDequeueSliceQueueSuccess(t *testing.T) {
	var q queue.SliceQueue

	q.Enqueue("one")
	q.Enqueue("two")

	element, ok := q.Dequeue()
	assert.Equal(t, "one", element)
	assert.True(t, ok)

	element, ok = q.Dequeue()
	assert.Equal(t, "two", element)
	assert.True(t, ok)
}

func TestDequeueSliceQueueEmptyQueue(t *testing.T) {
	var q queue.SliceQueue

	element, ok := q.Dequeue()
	assert.Equal(t, "", element)
	assert.False(t, ok)
}

// test Peek
func TestPeekSliceQueueSuccess(t *testing.T) {
	var q queue.SliceQueue

	q.Enqueue("one")

	element, ok := q.Peek()
	assert.Equal(t, "one", element)
	assert.True(t, ok)
}

func TestPeekEmptyQueueSliceQueue(t *testing.T) {
	var q queue.SliceQueue

	element, ok := q.Peek()
	assert.Equal(t, "", element)
	assert.False(t, ok)
}

// Testing singly linked list based queue:
//
// test Len
func TestLenListQueue(t *testing.T) {
	var q queue.ListQueue

	q.Enqueue("one")
	q.Enqueue("two")
	q.Enqueue("three")

	assert.Equal(t, 3, q.Len())
}

// test Enqueue
func TestEnqueueListQueue(t *testing.T) {
	var q queue.ListQueue

	q.Enqueue("one")

	element, ok := q.Peek()
	assert.Equal(t, "one", element)
	assert.True(t, ok)
}

// test Dequeue
func TestDequeueListQueueSuccess(t *testing.T) {
	var q queue.ListQueue

	q.Enqueue("one")
	q.Enqueue("two")

	element, ok := q.Dequeue()
	assert.Equal(t, "one", element)
	assert.True(t, ok)

	element, ok = q.Dequeue()
	assert.Equal(t, "two", element)
	assert.True(t, ok)
}

func TestDequeueListQueueEmptyQueue(t *testing.T) {
	var q queue.ListQueue

	element, ok := q.Dequeue()
	assert.Equal(t, "", element)
	assert.False(t, ok)
}

// test Peek
func TestPeekListQueueSuccess(t *testing.T) {
	var q queue.ListQueue

	q.Enqueue("one")

	element, ok := q.Peek()
	assert.Equal(t, "one", element)
	assert.True(t, ok)
}

func TestPeekEmptyQueueListQueue(t *testing.T) {
	var q queue.ListQueue

	element, ok := q.Peek()
	assert.Equal(t, "", element)
	assert.False(t, ok)
}
