package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/list"
)

// Testing singly linked list:
// HEAD
func TestHEADSinglyLinkedList(t *testing.T) {
	var l list.SinglyLinkedList

	// InsertHead
	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	// Len
	assert.Equal(t, 3, l.Len())

	// Success PeekHead and GetHead
	val, ok := l.PeekHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.PeekHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.PeekHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// PeekHead and GetHead in empty list
	val, ok = l.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BODY
func TestBODYSinglyLinkedList(t *testing.T) {
	var l list.SinglyLinkedList

	// Insert out of range
	ok := l.Insert(5, "one")
	assert.False(t, ok)

	ok = l.Insert(-1, "two")
	assert.False(t, ok)

	// Success Insert
	ok = l.Insert(0, "three")
	assert.True(t, ok)

	ok = l.Insert(0, "one")
	assert.True(t, ok)

	ok = l.Insert(1, "two")
	assert.True(t, ok)

	ok = l.Insert(3, "four")
	assert.True(t, ok)

	// Len
	assert.Equal(t, 4, l.Len())

	// Peek and Get out of range
	val, ok := l.Peek(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Peek(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Get(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Get(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	// Success Peek and Get
	val, ok = l.Peek(3)
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = l.Get(3)
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = l.Peek(2)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.Get(2)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.Peek(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Get(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Peek(0)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.Get(0)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// Peek and Get in empty list
	val, ok = l.Peek(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Get(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// TAIL
func TestTAILSinglyLinkedList(t *testing.T) {
	var l list.SinglyLinkedList

	// InsertTail
	l.InsertTail("one")
	l.InsertTail("two")
	l.InsertTail("three")

	// Len
	assert.Equal(t, 3, l.Len())

	// Success PeekTail and GetTail
	val, ok := l.PeekTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.PeekTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.PeekTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// PeekTail and GetTail in empty list
	val, ok = l.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// Testing doubly linked list:
// HEAD
func TestHEADDoublyLinkedList(t *testing.T) {
	var l list.DoublyLinkedList

	// InsertHead
	l.InsertHead("one")
	l.InsertHead("two")
	l.InsertHead("three")

	// Len
	assert.Equal(t, 3, l.Len())

	// Success PeekHead and GetHead
	val, ok := l.PeekHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.PeekHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.PeekHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// PeekHead and GetHead in empty list
	val, ok = l.PeekHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.GetHead()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// BODY
func TestBODYDoublyLinkedList(t *testing.T) {
	var l list.DoublyLinkedList

	// // Insert out of range
	ok := l.Insert(5, "one")
	assert.False(t, ok)

	ok = l.Insert(-1, "two")
	assert.False(t, ok)

	// Success Insert
	ok = l.Insert(0, "three")
	assert.True(t, ok)

	ok = l.Insert(0, "one")
	assert.True(t, ok)

	ok = l.Insert(1, "two")
	assert.True(t, ok)

	ok = l.Insert(3, "four")
	assert.True(t, ok)

	// Len
	assert.Equal(t, 4, l.Len())

	// // Peek and Get out of range
	val, ok := l.Peek(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Peek(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Get(5)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Get(-1)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	// Success Peek and Get
	val, ok = l.Peek(3)
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = l.Get(3)
	assert.Equal(t, "four", val)
	assert.True(t, ok)

	val, ok = l.Peek(2)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.Get(2)
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.Peek(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Get(1)
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.Peek(0)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.Get(0)
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// Peek and Get in empty list
	val, ok = l.Peek(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.Get(0)
	assert.Equal(t, "", val)
	assert.False(t, ok)
}

// TAIL
func TestDoublyLinkedList(t *testing.T) {
	var l list.DoublyLinkedList

	// InsertTail
	l.InsertTail("one")
	l.InsertTail("two")
	l.InsertTail("three")

	// Len
	assert.Equal(t, 3, l.Len())

	// Success PeekTail and GetTail
	val, ok := l.PeekTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "three", val)
	assert.True(t, ok)

	val, ok = l.PeekTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "two", val)
	assert.True(t, ok)

	val, ok = l.PeekTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "one", val)
	assert.True(t, ok)

	// PeekTail and GetTail in empty list
	val, ok = l.PeekTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)

	val, ok = l.GetTail()
	assert.Equal(t, "", val)
	assert.False(t, ok)
}
