package tree_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/tree"
)

type treeI interface {
	Insert(int)
	Remove(int)
	Find(int) bool
}

var datasetSmall = []int{1, 2, 3, 5521, 8, 111, 99, 13, 9, 10}
var datasetBig = rand.Perm(10000)

// Testing Trees:
func TestBSTSuccess(t *testing.T) {
	treeSuccess(t, &tree.BinarySearchTree{})
}

func TestAVLSuccess(t *testing.T) {
	treeSuccess(t, &tree.AVLTree{})
}

func treeSuccess(t *testing.T, tree treeI) {
	for _, value := range datasetBig {
		tree.Insert(value)
	}

	for _, value := range datasetBig {
		assert.True(t, tree.Find(value), "Value: %d %s", value, "been inserted, but not found")

		tree.Remove(value)

		assert.False(t, tree.Find(value), "Value: %d %s", value, "been removed, but were found")
	}
}

func TestBSTOutOfRange(t *testing.T) {
	treeOutOfRange(t, &tree.BinarySearchTree{})
}

func TestAVLOutOfRange(t *testing.T) {
	treeOutOfRange(t, &tree.AVLTree{})
}

func treeOutOfRange(t *testing.T, tree treeI) {
	// Search and Delete values in empty tree
	assert.False(t, tree.Find(50))
	tree.Remove(50)

	// Insert exists value
	tree.Insert(100)
	tree.Insert(100)

	// Search and Delete values out of range
	assert.False(t, tree.Find(0))
	tree.Remove(0)

	assert.False(t, tree.Find(1000))
	tree.Remove(1000)
}

// Testing Heap:
func TestHeapSuccess(t *testing.T) {
	var h tree.Heap

	expectedOut := []int{5521, 111, 99, 13, 10, 9, 8, 3, 2, 1}

	for _, value := range datasetSmall {
		h.Push(value)
	}

	for i, expected := range expectedOut {
		assert.Equal(t, len(expectedOut)-i, h.Len())

		val, ok := h.Peek()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Peek for expected value: \"%d\"", expected)

		val, ok = h.Pop()
		assert.Equal(t, expected, val)
		assert.True(t, ok, "Pop for expected value: \"%d\"", expected)
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

// BENCHMARKS
// Benchmarking BinarySearchTree:
func BenchmarkInsertBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var t tree.BinarySearchTree
		b.StartTimer()

		for _, value := range datasetSmall {
			t.Insert(value)
		}
	}
}

func BenchmarkRemoveBST(b *testing.B) {
	var t tree.BinarySearchTree
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range datasetSmall {
			t.Insert(value)
		}
		b.StartTimer()

		for _, value := range datasetSmall {
			t.Remove(value)
		}
	}
}

func BenchmarkFindBST(b *testing.B) {
	var t tree.BinarySearchTree

	for _, value := range datasetSmall {
		t.Insert(value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, value := range datasetSmall {
			t.Remove(value)
		}
	}
}

// Benchmarking AVLTree:
func BenchmarkInsertAVL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var t tree.AVLTree
		b.StartTimer()

		for _, value := range datasetSmall {
			t.Insert(value)
		}
	}
}

func BenchmarkRemoveAVL(b *testing.B) {
	var t tree.AVLTree
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range datasetSmall {
			t.Insert(value)
		}
		b.StartTimer()

		for _, value := range datasetSmall {
			t.Remove(value)
		}
	}
}

func BenchmarkFindAVL(b *testing.B) {
	var t tree.AVLTree

	for _, value := range datasetSmall {
		t.Insert(value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, value := range datasetSmall {
			t.Remove(value)
		}
	}
}

// Benchmarking Heap:
func BenchmarkInsertHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var h tree.Heap
		b.StartTimer()

		for _, value := range datasetSmall {
			h.Push(value)
		}
	}
}

func BenchmarkRemoveHeap(b *testing.B) {
	var h tree.Heap
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, value := range datasetSmall {
			h.Push(value)
		}
		b.StartTimer()

		for i := 0; i < len(datasetSmall); i++ {
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
