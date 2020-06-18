package list

/*
	DOUBLY LINKED LIST:

	Len()
	Best:    O(1)
	Average: O(1)
	Worst:	 O(1)

	InsertHead()         Insert()             InsertTail()
	Best:    O(1)        Best:    O(1)        Best:    O(n)
	Average: O(1)        Average: O(n)        Average: O(n)
	Worst:   O(1)        Worst:   O(n)        Worst:   O(n)


	GetHead()            Get()                GetTail()
	Best:    O(1)        Best:    O(1)        Best:    O(n)
	Average: O(1)        Average: O(n)        Average: O(n)
	Worst:   O(1)        Worst:   O(n)        Worst:   O(n)


	PeekHead()           Peek()               PeekTail()
	Best:    O(1)        Best:    O(1)        Best:    O(n)
	Average: O(1)        Average: O(n)        Average: O(n)
	Worst:   O(1)        Worst:   O(n)        Worst:   O(n)


	goos: linux
	goarch: amd64
	BenchmarkInsertHeadDoublyList-8          1000000               550 ns/op             320 B/op         10 allocs/op
	BenchmarkInsertDoublyList-8              1000000               684 ns/op             320 B/op         10 allocs/op
	BenchmarkInsertTailDoublyList-8          1000000               747 ns/op             320 B/op         10 allocs/op

	BenchmarkGetHeadDoublyList-8             1000000                86.6 ns/op             0 B/op          0 allocs/op
	BenchmarkGetDoublyList-8                 1000000               192 ns/op               0 B/op          0 allocs/op
	BenchmarkGetTailDoublyList-8             1000000               214 ns/op               0 B/op          0 allocs/op
	
	BenchmarkPeekHeadDoublyList-8            1000000                 7.55 ns/op            0 B/op          0 allocs/op
	BenchmarkPeekDoublyList-8                1000000                48.3 ns/op             0 B/op          0 allocs/op
	BenchmarkPeekTailDoublyList-8            1000000                71.9 ns/op             0 B/op          0 allocs/op
*/

// nodeDL - node of doubly linked list
type nodeDL struct {
	data     string  // Stored data
	prevNode *nodeDL // Pointer to previous node
	nextNode *nodeDL // Pointer to next node
}

// DoublyLinkedList -  data structure consisting of a collection of nodes,
// where every node contain pointer to previous and next node, which together represent a sequence.
type DoublyLinkedList struct {
	head *nodeDL // First element in linked list
	len  int     // Length of the linked list
}

// Len returns length of list
func (d *DoublyLinkedList) Len() int {
	return d.len
}

// InsertHead adds value in begin of list
func (d *DoublyLinkedList) InsertHead(val string) {
	if d.head == nil {
		d.head = &nodeDL{data: val, prevNode: nil, nextNode: nil}
		d.len++
		return
	}

	nextNode := d.head
	d.head = &nodeDL{data: val, prevNode: nil, nextNode: nextNode} // Move head node in second place and insert new node on its place
	nextNode.prevNode = d.head                                     // Second node prevNode field points to head node
	d.len++
}

// Insert adds value in list by index.
// If on index already exists value, it and all subsequent values will be moved by 1, and in its place will be inserted new value
func (d *DoublyLinkedList) Insert(index int, val string) (ok bool) {
	if index == 0 {
		d.InsertHead(val)
		return true
	}

	prevNode := d.nodeByIndex(index - 1) // Find node before of candidate node to be inserted
	if prevNode == nil {
		return false
	}
	if prevNode.nextNode != nil {
		nextNode := prevNode.nextNode
		newNode := &nodeDL{data: val, prevNode: prevNode, nextNode: nextNode}

		prevNode.nextNode = newNode
		nextNode.prevNode = newNode
		d.len++
		return true
	}
	prevNode.nextNode = &nodeDL{data: val, prevNode: prevNode, nextNode: nil}
	d.len++
	return true
}

// InsertTail adds value in end of list
func (d *DoublyLinkedList) InsertTail(val string) {
	lastNode := d.findLastNode()
	if lastNode == nil {
		d.head = &nodeDL{data: val, prevNode: nil, nextNode: nil}
		d.len++
		return
	}
	lastNode.nextNode = &nodeDL{data: val, prevNode: lastNode, nextNode: nil}
	d.len++
}

// GetHead obtains and deletes element from begin of list
func (d *DoublyLinkedList) GetHead() (val string, ok bool) {
	if d.head == nil {
		return "", false
	}

	val = d.head.data
	d.head = d.head.nextNode
	if d.head != nil {
		d.head.prevNode = nil
	}
	d.len--
	return val, true
}

// Get obtains and deletes element by index from list
func (d *DoublyLinkedList) Get(index int) (val string, ok bool) {
	if index == 0 {
		return d.GetHead()
	}

	node := d.nodeByIndex(index)
	if node == nil {
		return "", false
	}
	val = node.data

	node.prevNode.nextNode = node.nextNode // Previous node nextNode field now points to next node
	if node.nextNode != nil {
		node.nextNode.prevNode = node.prevNode // Next node prevNode field now points to previous node
	}
	d.len--

	return val, true
}

// GetTail obtains and deletes element from end of SinglyLinkedList
func (d *DoublyLinkedList) GetTail() (val string, ok bool) {
	lastNode := d.findLastNode()
	if lastNode == nil {
		return "", false
	}

	val = lastNode.data
	if lastNode.prevNode == nil { // Check if lastNode is only one in list
		d.head = nil
		d.len--
		return val, true
	}
	lastNode.prevNode.nextNode = nil // Second-to-last node now points to nil
	d.len--
	return val, true
}

// PeekHead obtains element from begin of list
func (d *DoublyLinkedList) PeekHead() (val string, ok bool) {
	if d.head == nil {
		return "", false
	}
	return d.head.data, true
}

// Peek obtains element by index from list
func (d *DoublyLinkedList) Peek(index int) (val string, ok bool) {
	node := d.nodeByIndex(index)
	if node == nil {
		return "", false
	}
	return node.data, true
}

// PeekTail obtains element from end of list
func (d *DoublyLinkedList) PeekTail() (val string, ok bool) {
	lastNode := d.findLastNode()
	if lastNode == nil {
		return "", false
	}
	return lastNode.data, true
}

// findLastNode finds last element in list
func (d *DoublyLinkedList) findLastNode() *nodeDL {
	node := d.head
	if node == nil {
		return nil
	}

	for {
		if node.nextNode == nil {
			return node
		}
		node = node.nextNode
	}
}

// nodeByIndex returns node from list by index
func (d *DoublyLinkedList) nodeByIndex(index int) *nodeDL {
	if index < 0 {
		return nil
	}

	node := d.head
	if node == nil {
		return nil
	}

	for i := 0; i < index; i++ {
		if node.nextNode == nil { // Checking if last node is reached before of index node
			return nil
		}
		node = node.nextNode
	}
	return node
}
