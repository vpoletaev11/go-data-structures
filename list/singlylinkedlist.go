package list

/*
	SINGLY LINKED LIST:

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
	BenchmarkInsertHeadSinglyList-8   	 1580721	       750 ns/op	     320 B/op	      10 allocs/op
	BenchmarkInsertSinglyList-8       	 1488368	       806 ns/op	     320 B/op	      10 allocs/op
	BenchmarkInsertTailSinglyList-8   	 1617968	       732 ns/op	     320 B/op	      10 allocs/op

	BenchmarkGetHeadSinglyList-8      	13920870	        86.5 ns/op	       0 B/op	       0 allocs/op
	BenchmarkGetSinglyList-8          	 6337293	       190 ns/op	       0 B/op	       0 allocs/op
	BenchmarkGetTailSinglyList-8      	 7814540	       154 ns/op	       0 B/op	       0 allocs/op
	
	BenchmarkPeekHeadSinglyList-8     	334253223	         3.08 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPeekSinglyList-8         	31299858	        40.6 ns/op	       0 B/op	       0 allocs/op
	BenchmarkPeekTailSinglyList-8     	21066404	        56.1 ns/op	       0 B/op	       0 allocs/op
*/

// nodeSL - node of singly linked list
type nodeSL struct {
	data     string  // Stored data
	nextNode *nodeSL // Pointer to next node
}

// SinglyLinkedList -  data structure consisting of a collection of nodes,
// where every node contain pointer to next node, which together represent a sequence.
type SinglyLinkedList struct {
	head *nodeSL // First element in linked list
	len  int     // Length of the linked list
}

// Len returns length of list
func (s *SinglyLinkedList) Len() int {
	return s.len
}

// InsertHead adds value in begin of list
func (s *SinglyLinkedList) InsertHead(val string) {
	if s.head == nil {
		s.head = &nodeSL{val, nil}
		s.len++
		return
	}

	s.head = &nodeSL{val, s.head}
	s.len++
}

// Insert adds value in list by index.
// If on index already exists value, it and all subsequent values will be moved by 1, and in its place will be inserted new value
func (s *SinglyLinkedList) Insert(index int, val string) (ok bool) {
	if index == 0 {
		s.InsertHead(val)
		return true
	}

	prevNode := s.nodeByIndex(index - 1) // Find node before of candidate node to be inserted
	if prevNode == nil {
		return false
	}
	prevNode.nextNode = &nodeSL{val, prevNode.nextNode} // prevNode now pointing on inserted node which pointing to node who was here before
	s.len++
	return true
}

// InsertTail adds value in end of list
func (s *SinglyLinkedList) InsertTail(val string) {
	if s.head == nil {
		s.head = &nodeSL{val, nil}
		s.len++
		return
	}

	lastNode := s.findLastNode()
	lastNode.nextNode = &nodeSL{val, nil}
	s.len++
}

// GetHead obtains and deletes element from begin of list
func (s *SinglyLinkedList) GetHead() (val string, ok bool) {
	if s.head == nil {
		return "", false
	}

	val = s.head.data
	s.head = s.head.nextNode
	s.len--
	return val, true
}

// Get obtains and deletes element by index from list
func (s *SinglyLinkedList) Get(index int) (val string, ok bool) {
	if index == 0 {
		return s.GetHead()
	}

	prevNode := s.nodeByIndex(index - 1) // Find node before of candidate node to be obtained and deleted
	if prevNode == nil || prevNode.nextNode == nil {
		return "", false
	}
	val = prevNode.nextNode.data                   // Obtain node data
	prevNode.nextNode = prevNode.nextNode.nextNode // prevNode now pointing to next node after deleted node.
	s.len--
	return val, true
}

// GetTail obtains and deletes element from end of SinglyLinkedList
func (s *SinglyLinkedList) GetTail() (val string, ok bool) {
	if s.head == nil {
		return "", false
	}

	p1, p2 := s.head, s.head // Initialize variables that stores pointers to the second-to-last and last nodes respectively
	for {
		if p2.nextNode == nil {
			break
		}

		p1 = p2
		p2 = p2.nextNode
	}

	if p1 == p2 { // checking if in list is only 1 element
		s.head = nil
	}

	val = p2.data
	p1.nextNode = nil // Second-to-last node points to nil i.e now it last node
	s.len--
	return val, true
}

// PeekHead obtains element from begin of list
func (s *SinglyLinkedList) PeekHead() (val string, ok bool) {
	if s.head == nil {
		return "", false
	}
	return s.head.data, true
}

// Peek obtains element by index from list
func (s *SinglyLinkedList) Peek(index int) (val string, ok bool) {
	node := s.nodeByIndex(index)
	if node == nil {
		return "", false
	}
	return node.data, true
}

// PeekTail obtains element from end of list
func (s *SinglyLinkedList) PeekTail() (val string, ok bool) {
	lastNode := s.findLastNode()
	if lastNode == nil {
		return "", false
	}
	return lastNode.data, true
}

// findLastNode finds last element in list
func (s *SinglyLinkedList) findLastNode() *nodeSL {
	node := s.head
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
func (s *SinglyLinkedList) nodeByIndex(index int) *nodeSL {
	if index < 0 {
		return nil
	}

	node := s.head
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
