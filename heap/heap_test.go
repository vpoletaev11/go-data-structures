package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	var heap Heap

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)

	if heap.Len() != 4 {
		t.Error("Pushed 4 elements, but length of heap ==", heap.Len())
	}
}

func TestPush(t *testing.T) {
	var heap Heap

	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5521)
	heap.Push(111)
	heap.Push(99)
	heap.Push(13)

	expected := []int{5521, 13, 111, 4, 3, 2, 99, 1}
	assert.Equal(t, expected, heap.items)
}

func TestPopSuccess(t *testing.T) {
	var heap Heap
	heap.items = []int{5521, 13, 111, 4, 3, 2, 99, 1}

	expected := []int{5521, 111, 99, 13, 4, 3, 2, 1}
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
	var heap Heap

	elem, ok := heap.Pop()
	if elem != 0 {
		t.Error("Nothing pushed, but popped value ==", elem)
	}
	if ok != false {
		t.Error("Nothing pushed, but pop status == true")
	}
}

func TestPeekSuccess(t *testing.T) {
	var heap Heap

	heap.items = []int{1}

	elem, ok := heap.Peek()
	assert.Equal(t, 1, elem)
	assert.Equal(t, true, ok)
}

func TestPeekEmptyHeap(t *testing.T) {
	var heap Heap

	elem, ok := heap.Peek()
	assert.Equal(t, 0, elem)
	assert.Equal(t, false, ok)
}

func TestIndexParent(t *testing.T) {
	var heap Heap

	heap.items = []int{5521, 13, 111, 4, 3, 2, 99, 1}

	parent := indexParent(1)
	assert.Equal(t, 0, parent)

	parent = indexParent(3)
	assert.Equal(t, 1, parent)

	parent = indexParent(5)
	assert.Equal(t, 2, parent)

	parent = indexParent(7)
	assert.Equal(t, 3, parent)
}

func TestIndexChildLeft(t *testing.T) {
	var heap Heap

	heap.items = []int{5521, 13, 111, 4, 3, 2, 99, 1}

	leftChildIndex, exist := heap.indexChildLeft(0)
	assert.Equal(t, 1, leftChildIndex)
	assert.Equal(t, true, exist)

	leftChildIndex, exist = heap.indexChildLeft(1)
	assert.Equal(t, 3, leftChildIndex)
	assert.Equal(t, true, exist)

	leftChildIndex, exist = heap.indexChildLeft(7)
	assert.Equal(t, 0, leftChildIndex)
	assert.Equal(t, false, exist)
}

func TestIndexChildRight(t *testing.T) {
	var heap Heap

	heap.items = []int{5521, 13, 111, 4, 3, 2, 99, 1}

	rightChildIndex, exist := heap.indexChildRight(0)
	assert.Equal(t, 2, rightChildIndex)
	assert.Equal(t, true, exist)

	rightChildIndex, exist = heap.indexChildRight(1)
	assert.Equal(t, 4, rightChildIndex)
	assert.Equal(t, true, exist)

	rightChildIndex, exist = heap.indexChildRight(7)
	assert.Equal(t, 0, rightChildIndex)
	assert.Equal(t, false, exist)
}

func TestSwap(t *testing.T) {
	var heap Heap
	heap.items = []int{1, 2}

	heap.swap(0, 1)

	expected := []int{2, 1}
	assert.Equal(t, expected, heap.items)
}

func TestHeapify(t *testing.T) {
	var heap Heap
	// testing heapifying left child
	heap.items = []int{4, 11, 9}

	heap.heapify(0)

	expected := []int{11, 4, 9}
	assert.Equal(t, expected, heap.items)

	// testing heapifying right child
	heap.items = []int{4, 3, 11}

	heap.heapify(0)

	expected = []int{11, 3, 4}
	assert.Equal(t, expected, heap.items)

	// testing recursively heapifying
	heap.items = []int{4, 11, 9, 10, 5, 6, 8, 1, 2, 4}

	heap.heapify(0)

	expected = []int{11, 10, 9, 4, 5, 6, 8, 1, 2, 4}
	assert.Equal(t, expected, heap.items)
}

// TestHeapDemo displays example of using the heap data structure
func TestHeapDemo(t *testing.T) {
	fmt.Println("Demonstration of using the heap data structure:")
	var heap Heap

	fmt.Println("Pushed: 1")
	heap.Push(1)

	fmt.Println("Pushed: 8")
	heap.Push(8)

	fmt.Println("Pushed: 2")
	heap.Push(2)

	fmt.Println("Pushed: 10")
	heap.Push(10)
	fmt.Print("\n")

	for heap.Len() > 0 {
		elem, ok := heap.Pop()
		if ok {
			fmt.Println("Popped:", elem)
		}
	}
}
