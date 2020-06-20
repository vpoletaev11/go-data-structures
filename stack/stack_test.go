package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/stack"
)

type stackI interface {
	Push(string)
	Pop() (string, bool)
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

var expectedOut = []string{
	"ten",
	"nine",
	"eight",
	"seven",
	"six",
	"five",
	"four",
	"three",
	"two",
	"one",
}

func TestSliceStackSuccess(t *testing.T) {
	stackSuccess(t, &stack.SliceStack{})
}

func TestListStackSuccess(t *testing.T) {
	stackSuccess(t, &stack.ListStack{})
}

func stackSuccess(t *testing.T, stack stackI) {
	for _, value := range dataset {
		stack.Push(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, stack.Len())

		val, ok := stack.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value: \"%s\"", expected)

		val, ok = stack.Pop()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Pop for expected value: \"%s\"", expected)
	}
}

func TestSliceStackOutOfRange(t *testing.T) {
	stackOutOfRange(t, &stack.SliceStack{})
}

func TestListStackOutOfRange(t *testing.T) {
	stackOutOfRange(t, &stack.ListStack{})
}

func stackOutOfRange(t *testing.T, stack stackI) {
	// Peek and Pop in empty stack
	val, ok := stack.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = stack.Pop()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BENCHMARKS
// Benchmarking ListStack:
func BenchmarkPushListStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var l stack.ListStack
		b.StartTimer()

		for _, value := range dataset {
			l.Push(value)
		}
	}
}

func BenchmarkPopListStack(b *testing.B) {
	var l stack.ListStack
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			l.Push(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			l.Pop()
		}
	}
}

func BenchmarkPeekListStack(b *testing.B) {
	var l stack.ListStack
	l.Push("value")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Peek()
	}
}

// Benchmarking SliceStack:
func BenchmarkPushSliceStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var s stack.SliceStack
		b.StartTimer()

		for _, value := range dataset {
			s.Push(value)
		}
	}
}

func BenchmarkPopSliceStack(b *testing.B) {
	var s stack.SliceStack
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			s.Push(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			s.Pop()
		}
	}
}

func BenchmarkPeekSliceStack(b *testing.B) {
	var s stack.SliceStack
	s.Push("value")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Peek()
	}
}
