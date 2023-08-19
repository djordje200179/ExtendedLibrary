package redblacktree

func (tree *Tree[K, V]) fixInsert(node *Node[K, V]) {
	if node.parent == nil {
		node.color = black

		return
	} else if node.parent.color == black {
		return
	}

	uncle := node.parent.Sibling()
	if nodeColor(uncle) == red {
		node.parent.color = black
		uncle.color = black
		node.parent.parent.color = red
		tree.fixInsert(node.parent.parent)

		return
	}

	grandparent := node.parent.parent
	if node == node.parent.rightChild && node.parent == grandparent.leftChild {
		tree.rotateLeft(node.parent)
		node = node.leftChild
	} else if node == node.parent.leftChild && node.parent == grandparent.rightChild {
		tree.rotateRight(node.parent)
		node = node.rightChild
	}

	node.parent.color = black
	grandparent.color = red

	if node == node.parent.leftChild && node.parent == grandparent.leftChild {
		tree.rotateRight(grandparent)
	} else {
		tree.rotateLeft(grandparent)
	}
}

func (tree *Tree[K, V]) fixRemove(node *Node[K, V]) {
	for node != tree.root && nodeColor(node) == black {
		if node == node.parent.leftChild {
			sibling := node.parent.rightChild

			if nodeColor(sibling) == red {
				sibling.color = black
				node.parent.color = red
				tree.rotateLeft(node.parent)
				sibling = node.parent.rightChild
			}

			if nodeColor(sibling.leftChild) == black && nodeColor(sibling.rightChild) == black {
				sibling.color = red
				node = node.parent
			} else {
				if nodeColor(sibling.rightChild) == black {
					sibling.leftChild.color = black
					sibling.color = red
					tree.rotateRight(sibling)
					sibling = node.parent.rightChild
				}

				sibling.color = node.parent.color
				node.parent.color = black
				sibling.rightChild.color = black
				tree.rotateLeft(node.parent)
				node = tree.root
			}
		} else {
			sibling := node.parent.leftChild

			if nodeColor(sibling) == red {
				sibling.color = black
				node.parent.color = red
				tree.rotateRight(node.parent)
				sibling = node.parent.leftChild
			}

			if nodeColor(sibling.leftChild) == black && nodeColor(sibling.rightChild) == black {
				sibling.color = red
				node = node.parent
			} else {
				if nodeColor(sibling.leftChild) == black {
					sibling.rightChild.color = black
					sibling.color = red
					tree.rotateLeft(sibling)
					sibling = node.parent.leftChild
				}

				sibling.color = node.parent.color
				node.parent.color = black
				sibling.leftChild.color = black
				tree.rotateRight(node.parent)
				node = tree.root
			}
		}
	}

	node.color = black
}

func (tree *Tree[K, V]) rotateLeft(node *Node[K, V]) {
	rightNode := node.rightChild
	node.rightChild = rightNode.leftChild
	if rightNode.leftChild != nil {
		rightNode.leftChild.parent = node
	}

	rightNode.parent = node.parent
	if node.parent == nil {
		tree.root = rightNode
	} else if node == node.parent.leftChild {
		node.parent.leftChild = rightNode
	} else {
		node.parent.rightChild = rightNode
	}

	rightNode.leftChild = node
	node.parent = rightNode
}

func (tree *Tree[K, V]) rotateRight(node *Node[K, V]) {
	leftNode := node.leftChild
	node.leftChild = leftNode.rightChild
	if leftNode.rightChild != nil {
		leftNode.rightChild.parent = node
	}

	leftNode.parent = node.parent
	if node.parent == nil {
		tree.root = leftNode
	} else if node == node.parent.rightChild {
		node.parent.rightChild = leftNode
	} else {
		node.parent.leftChild = leftNode
	}
	leftNode.rightChild = node
	node.parent = leftNode
}
