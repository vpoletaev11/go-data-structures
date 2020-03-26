package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/tree"
)

// testing Heap:

// test Len
func TestLenHeap(t *testing.T) {
	var heap tree.Heap

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)

	assert.Equal(t, 4, heap.Len())
}

// test Push
func TestPushHeap(t *testing.T) {
	var heap tree.Heap

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5521)
	heap.Push(111)
	heap.Push(99)
	heap.Push(13)

	actual := []int{}
	for heap.Len() > 0 {
		elem, ok := heap.Pop()
		if ok {
			actual = append(actual, elem)
		}
	}

	expected := []int{5521, 111, 99, 13, 4, 3, 2, 1}
	assert.Equal(t, expected, actual)
}

// test Pop
func TestPopHeapSuccess(t *testing.T) {
	var heap tree.Heap

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)

	expected := []int{4, 3, 2, 1}
	actual := []int{}
	for heap.Len() > 0 {
		elem, ok := heap.Pop()
		if ok {
			actual = append(actual, elem)
		}
	}

	assert.Equal(t, expected, actual)
}

func TestPopEmptyHeap(t *testing.T) {
	var heap tree.Heap

	elem, ok := heap.Pop()
	assert.Equal(t, 0, elem)
	assert.False(t, ok)
}

// test Peek
func TestPeekHeapSuccess(t *testing.T) {
	var heap tree.Heap

	heap.Push(1)

	elem, ok := heap.Peek()
	assert.Equal(t, 1, elem)
	assert.Equal(t, true, ok)
}

func TestPeekEmptyHeap(t *testing.T) {
	var heap tree.Heap

	elem, ok := heap.Peek()
	assert.Equal(t, 0, elem)
	assert.Equal(t, false, ok)
}

// Testing binary search tree:

// test Insert & Find
func TestInsertAndFindBST(t *testing.T) {
	var b tree.BinarySearchTree

	b.Insert(8)
	b.Insert(8)
	b.Insert(3)
	b.Insert(10)
	b.Insert(1)
	b.Insert(6)
	b.Insert(14)
	b.Insert(4)
	b.Insert(7)
	b.Insert(13)

	ok := b.Find(8)
	assert.True(t, ok)

	ok = b.Find(3)
	assert.True(t, ok)

	ok = b.Find(10)
	assert.True(t, ok)

	ok = b.Find(1)
	assert.True(t, ok)

	ok = b.Find(6)
	assert.True(t, ok)

	ok = b.Find(14)
	assert.True(t, ok)

	ok = b.Find(4)
	assert.True(t, ok)

	ok = b.Find(7)
	assert.True(t, ok)

	ok = b.Find(13)
	assert.True(t, ok)

	ok = b.Find(1000)
	assert.False(t, ok)

	ok = b.Find(0)
	assert.False(t, ok)
}

func TestFindEmptyTreeBST(t *testing.T) {
	var b tree.BinarySearchTree

	ok := b.Find(10)
	assert.False(t, ok)
}

// test Remove
func TestRemoveBST(t *testing.T) {
	var b tree.BinarySearchTree

	b.Insert(8)
	b.Insert(3)
	b.Insert(1)
	b.Insert(14)
	b.Insert(4)
	b.Insert(7)
	b.Insert(2)
	b.Insert(12)
	b.Insert(11)
	b.Insert(50)

	b.Remove(1000)
	b.Remove(0)

	b.Remove(8)
	assert.False(t, b.Find(8))

	b.Remove(1)
	assert.False(t, b.Find(1))

	b.Remove(3)
	assert.False(t, b.Find(3))

	b.Remove(7)
	assert.False(t, b.Find(7))

	b.Remove(2)
	assert.False(t, b.Find(2))

	b.Remove(14)
	assert.False(t, b.Find(14))

	b.Remove(50)
	assert.False(t, b.Find(50))

	b.Remove(4)
	assert.False(t, b.Find(4))

	b.Remove(11)
	assert.False(t, b.Find(11))

	b.Remove(12)
	assert.False(t, b.Find(12))

	b.Remove(0)
}
