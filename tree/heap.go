package tree

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

// Push appends element into Heap
func (h *Heap) Push(element int) {
	h.elements = append(h.elements, element)

	indexPushedElem := h.Len() - 1              // Get index of the last inserted element
	parentIndex := indexParent(indexPushedElem) // Get index of the parent of the last inserted element

	for {
		if h.elements[parentIndex] >= element {
			return
		}

		h.swap(indexPushedElem, parentIndex)
		indexPushedElem = parentIndex
		parentIndex = indexParent(indexPushedElem)
	}
}

// Pop obtains and deletes root element from heap
func (h *Heap) Pop() (element int, status bool) {
	if h.Len() == 0 {
		return 0, false
	}

	element = h.elements[0]               // Obtain the root element
	h.elements[0] = h.elements[h.Len()-1] // Set the last node in place of the root node
	h.elements = h.elements[:h.Len()-1]   // Remove the last node

	h.heapify(0)

	return element, true
}

// Peek obtains element from heap
func (h *Heap) Peek() (element int, status bool) {
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
