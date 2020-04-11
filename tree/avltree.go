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

// Insert adds value into AVL tree
func (b *AVLTree) Insert(val int) {
	if b.root == nil {
		b.root = &nodeAVL{
			data:       val,
			height:     1,
			parent:     nil,
			leftChild:  nil,
			rightChild: nil,
		}
		return
	}

	parent := b.root
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
				return
			}
			parent = parent.leftChild

		case parent.data == val:
			return
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

func (n *nodeAVL) findUnbalancedNode() *nodeAVL {
	node := n
	var balanceFactor int
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
			return node
		}
		if balanceFactor < -1 {
			return node
		}

		if node.parent == nil {
			return nil
		}
		node = node.parent
	}
}

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

func (n *nodeAVL) rotationRL() {
	n.rightChild.rotationR()
	n.rotationL()
}

func (n *nodeAVL) rotationLR() {
	n.leftChild.rotationL()
	n.rotationR()
}
