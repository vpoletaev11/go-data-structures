package tree

// nodeAVL - node of AVL tree
type nodeAVL struct {
	data       int
	height     int
	parent     *nodeAVL
	leftChild  *nodeAVL
	rightChild *nodeAVL
}

// AVLTree - balanced binary search tree, where for each of its vertices the height of two subtrees differs by no more than 1
type AVLTree struct {
	root *nodeAVL
}

// Insert adds value into tree
func (a *AVLTree) Insert(val int) {
	if a.root == nil {
		a.root = &nodeAVL{
			data:       val,
			height:     1,
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
		}
		return
	}

	parent := a.root
	for {
		switch {
		case parent.data < val:
			if parent.rightChild == nil {
				parent.rightChild = &nodeAVL{
					data:       val,
					height:     1,
					parent:     parent,
					leftChild:  nil,
					rightChild: nil,
				}

				parent.fixHeight()
				unbalancedNode, balanceFactor := parent.findUnbalancedNode()
				balanceNode(unbalancedNode, balanceFactor)
				return
			}
			parent = parent.rightChild

		case parent.data > val:
			if parent.leftChild == nil {
				parent.leftChild = &nodeAVL{
					data:       val,
					height:     1,
					parent:     parent,
					leftChild:  nil,
					rightChild: nil,
				}

				parent.fixHeight()
				unbalancedNode, balanceFactor := parent.findUnbalancedNode()
				balanceNode(unbalancedNode, balanceFactor)
				return
			}
			parent = parent.leftChild

		case parent.data == val:
			return
		}
	}
}

// Find checks if value exists in tree
func (a *AVLTree) Find(val int) bool {
	if a.root == nil {
		return false
	}
	parent := a.root
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
func (a *AVLTree) Remove(val int) {
	if a.root == nil {
		return
	}

	if a.root.data == val {
		if a.root.noChildren() {
			a.root = nil
			return
		}

		if a.root.hasTwoChildren() {
			data, cutNodeParent := a.root.cutLeftmostNodeValueFRS()
			a.root.data = data

			unbalancedNode, balanceFactor := cutNodeParent.findUnbalancedNode()
			balanceNode(unbalancedNode, balanceFactor)
			return
		}

		removeNodeWithOneChild(a.root)
		return
	}

	parent := a.root
	for {
		switch {
		case parent.data < val:
			if parent.rightChild == nil {
				return
			}

			if parent.rightChild.data == val {
				if parent.rightChild.noChildren() {
					parent.rightChild = nil

					parent.fixHeight()
					unbalancedNode, balanceFactor := parent.findUnbalancedNode()
					balanceNode(unbalancedNode, balanceFactor)
					return
				}

				if parent.rightChild.hasTwoChildren() {
					data, cutNodeParent := parent.rightChild.cutLeftmostNodeValueFRS()
					parent.rightChild.data = data

					unbalancedNode, balanceFactor := cutNodeParent.findUnbalancedNode()
					balanceNode(unbalancedNode, balanceFactor)
					return
				}

				removeNodeWithOneChild(parent.rightChild)

				parent.fixHeight()
				unbalancedNode, balanceFactor := parent.findUnbalancedNode()
				balanceNode(unbalancedNode, balanceFactor)
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

					parent.fixHeight()
					unbalancedNode, balanceFactor := parent.findUnbalancedNode()
					balanceNode(unbalancedNode, balanceFactor)
					return
				}

				if parent.leftChild.hasTwoChildren() {
					data, cutNodeParent := parent.leftChild.cutLeftmostNodeValueFRS()
					parent.leftChild.data = data

					unbalancedNode, balanceFactor := cutNodeParent.findUnbalancedNode()
					balanceNode(unbalancedNode, balanceFactor)
					return
				}

				removeNodeWithOneChild(parent.leftChild)

				parent.fixHeight()
				unbalancedNode, balanceFactor := parent.findUnbalancedNode()
				balanceNode(unbalancedNode, balanceFactor)
				return
			}
			parent = parent.leftChild
		}
	}
}

// fixHeight recalculates height of the node and node ancestors
func (n *nodeAVL) fixHeight() {
	if n == nil {
		return
	}

	switch {
	case n.rightChild == nil && n.leftChild == nil:
		n.height = 1
		n.parent.fixHeight()

	case n.leftChild == nil:
		n.height = n.rightChild.height + 1
		n.parent.fixHeight()

	case n.rightChild == nil:
		n.height = n.leftChild.height + 1
		n.parent.fixHeight()

	case n.leftChild.height > n.rightChild.height:
		n.height = n.leftChild.height + 1
		n.parent.fixHeight()

	default:
		n.height = n.rightChild.height + 1
		n.parent.fixHeight()
	}
}

