package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/list"
)

type listI interface {
	InsertHead(string)
	Insert(int, string) bool
	InsertTail(string)

	GetHead() (string, bool)
	Get(int) (string, bool)
	GetTail() (string, bool)

	PeekHead() (string, bool)
	Peek(int) (string, bool)
	PeekTail() (string, bool)

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

var datasetIV = []indexValue{
	{0, "two"},
	{0, "one"},
	{2, "three"},
	{3, "four"},
	{4, "six"},
	{4, "five"},
	{6, "ten"},
	{6, "nine"},
	{6, "eight"},
	{6, "seven"},
}

var expectedOutIV = []indexValue{
	{4, "five"},
	{1, "two"},
	{7, "ten"},
	{3, "six"},
	{5, "nine"},
	{3, "seven"},
	{3, "eight"},
	{0, "one"},
	{1, "four"},
	{0, "three"},
}

// HEAD
func TestSinglyListHeadSuccess(t *testing.T) {
	listHeadSuccess(t, &list.SinglyLinkedList{})
}

func TestDoublyListHeadSuccess(t *testing.T) {
	listHeadSuccess(t, &list.DoublyLinkedList{})
}

func listHeadSuccess(t *testing.T, list listI) {
	for _, value := range dataset {
		list.InsertHead(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, list.Len())

		val, ok := list.PeekHead()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "PeekHead for expected value: \"%s\"", expected)

		val, ok = list.GetHead()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "GetHead for expected value: \"%s\"", expected)
	}
}

func TestSinglyListHeadOutOfRange(t *testing.T) {
	listHeadOutOfRange(t, &list.SinglyLinkedList{})
}

func TestDoublyListHeadOutOfRange(t *testing.T) {
	listHeadOutOfRange(t, &list.DoublyLinkedList{})
}

func listHeadOutOfRange(t *testing.T, list listI) {
	// PeekHead and GetHead in empty list
	val, ok := list.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = list.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BODY
func TestSinglyListBodySuccess(t *testing.T) {
	listBodySuccess(t, &list.SinglyLinkedList{})
}

func TestDoublyListBodySuccess(t *testing.T) {
	listBodySuccess(t, &list.DoublyLinkedList{})
}

func listBodySuccess(t *testing.T, list listI) {
	for _, iv := range datasetIV {
		ok := list.Insert(iv.index, iv.value)
		assert.True(t, ok, "On insertion with index: %d and value \"%s\"", iv.index, iv.value)
	}

	for _, iv := range expectedOutIV {
		val, ok := list.Peek(iv.index)
		assert.Equal(t, iv.value, val, "On Peeking with index: %d", iv.index)
		assert.True(t, ok, "On Peeking with index: \"%d\"", iv.index)

		val, ok = list.Get(iv.index)
		assert.Equal(t, iv.value, val, "On Getting with index: %d", iv.index)
		assert.True(t, ok, "On Getting with index: %d", iv.index)
	}
}

func TestSinglyListBodyOutOfRange(t *testing.T) {
	listBodyOutOfRange(t, &list.SinglyLinkedList{})
}

func TestDoublyListBodyOutOfRange(t *testing.T) {
	listBodyOutOfRange(t, &list.DoublyLinkedList{})
}

func listBodyOutOfRange(t *testing.T, list listI) {
	// Insert out of range
	ok := list.Insert(5, "one")
	assert.False(t, ok)

	ok = list.Insert(-1, "two")
	assert.False(t, ok)

	// Peek and Get in empty list
	val, ok := list.Peek(10)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = list.Get(10)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	// for following tests list should be not empty
	list.InsertHead("stub")

	// Peek and Get out of range
	val, ok = list.Peek(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = list.Peek(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = list.Get(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = list.Get(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// TAIL
func TestSinglyListTailSuccess(t *testing.T) {
	listTailSuccess(t, &list.SinglyLinkedList{})
}

func TestDoublyListTailSuccess(t *testing.T) {
	listTailSuccess(t, &list.DoublyLinkedList{})
}

func listTailSuccess(t *testing.T, list listI) {
	for _, value := range dataset {
		list.InsertTail(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, list.Len())

		val, ok := list.PeekTail()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "PeekTail for expected value: \"%s\"", expected)

		val, ok = list.GetTail()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "GetTail for expected value: \"%s\"", expected)
	}
}

func TestSinglyListTailOutOfRange(t *testing.T) {
	listTailOutOfRange(t, &list.SinglyLinkedList{})
}

func TestDoublyListTailOutOfRange(t *testing.T) {
	listTailOutOfRange(t, &list.DoublyLinkedList{})
}

func listTailOutOfRange(t *testing.T, list listI) {
	// PeekTail and GetTail in empty list
	val, ok := list.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = list.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BENCHMARKS
// Benchmarking SinglyLinkedList:
func BenchmarkInsertHeadSinglyList(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var s list.SinglyLinkedList
		b.StartTimer()

		for _, iv := range datasetIV {
			s.Insert(iv.index, iv.value)
		}
	}
}

func BenchmarkInsertTailSinglyList(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var d list.DoublyLinkedList
		b.StartTimer()

		for _, iv := range datasetIV {
			d.Insert(iv.index, iv.value)
		}
	}
}

func BenchmarkInsertTailDoublyList(b *testing.B) {
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

type indexValue struct {
	index int
	value string
}
