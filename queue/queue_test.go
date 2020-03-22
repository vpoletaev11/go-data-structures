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

// Testing priority queue:
//
// test Len
func TestLenPriorityQueue(t *testing.T) {
	var q queue.PriorityQueue

	q.Push("one", 1)
	q.Push("two", 2)
	q.Push("three", 3)

	assert.Equal(t, 3, q.Len())
}

// test Push & Pop
func TestPushAndPopPriorityQueue(t *testing.T) {
	var queue queue.PriorityQueue

	queue.Push("one", 1)
	queue.Push("two", 2)
	queue.Push("three", 3)
	queue.Push("four", 4)
	queue.Push("five", 5521)
	queue.Push("six", 111)
	queue.Push("seven", 99)
	queue.Push("eight", 13)

	val, ok := queue.Pop()
	assert.Equal(t, "five", val)
	assert.True(t, ok)

	val, ok = queue.Pop()
	assert.Equal(t, "six", val)
	assert.True(t, ok)

	val, ok = queue.Pop()
	assert.Equal(t, "seven", val)
	assert.True(t, ok)

	val, ok = queue.Pop()
	assert.Equal(t, "eight", val)
	assert.True(t, ok)

	val, ok = queue.Pop()
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = queue.Pop()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = queue.Pop()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = queue.Pop()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = queue.Pop()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// test Peek
func TestPeekPriorityQueue(t *testing.T) {
	var q queue.PriorityQueue

	q.Push("one", 1)

	val, ok := q.Peek()
	assert.Equal(t, "one", val)
	assert.True(t, ok)
}

func TestPeekEmptyQueuePriorityQueue(t *testing.T) {
	var q queue.PriorityQueue

	val, ok := q.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}
