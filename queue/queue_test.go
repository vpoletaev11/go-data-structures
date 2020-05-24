package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/queue"
)

// Testing SliceQueue:
func TestSliceQueueSuccess(t *testing.T) {
	var s queue.SliceQueue

	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	for _, value := range dataset {
		s.Enqueue(value)
	}

	for i, expected := range dataset {
		assert.Equal(t, len(dataset)-i, s.Len())

		val, ok := s.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value:", expected, "should return true, but returns false")

		val, ok = s.Dequeue()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Dequeue for expected value:", expected, "should return true, but returns false")
	}
}

func TestSliceQueueOutOfRange(t *testing.T) {
	var s queue.SliceQueue

	// Peek and Dequeue in empty queue
	val, ok := s.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Dequeue()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing ListQueue:
func TestListQueueSuccess(t *testing.T) {
	var l queue.ListQueue

	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	for _, value := range dataset {
		l.Enqueue(value)
	}

	for i, expected := range dataset {
		assert.Equal(t, len(dataset)-i, l.Len())

		val, ok := l.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value:", expected, "should return true, but returns false")

		val, ok = l.Dequeue()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Dequeue for expected value:", expected, "should return true, but returns false")
	}
}

func TestListQueueOutOfRange(t *testing.T) {
	var l queue.ListQueue

	// Peek and Dequeue in empty queue
	val, ok := l.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Dequeue()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing PriorityQueue:
func TestPriorityQueueSuccess(t *testing.T) {
	var p queue.PriorityQueue

	dataset := []valuePriority{{"four", 80}, {"two", 1000}, {"ten", 1}, {"three", 100}, {"seven", 58}, {"eight", 54}, {"five", 75}, {"one", 10000}, {"six", 60}, {"nine", 30}}
	expectedOut := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	for _, valPrior := range dataset {
		p.Enqueue(valPrior.value, valPrior.priority)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, p.Len())

		val, ok := p.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value:", expected, "should return true, but returns false")

		val, ok = p.Dequeue()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Dequeue for expected value:", expected, "should return true, but returns false")
	}
}

func TestPriorityQueueOutOfRange(t *testing.T) {
	var p queue.PriorityQueue

	// Peek and Pop in empty queue
	val, ok := p.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = p.Dequeue()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

type valuePriority struct {
	value    string
	priority int
}
