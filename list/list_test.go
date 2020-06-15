package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/list"
)

// Testing SinglyLinkedList:
// HEAD
func TestSinglyListHEADSuccess(t *testing.T) {
	var s list.SinglyLinkedList

	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	expectedOut := []string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"}

	for _, value := range dataset {
		s.InsertHead(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, s.Len())

		val, ok := s.PeekHead()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "PeekHead for expected value: \"%s\" %s", expected, "should return true, but returns false")

		val, ok = s.GetHead()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "GetHead for expected value: \"%s\" %s", expected, "should return true, but returns false")
	}
}

func TestSinglyListHEADOutOfRange(t *testing.T) {
	var s list.SinglyLinkedList

	// PeekHead and GetHead in empty list
	val, ok := s.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BODY
func TestSinglyListBODYSuccess(t *testing.T) {
	var s list.SinglyLinkedList

	ok := s.Insert(0, "three")
	assert.True(t, ok)

	ok = s.Insert(0, "one")
	assert.True(t, ok)

	ok = s.Insert(1, "two")
	assert.True(t, ok)

	ok = s.Insert(3, "four")
	assert.True(t, ok)

	val, ok := s.Peek(3)
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = s.Get(3)
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = s.Peek(2)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.Get(2)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.Peek(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.Get(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.Peek(0)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = s.Get(0)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

}

func TestSinglyListBODYOutOfRange(t *testing.T) {
	var s list.SinglyLinkedList

	// Insert out of range
	ok := s.Insert(5, "one")
	assert.False(t, ok)

	ok = s.Insert(-1, "two")
	assert.False(t, ok)

	// Peek and Get in empty list
	val, ok := s.Peek(10)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Get(10)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	// for following tests list should be not empty
	s.InsertHead("stub")

	// Peek and Get out of range
	val, ok = s.Peek(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Peek(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Get(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Get(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// TAIL
func TestSinglyListTailSuccess(t *testing.T) {
	var s list.SinglyLinkedList

	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	expectedOut := []string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"}

	for _, value := range dataset {
		s.InsertTail(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, s.Len())

		val, ok := s.PeekTail()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "PeekTail for expected value: \"%s\" %s", expected, "should return true, but returns false")

		val, ok = s.GetTail()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "GetTail for expected value: \"%s\" %s", expected, "should return true, but returns false")
	}
}

func TestSinglyListTailOutOfRange(t *testing.T) {
	var s list.SinglyLinkedList

	// PeekTail and GetTail in empty list
	val, ok := s.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing DoublyLinkedList:
// HEAD
func TestDoublyListHEADSuccess(t *testing.T) {
	var d list.DoublyLinkedList

	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	expectedOut := []string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"}

	for _, value := range dataset {
		d.InsertHead(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, d.Len())

		val, ok := d.PeekHead()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "PeekHead for expected value: \"%s\" %s", expected, "should return true, but returns false")

		val, ok = d.GetHead()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "GetHead for expected value: \"%s\" %s", expected, "should return true, but returns false")
	}
}

func TestDoublyListHEADOutOfRange(t *testing.T) {
	var d list.DoublyLinkedList

	// PeekHead and GetHead in empty list
	val, ok := d.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BODY
func TestDoublyListBODYSuccess(t *testing.T) {
	var d list.DoublyLinkedList

	ok := d.Insert(0, "three")
	assert.True(t, ok)

	ok = d.Insert(0, "one")
	assert.True(t, ok)

	ok = d.Insert(1, "two")
	assert.True(t, ok)

	ok = d.Insert(3, "four")
	assert.True(t, ok)

	val, ok := d.Peek(3)
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = d.Get(3)
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = d.Peek(2)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = d.Get(2)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = d.Peek(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = d.Get(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = d.Peek(0)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = d.Get(0)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

}

func TestDoublyListBODYOutOfRange(t *testing.T) {
	var d list.DoublyLinkedList

	// Insert out of range
	ok := d.Insert(5, "one")
	assert.False(t, ok)

	ok = d.Insert(-1, "two")
	assert.False(t, ok)

	// Peek and Get in empty list
	val, ok := d.Peek(10)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.Get(10)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	// for following tests list should be not empty
	d.InsertHead("stub")

	// Peek and Get out of range
	val, ok = d.Peek(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.Peek(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.Get(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.Get(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// TAIL
func TestDoublyListTAILSuccess(t *testing.T) {
	var d list.DoublyLinkedList

	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	expectedOut := []string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"}

	for _, value := range dataset {
		d.InsertTail(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, d.Len())

		val, ok := d.PeekTail()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "PeekTail for expected value: \"%s\" %s", expected, "should return true, but returns false")

		val, ok = d.GetTail()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "GetTail for expected value: \"%s\" %s", expected, "should return true, but returns false")
	}
}

func TestDoublyListTAILOutOfRange(t *testing.T) {
	var d list.DoublyLinkedList

	// PeekTail and GetTail in empty list
	val, ok := d.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BENCHMARKS
// Benchmarking SinglyLinkedList:
func BenchmarkInsertHeadSinglyList(b *testing.B) {
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var s list.SinglyLinkedList
		b.StartTimer()

		for _, value := range dataset {
			s.InsertHead(value)
		}
	}
}

func BenchmarkInsertSinglyList(b *testing.B) {
	dataset := []valueIndex{{"two", 0}, {"one", 0}, {"three", 2}, {"four", 3}, {"six", 4}, {"five", 4}, {"ten", 6}, {"nine", 6}, {"eight", 6}, {"seven", 6}}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var s list.SinglyLinkedList
		b.StartTimer()

		for _, valPrior := range dataset {
			s.Insert(valPrior.index, valPrior.value)
		}
	}
}

func BenchmarkInsertTailSinglyList(b *testing.B) {
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var s list.SinglyLinkedList
		b.StartTimer()

		for _, value := range dataset {
			s.InsertTail(value)
		}
	}
}

func BenchmarkGetHeadSinglyList(b *testing.B) {
	var s list.SinglyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			s.InsertHead(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			s.GetHead()
		}
	}
}

func BenchmarkGetSinglyList(b *testing.B) {
	var s list.SinglyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	indexset := []int{6, 6, 6, 6, 4, 4, 3, 2, 0, 0}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			s.InsertHead(value)
		}
		b.StartTimer()

		for _, index := range indexset {
			s.Get(index)
		}
	}
}

func BenchmarkGetTailSinglyList(b *testing.B) {
	var s list.SinglyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			s.InsertHead(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			s.GetTail()
		}
	}
}

func BenchmarkPeekHeadSinglyList(b *testing.B) {
	var s list.SinglyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	s.InsertHead("value")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(dataset); i++ {
			s.PeekHead()
		}
	}
}

func BenchmarkPeekSinglyList(b *testing.B) {
	var s list.SinglyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	indexset := []int{6, 6, 6, 6, 4, 4, 3, 2, 0, 0}

	for _, value := range dataset {
		s.InsertHead(value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, index := range indexset {
			s.Peek(index)
		}
	}
}

func BenchmarkPeekTailSinglyList(b *testing.B) {
	var s list.SinglyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	for _, value := range dataset {
		s.InsertHead(value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(dataset); i++ {
			s.PeekTail()
		}
	}
}

// Benchmarking DoublyLinkedList:
func BenchmarkInsertHeadDoublyList(b *testing.B) {
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var d list.DoublyLinkedList
		b.StartTimer()

		for _, value := range dataset {
			d.InsertHead(value)
		}
	}
}

func BenchmarkInsertDoublyList(b *testing.B) {
	dataset := []valueIndex{{"two", 0}, {"one", 0}, {"three", 2}, {"four", 3}, {"six", 4}, {"five", 4}, {"ten", 6}, {"nine", 6}, {"eight", 6}, {"seven", 6}}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var d list.DoublyLinkedList
		b.StartTimer()

		for _, valPrior := range dataset {
			d.Insert(valPrior.index, valPrior.value)
		}
	}
}

func BenchmarkInsertTailDoublyList(b *testing.B) {
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var d list.DoublyLinkedList
		b.StartTimer()

		for _, value := range dataset {
			d.InsertTail(value)
		}
	}
}

func BenchmarkGetHeadDoublyList(b *testing.B) {
	var d list.DoublyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			d.InsertHead(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			d.GetHead()
		}
	}
}

func BenchmarkGetDoublyList(b *testing.B) {
	var d list.DoublyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	indexset := []int{6, 6, 6, 6, 4, 4, 3, 2, 0, 0}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			d.InsertHead(value)
		}
		b.StartTimer()

		for _, index := range indexset {
			d.Get(index)
		}
	}
}

func BenchmarkGetTailDoublyList(b *testing.B) {
	var d list.DoublyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			d.InsertHead(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			d.GetTail()
		}
	}
}

func BenchmarkPeekHeadDoublyList(b *testing.B) {
	var d list.DoublyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	d.InsertHead("value")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(dataset); i++ {
			d.PeekHead()
		}
	}
}

func BenchmarkPeekDoublyList(b *testing.B) {
	var d list.DoublyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	indexset := []int{6, 6, 6, 6, 4, 4, 3, 2, 0, 0}

	for _, value := range dataset {
		d.InsertHead(value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, index := range indexset {
			d.Peek(index)
		}
	}
}

func BenchmarkPeekTailDoublyList(b *testing.B) {
	var d list.DoublyLinkedList
	dataset := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	for _, value := range dataset {
		d.InsertHead(value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(dataset); i++ {
			d.PeekTail()
		}
	}
}

type valueIndex struct {
	value string
	index int
}
