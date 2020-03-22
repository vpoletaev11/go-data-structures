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
func (q *PriorityQueue) Len() int {
	return len(q.nodes)
}

// Push appends element into PriorityQueue
func (q *PriorityQueue) Push(val string, priority int) {
	q.nodes = append(q.nodes, nodePQ{data: val, priority: priority})

	indexPushedElem := q.Len() - 1              // Get index of the last inserted element
	parentIndex := indexParent(indexPushedElem) // Get index of the parent of the last inserted element

	for {
		if q.nodes[parentIndex].priority >= priority {
			return
		}

		q.swap(indexPushedElem, parentIndex)
		indexPushedElem = parentIndex
		parentIndex = indexParent(indexPushedElem)
	}
}

// Pop obtains and deletes element with biggest priority from PriorityQueue
func (q *PriorityQueue) Pop() (val string, ok bool) {
	if q.Len() == 0 {
		return "", false
	}

	val = q.nodes[0].data           // Obtain the root element
	q.nodes[0] = q.nodes[q.Len()-1] // Set the last node in place of the root node
	q.nodes = q.nodes[:q.Len()-1]   // Remove the last node

	q.heapify(0)

	return val, true
}

// Peek obtains element from heap
func (q *PriorityQueue) Peek() (val string, ok bool) {
	if q.Len() == 0 {
		return "", false
	}

	return q.nodes[0].data, true
}

// indexParent returns parent node index for inserted node index
func indexParent(i int) int {
	return (i - 1) / 2
}

// indexChildLeft returns index of the left child for inserted parent index
func (q *PriorityQueue) indexChildLeft(parentIndex int) (childIndex int, exist bool) {
	childIndex = 2*parentIndex + 1
	if childIndex > q.Len()-1 {
		return 0, false
	}
	return childIndex, true
}

// indexChildRight returns index of the right child for inserted parent index
func (q *PriorityQueue) indexChildRight(parentIndex int) (childIndex int, exist bool) {
	childIndex = 2*parentIndex + 2
	if childIndex > q.Len()-1 {
		return 0, false
	}
	return childIndex, true
}

// swap swaps tho nodes
func (q *PriorityQueue) swap(index1, index2 int) {
	q.nodes[index1], q.nodes[index2] = q.nodes[index2], q.nodes[index1]
}

// heapify restores the property of ordering in the entire subtree, the root of which is the rootIndex
func (q *PriorityQueue) heapify(rootIndex int) {
	largest := rootIndex

	indexLeftChild, exist := q.indexChildLeft(rootIndex)
	if exist && (q.nodes[indexLeftChild].priority > q.nodes[largest].priority) {
		largest = indexLeftChild
	}

	indexRightChild, exist := q.indexChildRight(rootIndex)
	if exist && (q.nodes[indexRightChild].priority > q.nodes[largest].priority) {
		largest = indexRightChild
	}

	if rootIndex != largest {
		q.swap(rootIndex, largest)
		q.heapify(largest)
	}
}
