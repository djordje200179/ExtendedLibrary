package redblacktree

func (tree *Tree[K, V]) fixInsert(node *Node[K, V]) {
	if node.parent == nil {
		node.color = black

		return
	} else if node.parent.color == black {
		return
	}

	uncle := node.parent.Sibling()
	if uncle != nil && uncle.color == red {
		node.parent.color = black
		uncle.color = black
		node.parent.parent.color = red
		tree.fixInsert(node.parent.parent)

		return
	}

	grandparent := node.parent.parent
	if node == node.parent.right && node.parent == grandparent.left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == grandparent.right {
		tree.rotateRight(node.parent)
		node = node.right
	}

	node.parent.color = black
	grandparent.color = red

	if node == node.parent.left && node.parent == grandparent.left {
		tree.rotateRight(grandparent)
	} else {
		tree.rotateLeft(grandparent)
	}
}

func (tree *Tree[K, V]) rotateLeft(node *Node[K, V]) {
	leftNode := node
	rightNode := node.right
	middleNode := rightNode.left

	rightNode.left = leftNode
	leftNode.right = middleNode
	if middleNode != nil {
		middleNode.parent = leftNode
	}

	parent := leftNode.parent

	leftNode.parent = rightNode
	rightNode.parent = parent

	if parent != nil {
		if parent.left == leftNode {
			parent.left = rightNode
		} else {
			parent.right = rightNode
		}
	} else {
		tree.root = rightNode
	}
}

func (tree *Tree[K, V]) rotateRight(node *Node[K, V]) {
	leftNode := node.left
	rightNode := node
	middleNode := leftNode.right

	leftNode.right = rightNode
	rightNode.left = middleNode

	leftNode.parent = rightNode.parent
	rightNode.parent = leftNode
	if middleNode != nil {
		middleNode.parent = rightNode
	}

	parent := rightNode.parent

	rightNode.parent = leftNode
	leftNode.parent = parent

	if parent != nil {
		if parent.left == rightNode {
			parent.left = leftNode
		} else {
			parent.right = leftNode
		}
	} else {
		tree.root = leftNode
	}
}
