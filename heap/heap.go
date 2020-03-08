package heap

import (
	"fmt"
)

// Heap - binary tree where:
// 1) The value at any vertex is no less than the values ​​of its descendants;
// 2) The depth of all leaves (distance to the root) differs by no more than 1 layer;
// 3) The last layer is filled from left to right without gaps.
type Heap []int

// Push appends element into Heap
func (h *Heap) Push(x int) {
	*h = append(*h, x)

	pushedElemIndex := len(*h) - 1
	parentIndex := h.parentIndex(pushedElemIndex)

	for {
		if (*h)[parentIndex] < x {
			(*h)[pushedElemIndex], (*h)[parentIndex] = (*h)[parentIndex], (*h)[pushedElemIndex]
		} else {
			return
		}
		pushedElemIndex = parentIndex
		parentIndex = h.parentIndex(pushedElemIndex)
	}
}

func (h *Heap) parentIndex(i int) int {
	return (i - 1) / 2
}

// Pop obtains and deletes root element from heap
func (h *Heap) Pop() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}

	elem := (*h)[0]
	*h = (*h)[1:]
	return elem, true
}

// Demo displays example of using the heap data structure
func Demo() {
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

	for len(heap) > 0 {
		elem, ok := heap.Pop()
		if ok {
			fmt.Println("Popped:", elem)
		}
	}
}
