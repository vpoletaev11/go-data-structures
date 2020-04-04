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
				return
			}
			parent = parent.leftChild

		case parent.data == val:
			return
		}
	}
}
