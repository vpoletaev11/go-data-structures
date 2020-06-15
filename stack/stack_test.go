package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/stack"
)

// Testing SliceStack
func TestSliceStackSuccess(t *testing.T) {
	var s stack.SliceStack

	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	expectedOut := []string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"}

	for _, value := range dataset {
		s.Push(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, s.Len())

		val, ok := s.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value: \"%s\" %s", expected, "should return true, but returns false")

		val, ok = s.Pop()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Pop for expected value: \"%s\" %s", expected, "should return true, but returns false")
	}
}

func TestSliceStackOutOfRange(t *testing.T) {
	var s stack.SliceStack

	// Peek and Pop in empty stack
	val, ok := s.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Pop()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing ListStack
func TestListStackSuccess(t *testing.T) {
	var l stack.ListStack

	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	expectedOut := []string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"}

	for _, value := range dataset {
		l.Push(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, l.Len())

		val, ok := l.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value: \"%s\" %s", expected, "should return true, but returns false")

		val, ok = l.Pop()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Pop for expected value: \"%s\" %s", expected, "should return true, but returns false")
	}
}

func TestListStackOutOfRange(t *testing.T) {
	var l stack.ListStack

	// Peek and Pop in empty stack
	val, ok := l.Peek()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Pop()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BENCHMARKS
// Benchmarking ListStack:
func BenchmarkPushListStack(b *testing.B) {
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

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
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
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
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

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
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
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
