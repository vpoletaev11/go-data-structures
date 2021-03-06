package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/queue"
)

type queueI interface {
	Enqueue(string)
	Dequeue() (string, bool)
	Peek() (string, bool)
	Len() int
}

var dataset = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
}

var datasetPQ = []valuePriority{
	{"four", 80},
	{"two", 1000},
	{"ten", 1},
	{"three", 100},
	{"seven", 58},
	{"eight", 54},
	{"five", 75},
	{"one", 10000},
	{"six", 60},
	{"nine", 30},
}

// Testing SliceQueue and ListQueue:
func TestSliceQueueSuccess(t *testing.T) {
	queueSuccess(t, &queue.SliceQueue{})
}

func TestListQueueSuccess(t *testing.T) {
	queueSuccess(t, &queue.ListQueue{})
}

func queueSuccess(t *testing.T, queue queueI) {
	for _, value := range dataset {
		queue.Enqueue(value)
	}

	for i, expected := range dataset {
		assert.Equal(t, len(dataset)-i, queue.Len())

		val, ok := queue.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value: \"%s\"", expected)

		val, ok = queue.Dequeue()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Dequeue for expected value: \"%s\"", expected)
	}
}

func TestSliceQueueOutOfRange(t *testing.T) {
	queueOutOfRange(t, &queue.SliceQueue{})
}

func TestListQueueOutOfRange(t *testing.T) {
	queueOutOfRange(t, &queue.ListQueue{})
}

func queueOutOfRange(t *testing.T, queue queueI) {
	// Peek and Dequeue in empty queue
	val, ok := queue.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = queue.Dequeue()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing PriorityQueue:
func TestPriorityQueueSuccess(t *testing.T) {
	var p queue.PriorityQueue

	for _, valPrior := range datasetPQ {
		p.Enqueue(valPrior.value, valPrior.priority)
	}

	for i, expected := range dataset {
		assert.Equal(t, len(dataset)-i, p.Len())

		val, ok := p.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value: \"%s\"", expected)

		val, ok = p.Dequeue()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Dequeue for expected value: \"%s\"", expected)
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

// BENCHMARKS
// Benchmarking ListQueue:
func BenchmarkEnqueueListQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var l queue.ListQueue
		b.StartTimer()

		for _, value := range dataset {
			l.Enqueue(value)
		}
	}
}

func BenchmarkDequeueListQueue(b *testing.B) {
	var l queue.ListQueue
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			l.Enqueue(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			l.Dequeue()
		}
	}
}

func BenchmarkPeekListQueue(b *testing.B) {
	var l queue.ListQueue
	l.Enqueue("value")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Peek()
	}
}

// Benchmarking SliceQueue:
func BenchmarkEnqueueSliceQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var s queue.SliceQueue
		b.StartTimer()

		for _, value := range dataset {
			s.Enqueue(value)
		}
	}
}

func BenchmarkDequeueSliceQueue(b *testing.B) {
	var s queue.SliceQueue
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			s.Enqueue(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			s.Dequeue()
		}
	}
}

func BenchmarkPeekSliceQueue(b *testing.B) {
	var s queue.SliceQueue
	s.Enqueue("value")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Peek()
	}
}

// Benchmarking PriorityQueue:
func BenchmarkEnqueuePriorityQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var p queue.PriorityQueue
		b.StartTimer()

		for _, valPrior := range datasetPQ {
			p.Enqueue(valPrior.value, valPrior.priority)
		}
	}
}

func BenchmarkDequeuePriorityQueue(b *testing.B) {
	var p queue.PriorityQueue
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, valPrior := range datasetPQ {
			p.Enqueue(valPrior.value, valPrior.priority)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			p.Dequeue()
		}
	}
}

func BenchmarkPeekPriorityQueue(b *testing.B) {
	var p queue.PriorityQueue
	p.Enqueue("value", 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p.Peek()
	}
}

type valuePriority struct {
	value    string
	priority int
}
