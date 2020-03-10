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
	parentIndex := indexParentElement(indexPushedElem)

	for {
		if h.items[parentIndex] >= x {
			return
		}

		h.items[indexPushedElem], h.items[parentIndex] = h.items[parentIndex], h.items[indexPushedElem]

		indexPushedElem = parentIndex
		parentIndex = indexParentElement(indexPushedElem)
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

	return elem, true
}

func indexParentElement(i int) int {
	return (i - 1) / 2
}
