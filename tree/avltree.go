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
		height:     n.leftChild.height,
		parent:     n,
		leftChild:  n.leftChild.rightChild,
		rightChild: n.rightChild,
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
		height:     n.rightChild.height,
		parent:     n,
		leftChild:  n.leftChild,
		rightChild: n.rightChild.leftChild,
	}

	n.rightChild = n.rightChild.rightChild
	if n.rightChild != nil {
		n.rightChild.parent = n
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

	switch {

	case balanceFactor == 2 && unbalancedNode.leftChild.leftChild == nil:
		unbalancedNode.rotationLR()

	case balanceFactor == -2 && unbalancedNode.rightChild.rightChild == nil:
		unbalancedNode.rotationRL()

	case balanceFactor == -2 && unbalancedNode.rightChild.leftChild.height > unbalancedNode.rightChild.rightChild.height:
		unbalancedNode.rotationRL()

	case balanceFactor == 2 && unbalancedNode.leftChild.rightChild.height > unbalancedNode.leftChild.leftChild.height:
		unbalancedNode.rotationLR()

	case balanceFactor == 2:
		unbalancedNode.rotationR()

	case balanceFactor == -2:
		unbalancedNode.rotationL()
	}
}
