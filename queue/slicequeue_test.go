package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/queue"
)

func TestLen(t *testing.T) {
	var q queue.SliceQueue

	q.Enqueue("one")
	q.Enqueue("two")
	q.Enqueue("three")

	assert.Equal(t, 3, q.Len())
}

func TestEnqueue(t *testing.T) {
	var q queue.SliceQueue

	q.Enqueue("one")

	element, ok := q.Peek()
	assert.Equal(t, "one", element)
	assert.True(t, ok)
}

func TestDequeueSuccess(t *testing.T) {
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

func TestDequeueEmptyQueue(t *testing.T) {
	var q queue.SliceQueue

	element, ok := q.Dequeue()
	assert.Equal(t, "", element)
	assert.False(t, ok)
}

func TestPeekSuccess(t *testing.T) {
	var q queue.SliceQueue

	q.Enqueue("one")

	element, ok := q.Peek()
	assert.Equal(t, "one", element)
	assert.True(t, ok)
}

func TestPeekEmptyQueue(t *testing.T) {
	var q queue.SliceQueue

	element, ok := q.Peek()
	assert.Equal(t, "", element)
	assert.False(t, ok)
}
