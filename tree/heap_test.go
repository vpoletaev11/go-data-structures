package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-data-structures/tree"
)

func TestLen(t *testing.T) {
	var heap tree.Heap

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)

	assert.Equal(t, 4, heap.Len())
}

func TestPush(t *testing.T) {
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

func TestPopSuccess(t *testing.T) {
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

func TestPeekSuccess(t *testing.T) {
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
