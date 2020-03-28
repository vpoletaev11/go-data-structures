package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/tree"
)

// testing Heap:
func TestPushPopPeekLenHeap(t *testing.T) {
	var h tree.Heap

	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5521)
	h.Push(111)
	h.Push(99)
	h.Push(13)

	// Len
	assert.Equal(t, 8, h.Len())

	// Peek valid
	val, ok := h.Peek()
	assert.Equal(t, 5521, val)
	assert.True(t, ok)

	// Pop valid
	val, ok = h.Pop()
	assert.Equal(t, 5521, val)
	assert.True(t, ok)

	val, ok = h.Pop()
	assert.Equal(t, 111, val)
	assert.True(t, ok)

	val, ok = h.Pop()
	assert.Equal(t, 99, val)
	assert.True(t, ok)

	val, ok = h.Pop()
	assert.Equal(t, 13, val)
	assert.True(t, ok)

	val, ok = h.Pop()
	assert.Equal(t, 4, val)
	assert.True(t, ok)

	val, ok = h.Pop()
	assert.Equal(t, 3, val)
	assert.True(t, ok)

	val, ok = h.Pop()
	assert.Equal(t, 2, val)
	assert.True(t, ok)

	val, ok = h.Pop()
	assert.Equal(t, 1, val)
	assert.True(t, ok)

	// Pop empty heap
	elem, ok := h.Pop()
	assert.Equal(t, 0, elem)
	assert.False(t, ok)

	// Peek empty heap
	val, ok = h.Peek()
	assert.Equal(t, 0, val)
	assert.False(t, ok)
}

// Testing Binary Search Tree:
func TestInsertRemoveFindBST(t *testing.T) {
	var b tree.BinarySearchTree

	b.Insert(50)
	b.Insert(50) // Inserting exists value
	b.Insert(70)
	b.Insert(60)
	b.Insert(90)
	b.Insert(65)
	b.Insert(62)
	b.Insert(100)
	b.Insert(95)
	b.Insert(110)
	b.Insert(48)
	b.Insert(49)
	b.Insert(41)
	b.Insert(42)

	// Search and Delete values out of range
	assert.False(t, b.Find(0))
	b.Remove(0)

	assert.False(t, b.Find(0))
	b.Remove(1000)

	// Valid Finding and Removing; checking that the value does not exist after deletion
	assert.True(t, b.Find(50))
	b.Remove(50)
	assert.False(t, b.Find(50))

	assert.True(t, b.Find(60))
	b.Remove(60)
	assert.False(t, b.Find(60))

	assert.True(t, b.Find(65))
	b.Remove(65)
	assert.False(t, b.Find(65))

	assert.True(t, b.Find(62))
	b.Remove(62)
	assert.False(t, b.Find(62))

	assert.True(t, b.Find(70))
	b.Remove(70)
	assert.False(t, b.Find(70))

	assert.True(t, b.Find(100))
	b.Remove(100)
	assert.False(t, b.Find(100))

	assert.True(t, b.Find(110))
	b.Remove(110)
	assert.False(t, b.Find(110))

	assert.True(t, b.Find(95))
	b.Remove(95)
	assert.False(t, b.Find(95))

	assert.True(t, b.Find(48))
	b.Remove(48)
	assert.False(t, b.Find(48))

	assert.True(t, b.Find(41))
	b.Remove(41)
	assert.False(t, b.Find(41))

	assert.True(t, b.Find(49))
	b.Remove(49)
	assert.False(t, b.Find(49))

	assert.True(t, b.Find(90))
	b.Remove(90)
	assert.False(t, b.Find(90))

	assert.True(t, b.Find(42))
	b.Remove(42)
	assert.False(t, b.Find(42))

	// Search and Delete values in empty tree
	assert.False(t, b.Find(0))
	b.Remove(0)

}
