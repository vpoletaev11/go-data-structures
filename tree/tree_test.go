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

	b.Insert(8)
	b.Insert(8) // Inserting exists value
	b.Insert(3)
	b.Insert(1)
	b.Insert(14)
	b.Insert(4)
	b.Insert(7)
	b.Insert(2)
	b.Insert(12)
	b.Insert(11)
	b.Insert(50)

	// Search and Delete values out of range
	assert.False(t, b.Find(1000))
	b.Remove(1000)

	assert.False(t, b.Find(0))
	b.Remove(0)

	// Valid Finding and Removing; checking that the value does not exist after deletion
	assert.True(t, b.Find(8))
	b.Remove(8)
	assert.False(t, b.Find(8))

	assert.True(t, b.Find(1))
	b.Remove(1)
	assert.False(t, b.Find(1))

	assert.True(t, b.Find(3))
	b.Remove(3)
	assert.False(t, b.Find(3))

	assert.True(t, b.Find(7))
	b.Remove(7)
	assert.False(t, b.Find(7))

	assert.True(t, b.Find(2))
	b.Remove(2)
	assert.False(t, b.Find(2))

	assert.True(t, b.Find(14))
	b.Remove(14)
	assert.False(t, b.Find(14))

	assert.True(t, b.Find(50))
	b.Remove(50)
	assert.False(t, b.Find(50))

	assert.True(t, b.Find(4))
	b.Remove(4)
	assert.False(t, b.Find(4))

	assert.True(t, b.Find(11))
	b.Remove(11)
	assert.False(t, b.Find(11))

	assert.True(t, b.Find(12))
	b.Remove(12)
	assert.False(t, b.Find(12))

	// Search and Delete values in empty tree
	assert.False(t, b.Find(0))
	b.Remove(0)
}
