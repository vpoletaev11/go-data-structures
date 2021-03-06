package tree

/*
	HEAP:

	Len()                Push()                   Pop()                    Peek()
	Best:    O(1)        Best:    O(1)            Best:    O(1)            Best:    O(1)
	Average: O(1)        Average: O(log n)        Average: O(log n)        Average: O(1)
	Worst:   O(1)        Worst:   O(log n)        Worst:   O(log n)        Worst:   O(1)


	goos: linux
	goarch: amd64
	BenchmarkInsertHeap-8   	 1658200	       727 ns/op	     256 B/op	       5 allocs/op
	BenchmarkRemoveHeap-8   	 5368878	       220 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPeekHeap-8     	1000000000	         0.575 ns/op	       0 B/op	       0 allocs/op
*/

// Heap - binary tree where:
// 1) The value at any parent node is no less than the values ​​of its children;
// 2) The depth of all leaves (distance to the root) differs by no more than 1 layer;
// 3) The last layer is filled from left to right without gaps.
type Heap struct {
	elements []int
}

// Len returns length of the heap
func (h *Heap) Len() int {
	return len(h.elements)
}

// Push appends value into Heap
func (h *Heap) Push(val int) {
	h.elements = append(h.elements, val)

	indexPushedElem := h.Len() - 1              // Get index of the last inserted element
	parentIndex := indexParent(indexPushedElem) // Get index of the parent of the last inserted element

	for {
		if h.elements[parentIndex] >= val {
			return
		}

		h.swap(indexPushedElem, parentIndex)
		indexPushedElem = parentIndex
		parentIndex = indexParent(indexPushedElem)
	}
}

// Pop obtains and deletes root element from heap
func (h *Heap) Pop() (val int, status bool) {
	if h.Len() == 0 {
		return 0, false
	}

	val = h.elements[0]                   // Obtain the root element
	h.elements[0] = h.elements[h.Len()-1] // Set the last node in place of the root node
	h.elements = h.elements[:h.Len()-1]   // Remove the last node

	h.heapify(0)

	return val, true
}

// Peek obtains value from heap
func (h *Heap) Peek() (val int, status bool) {
	if h.Len() == 0 {
		return 0, false
	}

	return h.elements[0], true
}

// indexParent returns parent node index for inserted node index
func indexParent(i int) int {
	return (i - 1) / 2
}

// indexChildLeft returns index of the left child for inserted parent index
func (h *Heap) indexChildLeft(parentIndex int) (childIndex int, exist bool) {
	childIndex = 2*parentIndex + 1
	if childIndex > h.Len()-1 {
		return 0, false
	}
	return childIndex, true
}

// indexChildRight returns index of the right child for inserted parent index
func (h *Heap) indexChildRight(parentIndex int) (childIndex int, exist bool) {
	childIndex = 2*parentIndex + 2
	if childIndex > h.Len()-1 {
		return 0, false
	}
	return childIndex, true
}

// swap swaps tho nodes
func (h *Heap) swap(index1, index2 int) {
	h.elements[index1], h.elements[index2] = h.elements[index2], h.elements[index1]
}

// heapify restores the property of ordering in the entire subtree, the root of which is the rootIndex
func (h *Heap) heapify(rootIndex int) {
	largest := rootIndex

	indexLeftChild, exist := h.indexChildLeft(rootIndex)
	if exist && (h.elements[indexLeftChild] > h.elements[largest]) {
		largest = indexLeftChild
	}

	indexRightChild, exist := h.indexChildRight(rootIndex)
	if exist && (h.elements[indexRightChild] > h.elements[largest]) {
		largest = indexRightChild
	}

	if rootIndex != largest {
		h.swap(rootIndex, largest)
		h.heapify(largest)
	}
}
