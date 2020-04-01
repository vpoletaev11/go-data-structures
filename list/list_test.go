package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/list"
)

// Testing singly linked list:
// HEAD
func TestHEADSinglyLinkedList(t *testing.T) {
	var s list.SinglyLinkedList

	// InsertHead
	s.InsertHead("one")
	s.InsertHead("two")
	s.InsertHead("three")

	// Len
	assert.Equal(t, 3, s.Len())

	// Success PeekHead and GetHead
	val, ok := s.PeekHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.GetHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.PeekHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.GetHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.PeekHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = s.GetHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// PeekHead and GetHead in empty list
	val, ok = s.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BODY
func TestBODYSinglyLinkedList(t *testing.T) {
	var s list.SinglyLinkedList

	// Insert out of range
	ok := s.Insert(5, "one")
	assert.False(t, ok)

	ok = s.Insert(-1, "two")
	assert.False(t, ok)

	// Success Insert
	ok = s.Insert(0, "three")
	assert.True(t, ok)

	ok = s.Insert(0, "one")
	assert.True(t, ok)

	ok = s.Insert(1, "two")
	assert.True(t, ok)

	ok = s.Insert(3, "four")
	assert.True(t, ok)

	// Len
	assert.Equal(t, 4, s.Len())

	// Peek and Get out of range
	val, ok := s.Peek(5)
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

	// Success Peek and Get
	val, ok = s.Peek(3)
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

	// Peek and Get in empty list
	val, ok = s.Peek(10)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.Get(10)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// TAIL
func TestTAILSinglyLinkedList(t *testing.T) {
	var s list.SinglyLinkedList

	// InsertTail
	s.InsertTail("one")
	s.InsertTail("two")
	s.InsertTail("three")

	// Len
	assert.Equal(t, 3, s.Len())

	// Success PeekTail and GetTail
	val, ok := s.PeekTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.GetTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = s.PeekTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.GetTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = s.PeekTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = s.GetTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// PeekTail and GetTail in empty list
	val, ok = s.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = s.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing doubly linked list:
// HEAD
func TestHEADDoublyLinkedList(t *testing.T) {
	var d list.DoublyLinkedList

	// InsertHead
	d.InsertHead("one")
	d.InsertHead("two")
	d.InsertHead("three")

	// Len
	assert.Equal(t, 3, d.Len())

	// Success PeekHead and GetHead
	val, ok := d.PeekHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = d.GetHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = d.PeekHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = d.GetHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = d.PeekHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = d.GetHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// PeekHead and GetHead in empty list
	val, ok = d.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BODY
func TestBODYDoublyLinkedList(t *testing.T) {
	var d list.DoublyLinkedList

	// // Insert out of range
	ok := d.Insert(5, "one")
	assert.False(t, ok)

	ok = d.Insert(-1, "two")
	assert.False(t, ok)

	// Success Insert
	ok = d.Insert(0, "three")
	assert.True(t, ok)

	ok = d.Insert(0, "one")
	assert.True(t, ok)

	ok = d.Insert(1, "two")
	assert.True(t, ok)

	ok = d.Insert(3, "four")
	assert.True(t, ok)

	// Len
	assert.Equal(t, 4, d.Len())

	// // Peek and Get out of range
	val, ok := d.Peek(5)
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

	// Success Peek and Get
	val, ok = d.Peek(3)
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

	// Peek and Get in empty list
	val, ok = d.Peek(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.Get(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// TAIL
func TestDoublyLinkedList(t *testing.T) {
	var d list.DoublyLinkedList

	// InsertTail
	d.InsertTail("one")
	d.InsertTail("two")
	d.InsertTail("three")

	// Len
	assert.Equal(t, 3, d.Len())

	// Success PeekTail and GetTail
	val, ok := d.PeekTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = d.GetTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = d.PeekTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = d.GetTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = d.PeekTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = d.GetTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// PeekTail and GetTail in empty list
	val, ok = d.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = d.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}
