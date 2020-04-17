package tree_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/tree"
)

// Testing Heap:
func TestHeapSuccess(t *testing.T) {
	var h tree.Heap

	dataset := []int{1, 2, 3, 4, 5521, 111, 99, 13}
	expectedOut := []int{5521, 111, 99, 13, 4, 3, 2, 1}

	for _, value := range dataset {
		h.Push(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, h.Len())

		val, ok := h.Peek()
		assert.Equal(t, expected, val)
		if !ok {
			t.Error("Peek for expected value:", expected, "should return true, but returns false")
		}

		val, ok = h.Pop()
		assert.Equal(t, expected, val)
		if !ok {
			t.Error("Pop for expected value:", expected, "should return true, but returns false")
		}
	}
}

func TestHeapOutOfRange(t *testing.T) {
	var h tree.Heap

	assert.Equal(t, 0, h.Len())

	// Peek and Pop in empty heap
	val, ok := h.Peek()
	assert.Equal(t, 0, val)
	assert.False(t, ok)

	val, ok = h.Pop()
	assert.Equal(t, 0, val)
	assert.False(t, ok)
}

// Testing BinarySearchTree:
func TestBSTSuccess(t *testing.T) {
	var b tree.BinarySearchTree

	dataset := rand.Perm(10000)

	for _, value := range dataset {
		b.Insert(value)
	}

	for _, value := range dataset {
		if !b.Find(value) {
			t.Error("Value:", value, "been inserted, but not found")
		}
		b.Remove(value)
		if b.Find(value) {
			t.Error("Value:", value, "been removed, but were found")
		}
	}
}

func TestBSTOutOfRange(t *testing.T) {
	var b tree.BinarySearchTree

	// Search and Delete values in empty tree
	assert.False(t, b.Find(50))
	b.Remove(50)

	// Insert exists value
	b.Insert(100)
	b.Insert(100)

	// Search and Delete values out of range
	assert.False(t, b.Find(0))
	b.Remove(0)

	assert.False(t, b.Find(1000))
	b.Remove(1000)
}

// Testing AVLTree:
func TestAVLSuccess(t *testing.T) {
	var a tree.AVLTree

	dataset := rand.Perm(10000)

	for _, value := range dataset {
		a.Insert(value)
	}

	for _, value := range dataset {
		if !a.Find(value) {
			t.Error("Value:", value, "been inserted, but not found")
		}
		a.Remove(value)
		if a.Find(value) {
			t.Error("Value:", value, "been removed, but were found")
		}
	}
}

func TestAVLOutOfRange(t *testing.T) {
	var a tree.AVLTree

	// Search and Delete values in empty tree
	assert.False(t, a.Find(50))
	a.Remove(50)

	// Insert exists value
	a.Insert(100)
	a.Insert(100)

	// Search and Delete values out of range
	assert.False(t, a.Find(0))
	a.Remove(0)

	assert.False(t, a.Find(1000))
	a.Remove(1000)
}
