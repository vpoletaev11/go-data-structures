package heap

// Heap - binary tree where:
// 1) The value at any vertex is no less than the values ​​of its descendants;
// 2) The depth of all leaves (distance to the root) differs by no more than 1 layer;
// 3) The last layer is filled from left to right without gaps.
type Heap struct {
	items []int
}

// Len returns length of the heap
func (h *Heap) Len() int {
	return len(h.items)
}

// Push appends element into Heap
func (h *Heap) Push(x int) {
	h.items = append(h.items, x)

	indexPushedElem := h.Len() - 1
	parentIndex := indexParent(indexPushedElem)

	for {
		if h.items[parentIndex] >= x {
			return
		}

		h.items[indexPushedElem], h.items[parentIndex] = h.items[parentIndex], h.items[indexPushedElem]

		indexPushedElem = parentIndex
		parentIndex = indexParent(indexPushedElem)
	}
}

// Pop obtains and deletes root element from heap
func (h *Heap) Pop() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}

	elem := h.items[0]
	h.items[0] = h.items[h.Len()-1]
	h.items = h.items[:h.Len()-1]

	h.heapify(0)

	return elem, true
}

func indexParent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) indexChildLeft(rootIndex int) (childIndex int, exist bool) {
	childIndex = 2*rootIndex + 1
	if childIndex > h.Len()-1 {
		return 0, false
	}
	return childIndex, true
}

func (h *Heap) indexChildRight(rootIndex int) (childIndex int, exist bool) {
	childIndex = 2*rootIndex + 2
	if childIndex > h.Len()-1 {
		return 0, false
	}
	return childIndex, true
}

func (h *Heap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}

func (h *Heap) heapify(rootIndex int) {
	largest := rootIndex
	indexLeftChild, exist := h.indexChildLeft(rootIndex)
	if exist && (h.items[indexLeftChild] > h.items[largest]) {
		largest = indexLeftChild
	}

	indexRightChild, exist := h.indexChildRight(rootIndex)
	if exist && (h.items[indexRightChild] > h.items[largest]) {
		largest = indexRightChild
	}
	if rootIndex != largest {
		h.swap(rootIndex, largest)
		h.heapify(largest)
	}
}
