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
