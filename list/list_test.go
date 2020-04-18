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
		if !ok {
			t.Error("PeekHead for expected value:", expected, "should return true, but returns false")
		}

		val, ok = s.GetHead()
		assert.Equal(t, expected, val)
		if !ok {
			t.Error("GetHead for expected value:", expected, "should return true, but returns false")
		}
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
		if !ok {
			t.Error("PeekTail for expected value:", expected, "should return true, but returns false")
		}

		val, ok = s.GetTail()
		assert.Equal(t, expected, val)
		if !ok {
			t.Error("GetTail for expected value:", expected, "should return true, but returns false")
		}
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
		if !ok {
			t.Error("PeekHead for expected value:", expected, "should return true, but returns false")
		}

		val, ok = d.GetHead()
		assert.Equal(t, expected, val)
		if !ok {
			t.Error("GetHead for expected value:", expected, "should return true, but returns false")
		}
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
		if !ok {
			t.Error("PeekTail for expected value:", expected, "should return true, but returns false")
		}

		val, ok = d.GetTail()
		assert.Equal(t, expected, val)
		if !ok {
			t.Error("GetTail for expected value:", expected, "should return true, but returns false")
		}
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
