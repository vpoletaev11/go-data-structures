package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	var queue Queue

	queue.Enqueue("one")
	queue.Enqueue("two")
	queue.Enqueue("three")

	assert.Equal(t, 3, queue.Len())
}

func TestEnqueue(t *testing.T) {
	var queue Queue

	queue.Enqueue("one")

	element, ok := queue.Peek()
	assert.Equal(t, "one", element)
	assert.True(t, ok)
}

func TestDequeueSuccess(t *testing.T) {
	var queue Queue

	queue.Enqueue("one")
	queue.Enqueue("two")

	element, ok := queue.Dequeue()
	assert.Equal(t, "one", element)
	assert.True(t, ok)

	element, ok = queue.Dequeue()
	assert.Equal(t, "two", element)
	assert.True(t, ok)
}

func TestDequeueEmptyQueue(t *testing.T) {
	var queue Queue

	element, ok := queue.Dequeue()
	assert.Equal(t, "", element)
	assert.False(t, ok)
}

func TestPeekSuccess(t *testing.T) {
	var queue Queue

	queue.Enqueue("one")

	element, ok := queue.Peek()
	assert.Equal(t, "one", element)
	assert.True(t, ok)
}

func TestPeekEmptyQueue(t *testing.T) {
	var queue Queue

	element, ok := queue.Peek()
	assert.Equal(t, "", element)
	assert.False(t, ok)
}
