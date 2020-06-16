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
		assert.True(t, ok, "Peek for expected value: \"%d\" %s", expected, "should return true, but returns false")

		val, ok = h.Pop()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Pop for expected value: \"%d\" %s", expected, "should return true, but returns false")
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
		assert.True(t, b.Find(value), "Value: %d %s", value, "been inserted, but not found")

		b.Remove(value)

		assert.False(t, b.Find(value), "Value: %d %s", value, "been removed, but were found")
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
		assert.True(t, a.Find(value), "Value: %d %s", value, "been inserted, but not found")

		a.Remove(value)

		assert.False(t, a.Find(value), "Value: %d %s", value, "been removed, but were found")
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

// BENCHMARKS
// Benchmarking BinarySearchTree:
func BenchmarkInsertBST(b *testing.B) {
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var t tree.BinarySearchTree
		b.StartTimer()

		for _, value := range dataset {
			t.Insert(value)
		}
	}
}

func BenchmarkRemoveBST(b *testing.B) {
	var t tree.BinarySearchTree
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			t.Insert(value)
		}
		b.StartTimer()

		for _, value := range dataset {
			t.Remove(value)
		}
	}
}

func BenchmarkFindBST(b *testing.B) {
	var t tree.BinarySearchTree
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range dataset {
		t.Insert(value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, value := range dataset {
			t.Remove(value)
		}
	}
}

// Benchmarking AVLTree:
func BenchmarkInsertAVL(b *testing.B) {
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var t tree.AVLTree
		b.StartTimer()

		for _, value := range dataset {
			t.Insert(value)
		}
	}
}

func BenchmarkRemoveAVL(b *testing.B) {
	var t tree.AVLTree
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			t.Insert(value)
		}
		b.StartTimer()

		for _, value := range dataset {
			t.Remove(value)
		}
	}
}

func BenchmarkFindAVL(b *testing.B) {
	var t tree.AVLTree
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range dataset {
		t.Insert(value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, value := range dataset {
			t.Remove(value)
		}
	}
}

// Benchmarking Heap:
func BenchmarkInsertHeap(b *testing.B) {
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var h tree.Heap
		b.StartTimer()

		for _, value := range dataset {
			h.Push(value)
		}
	}
}

func BenchmarkRemoveHeap(b *testing.B) {
	var h tree.Heap
	dataset := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range dataset {
			h.Push(value)
		}
		b.StartTimer()

		for i := 0; i < len(dataset); i++ {
			h.Pop()
		}
	}
}

func BenchmarkPeekHeap(b *testing.B) {
	var h tree.Heap
	h.Push(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h.Peek()
	}
}
