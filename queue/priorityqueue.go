package queue

// nodePQ - node of priority queue
type nodePQ struct {
	priority int
	data     string
}

// PriorityQueue - heap based list of elements organized by priority of contained elements
type PriorityQueue struct {
	nodes []nodePQ
}

// Len returns length of the PriorityQueue
func (p *PriorityQueue) Len() int {
	return len(p.nodes)
}

// Enqueue appends element into PriorityQueue
func (p *PriorityQueue) Enqueue(val string, priority int) {
	p.nodes = append(p.nodes, nodePQ{data: val, priority: priority})

	indexPushedElem := p.Len() - 1              // Get index of the last inserted element
	parentIndex := indexParent(indexPushedElem) // Get index of the parent of the last inserted element

	for {
		if p.nodes[parentIndex].priority >= priority { // Check if parent node priority >= newly inserted node priority
			return
		}

		p.swap(indexPushedElem, parentIndex)
		indexPushedElem = parentIndex
		parentIndex = indexParent(indexPushedElem)
	}
}

// Dequeue obtains and deletes element with biggest priority from PriorityQueue
func (p *PriorityQueue) Dequeue() (val string, ok bool) {
	if p.Len() == 0 {
		return "", false
	}

	val = p.nodes[0].data           // Obtain the root element
	p.nodes[0] = p.nodes[p.Len()-1] // Set the last node in place of the root node
	p.nodes = p.nodes[:p.Len()-1]   // Remove the last node

	p.heapify(0)

	return val, true
}

// Peek obtains element from PriorityQueue
func (p *PriorityQueue) Peek() (val string, ok bool) {
	if p.Len() == 0 {
		return "", false
	}

	return p.nodes[0].data, true
}

// indexParent returns parent node index for inserted node index
func indexParent(i int) int {
	return (i - 1) / 2
}

// indexChildLeft returns index of the left child for inserted parent index
func (p *PriorityQueue) indexChildLeft(parentIndex int) (childIndex int, exist bool) {
	childIndex = 2*parentIndex + 1
	if childIndex > p.Len()-1 {
		return 0, false
	}
	return childIndex, true
}

// indexChildRight returns index of the right child for inserted parent index
func (p *PriorityQueue) indexChildRight(parentIndex int) (childIndex int, exist bool) {
	childIndex = 2*parentIndex + 2
	if childIndex > p.Len()-1 {
		return 0, false
	}
	return childIndex, true
}

// swap swaps tho nodes
func (p *PriorityQueue) swap(index1, index2 int) {
	p.nodes[index1], p.nodes[index2] = p.nodes[index2], p.nodes[index1]
}

// heapify restores the property of ordering in the entire subtree, the root of which is the rootIndex
func (p *PriorityQueue) heapify(rootIndex int) {
	largest := rootIndex

	// Check if left child priority bigger than rootIndex
	indexLeftChild, exist := p.indexChildLeft(rootIndex)
	if exist && (p.nodes[indexLeftChild].priority > p.nodes[largest].priority) {
		largest = indexLeftChild
	}

	// Check if right child priority bigger than (rootIndex or left child) priority
	indexRightChild, exist := p.indexChildRight(rootIndex)
	if exist && (p.nodes[indexRightChild].priority > p.nodes[largest].priority) {
		largest = indexRightChild
	}

	if rootIndex != largest {
		p.swap(rootIndex, largest)
		p.heapify(largest)
	}
}
