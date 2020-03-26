package tree

// nodeBST - node of binary search tree
type nodeBST struct {
	data       int
	leftChild  *nodeBST
	rightChild *nodeBST
}

// BinarySearchTree - binary tree where:
// 1) Left and right subtrees are binary search trees
// 2) All nodes of the left subtree of an random node X have value ​​less than value of node X itself
// 3) All nodes of the right subtree of an random node X have value ​​>= value of node X itself
type BinarySearchTree struct {
	root *nodeBST
}

// Insert adds value into BinarySearchTree
func (t *BinarySearchTree) Insert(val int) {
	if t.root == nil {
		t.root = &nodeBST{data: val, leftChild: nil, rightChild: nil}
		return
	}
	parent := t.root
	for {
		switch {
		case parent.data < val:
			if parent.rightChild == nil {
				parent.rightChild = &nodeBST{data: val, leftChild: nil, rightChild: nil}
				return
			}
			parent = parent.rightChild

		case parent.data > val:
			if parent.leftChild == nil {
				parent.leftChild = &nodeBST{data: val, leftChild: nil, rightChild: nil}
				return
			}
			parent = parent.leftChild

		case parent.data == val:
			return
		}
	}
}

// Find checks if value exists in tree
func (t *BinarySearchTree) Find(val int) bool {
	if t.root == nil {
		return false
	}
	parent := t.root
	for {
		switch {
		case parent.data < val:
			if parent.rightChild == nil {
				return false
			}
			if parent.rightChild.data == val {
				return true
			}
			parent = parent.rightChild

		case parent.data > val:
			if parent.leftChild == nil {
				return false
			}
			if parent.leftChild.data == val {
				return true
			}
			parent = parent.leftChild

		case parent.data == val:
			return true
		}
	}
}

// Remove removes value from tree
func (t *BinarySearchTree) Remove(val int) {
	if t.root == nil {
		return
	}

	if t.root.data == val {
		if t.root.noChildren() {
			t.root = nil
			return
		}

		if t.root.hasTwoChild() {
			t.root.data = t.root.cutLeftmostNodeValueFRS()
			return
		}

		t.root.removeNodeWithOneChild()
		return
	}

	parent := t.root
	for {
		switch {
		case parent.data < val:
			if parent.rightChild == nil {
				return
			}

			if parent.rightChild.data == val {
				if parent.rightChild.noChildren() {
					parent.rightChild = nil
					return
				}

				if parent.rightChild.hasTwoChild() {
					parent.rightChild.data = parent.rightChild.cutLeftmostNodeValueFRS()
					return
				}

				parent.rightChild.removeNodeWithOneChild()
				return
			}
			parent = parent.rightChild

		case parent.data > val:
			if parent.leftChild == nil {
				return
			}
			if parent.leftChild.data == val {
				if parent.leftChild.noChildren() {
					parent.leftChild = nil
					return
				}

				if parent.leftChild.hasTwoChild() {
					parent.leftChild.data = parent.leftChild.cutLeftmostNodeValueFRS()
					return
				}

				parent.leftChild.removeNodeWithOneChild()
				return
			}
			parent = parent.leftChild
		}
	}
}

// noChildren checks that the node has no children
func (n *nodeBST) noChildren() bool {
	if n.leftChild == nil && n.rightChild == nil {
		return true
	}
	return false
}

// hasTwoChild check that the node has two child
func (n *nodeBST) hasTwoChild() bool {
	if n.leftChild != nil && n.rightChild != nil {
		return true
	}
	return false
}

// removeNodeWithOneChild removes node who have only one child
func (n *nodeBST) removeNodeWithOneChild() {
	if n.leftChild != nil {
		*n = *n.leftChild
		return
	}
	if n.rightChild != nil {
		*n = *n.rightChild
		return
	}
}

// cutLeftmostNodeFRS obtains leftmost node data and deletes leftmost node From Right Subtree of node n
func (n *nodeBST) cutLeftmostNodeValueFRS() int {
	if n.rightChild.leftChild == nil {
		data := n.rightChild.data
		if n.rightChild.rightChild == nil {
			n.rightChild = nil
			return data
		}
		n.rightChild = n.rightChild.rightChild
		return data
	}

	root := n.rightChild
	for {
		if root.leftChild.leftChild == nil {
			data := root.leftChild.data
			root.leftChild = nil
			return data
		}
		root = root.leftChild
	}
}
