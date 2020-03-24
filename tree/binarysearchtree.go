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
