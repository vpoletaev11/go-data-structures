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
func (b *BinarySearchTree) Insert(val int) {
	if b.root == nil {
		b.root = &nodeBST{data: val, leftChild: nil, rightChild: nil}
		return
	}
	parent := b.root
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
func (b *BinarySearchTree) Find(val int) bool {
	if b.root == nil {
		return false
	}
	parent := b.root
	for {
		switch {
		case parent.data < val:
			if parent.rightChild == nil {
				return false
			}
			parent = parent.rightChild

		case parent.data > val:
			if parent.leftChild == nil {
				return false
			}
			parent = parent.leftChild

		case parent.data == val:
			return true
		}
	}
}

// Remove removes value from tree
func (b *BinarySearchTree) Remove(val int) {
	if b.root == nil {
		return
	}

	if b.root.data == val {
		if b.root.noChildren() {
			b.root = nil
			return
		}

		if b.root.hasTwoChild() {
			b.root.data = b.root.cutLeftmostNodeValueFRS()
			return
		}

		b.root.removeNodeWithOneChild()
		return
	}

	parent := b.root
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

// cutLeftmostNodeFRS obtains leftmost node data and deletes leftmost node From Right Subtree
func (n *nodeBST) cutLeftmostNodeValueFRS() int {
	if n.rightChild.leftChild == nil { // case when right subtree haven't left child
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
			if root.leftChild.rightChild == nil {
				root.leftChild = nil
				return data
			}
			root.leftChild = root.leftChild.rightChild
			return data
		}
		root = root.leftChild
	}
}
