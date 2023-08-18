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

func (tree *Tree[K, V]) fixRemove(node *Node[K, V]) {
	for node != tree.root && nodeColor(node) == black {
		if node == node.parent.left {
			sibling := node.parent.right

			if nodeColor(sibling) == red {
				sibling.color = black
				node.parent.color = red
				tree.rotateLeft(node.parent)
				sibling = node.parent.right
			}

			if nodeColor(sibling.left) == black && nodeColor(sibling.right) == black {
				sibling.color = red
				node = node.parent
			} else {
				if nodeColor(sibling.right) == black {
					sibling.left.color = black
					sibling.color = red
					tree.rotateRight(sibling)
					sibling = node.parent.right
				}

				sibling.color = node.parent.color
				node.parent.color = black
				sibling.right.color = black
				tree.rotateLeft(node.parent)
				node = tree.root
			}
		} else {
			sibling := node.parent.left

			if nodeColor(sibling) == red {
				sibling.color = black
				node.parent.color = red
				tree.rotateRight(node.parent)
				sibling = node.parent.left
			}

			if nodeColor(sibling.left) == black && nodeColor(sibling.right) == black {
				sibling.color = red
				node = node.parent
			} else {
				if nodeColor(sibling.left) == black {
					sibling.right.color = black
					sibling.color = red
					tree.rotateLeft(sibling)
					sibling = node.parent.left
				}

				sibling.color = node.parent.color
				node.parent.color = black
				sibling.left.color = black
				tree.rotateRight(node.parent)
				node = tree.root
			}
		}
	}

	node.color = black
}

func (tree *Tree[K, V]) rotateLeft(node *Node[K, V]) {
	rightNode := node.right
	node.right = rightNode.left
	if rightNode.left != nil {
		rightNode.left.parent = node
	}

	rightNode.parent = node.parent
	if node.parent == nil {
		tree.root = rightNode
	} else if node == node.parent.left {
		node.parent.left = rightNode
	} else {
		node.parent.right = rightNode
	}

	rightNode.left = node
	node.parent = rightNode
}

func (tree *Tree[K, V]) rotateRight(node *Node[K, V]) {
	leftNode := node.left
	node.left = leftNode.right
	if leftNode.right != nil {
		leftNode.right.parent = node
	}

	leftNode.parent = node.parent
	if node.parent == nil {
		tree.root = leftNode
	} else if node == node.parent.right {
		node.parent.right = leftNode
	} else {
		node.parent.left = leftNode
	}
	leftNode.right = node
	node.parent = leftNode
}