// findUnbalancedNode checks if inputted node and it ancestors are unbalanced.
// In case when node are unbalanced will be returned unbalanced node and it balance factor.
// In case when unbalanced node not found will be returned nil and 0.
func (n *nodeAVL) findUnbalancedNode() (unbalancedNode *nodeAVL, balanceFactor int) {
	node := n
	for {
		switch {
		case node.leftChild == nil && node.rightChild == nil:
			balanceFactor = 0

		case node.leftChild == nil:
			balanceFactor = 0 - node.rightChild.height

		case node.rightChild == nil:
			balanceFactor = node.leftChild.height - 0

		default:
			balanceFactor = node.leftChild.height - node.rightChild.height
		}

		if balanceFactor > 1 {
			return node, balanceFactor
		}
		if balanceFactor < -1 {
			return node, balanceFactor
		}

		if node.parent == nil {
			return nil, 0
		}
		node = node.parent
	}
}

// rotationR makes right rotation
func (n *nodeAVL) rotationR() {
	n.data, n.leftChild.data = n.leftChild.data, n.data

	n.rightChild = &nodeAVL{
		data:       n.leftChild.data,
		parent:     n,
		leftChild:  n.leftChild.rightChild,
		rightChild: n.rightChild,
	}
	if n.rightChild.leftChild != nil {
		n.rightChild.leftChild.parent = n.rightChild
	}
	if n.rightChild.rightChild != nil {
		n.rightChild.rightChild.parent = n.rightChild
	}

	n.leftChild = n.leftChild.leftChild
	if n.leftChild != nil {
		n.leftChild.parent = n
	}

	n.rightChild.fixHeight()
}

// rotationL makes left rotation
func (n *nodeAVL) rotationL() {
	n.data, n.rightChild.data = n.rightChild.data, n.data

	n.leftChild = &nodeAVL{
		data:       n.rightChild.data,
		parent:     n,
		leftChild:  n.leftChild,
		rightChild: n.rightChild.leftChild,
	}
	if n.leftChild.rightChild != nil {
		n.leftChild.rightChild.parent = n.leftChild
	}

	n.rightChild = n.rightChild.rightChild
	if n.rightChild != nil {
		n.rightChild.parent = n
	}
	if n.leftChild.leftChild != nil {
		n.leftChild.leftChild.parent = n.leftChild
	}

	n.leftChild.fixHeight()
}

// rotationRL makes right-left (big left) rotation
func (n *nodeAVL) rotationRL() {
	n.rightChild.rotationR()
	n.rotationL()
}

// rotationLR makes left-right (big right) rotation
func (n *nodeAVL) rotationLR() {
	n.leftChild.rotationL()
	n.rotationR()
}

// balanceNode balances node by using one of rotaions if it needed
func balanceNode(unbalancedNode *nodeAVL, balanceFactor int) {
	if unbalancedNode == nil {
		return
	}

	if balanceFactor == 2 {
		if unbalancedNode.leftChild.leftChild == nil {
			unbalancedNode.rotationLR()
			return
		}

		if unbalancedNode.leftChild.rightChild != nil {
			if unbalancedNode.leftChild.rightChild.height > unbalancedNode.leftChild.leftChild.height {
				unbalancedNode.rotationLR()
				return
			}
		}

		unbalancedNode.rotationR()
		return
	}

	if balanceFactor == -2 {
		if unbalancedNode.rightChild.rightChild == nil {
			unbalancedNode.rotationRL()
			return
		}

		if unbalancedNode.rightChild.leftChild != nil {
			if unbalancedNode.rightChild.leftChild.height > unbalancedNode.rightChild.rightChild.height {
				unbalancedNode.rotationRL()
				return
			}
		}

		unbalancedNode.rotationL()
	}
}

// noChildren checks that the node has no children
func (n *nodeAVL) noChildren() bool {
	if n.leftChild == nil && n.rightChild == nil {
		return true
	}
	return false
}

// hasTwoChild check that the node has two child
func (n *nodeAVL) hasTwoChildren() bool {
	if n.leftChild != nil && n.rightChild != nil {
		return true
	}
	return false
}

// removeNodeWithOneChild removes node who have only one child
func removeNodeWithOneChild(n *nodeAVL) {
	if n.leftChild != nil {
		parent := n.parent
		*n = *n.leftChild
		n.parent = parent
		return
	}
	if n.rightChild != nil {
		parent := n.parent
		*n = *n.rightChild
		n.parent = parent
		return
	}
}

// cutLeftmostNodeFRS obtains leftmost node data and deletes leftmost node From Right Subtree;
// return leftmost node data and it parent.
func (n *nodeAVL) cutLeftmostNodeValueFRS() (data int, parent *nodeAVL) {
	if n.rightChild.leftChild == nil { // case when right subtree haven't left child
		data = n.rightChild.data
		parent = n.rightChild.parent

		if n.rightChild.rightChild == nil {
			n.rightChild = nil

			n.fixHeight()
			return data, parent
		}

		n.rightChild = n.rightChild.rightChild
		n.rightChild.parent = n
		n.rightChild.fixHeight()
		return data, parent
	}

	root := n.rightChild
	for {
		if root.leftChild.leftChild == nil {
			data = root.leftChild.data
			parent = root

			if root.leftChild.rightChild == nil {
				root.leftChild = nil
				root.fixHeight()
				return data, parent
			}

			root.leftChild = root.leftChild.rightChild
			root.leftChild.parent = root
			root.leftChild.fixHeight()
			return data, parent
		}
		root = root.leftChild
	}
}
