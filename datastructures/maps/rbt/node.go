package rbt

type color bool

const (
	red   color = false
	black color = true
)

// Node is a node of a red-black tree.
// It should not be created directly.
type Node[K, V any] struct {
	key   K
	Value V // Value is the value stored in the node.

	color color

	leftChild  *Node[K, V]
	rightChild *Node[K, V]
	parent     *Node[K, V]
}

// Key returns the key of the node.
func (node *Node[K, V]) Key() K {
	return node.key
}

// LeftChild returns the left child of the node.
func (node *Node[K, V]) LeftChild() *Node[K, V] {
	return node.leftChild
}

// RightChild returns the right child of the node.
func (node *Node[K, V]) RightChild() *Node[K, V] {
	return node.rightChild
}

// Parent returns the parent of the node.
func (node *Node[K, V]) Parent() *Node[K, V] {
	return node.parent
}

// Sibling returns the sibling of the node.
// If the node is the root, it returns nil.
func (node *Node[K, V]) Sibling() *Node[K, V] {
	if node.parent == nil {
		return nil
	}

	if node.parent.leftChild == node {
		return node.parent.rightChild
	} else {
		return node.parent.leftChild
	}
}

// Prev returns the previous node in the tree
// in the order of the keys.
func (node *Node[K, V]) Prev() *Node[K, V] {
	if node.leftChild != nil {
		return node.leftChild.Max()
	}

	for prev, curr := node.parent, node; prev != nil && curr == prev.leftChild; curr, prev = prev, prev.parent {
		if curr != prev.leftChild {
			return prev
		}
	}

	return nil
}

// Next returns the next node in the tree
// in the order of the keys.
func (node *Node[K, V]) Next() *Node[K, V] {
	if node.rightChild != nil {
		return node.rightChild.Min()
	}

	for prev, curr := node.parent, node; prev != nil; curr, prev = prev, prev.parent {
		if curr != prev.rightChild {
			return prev
		}
	}

	return nil
}

// Min returns the minimum node in the subtree rooted at the node.
func (node *Node[K, V]) Min() *Node[K, V] {
	for curr := node; curr != nil; curr = curr.leftChild {
		if curr.leftChild == nil {
			return curr
		}
	}

	return nil
}

// Max returns the maximum node in the subtree rooted at the node.
func (node *Node[K, V]) Max() *Node[K, V] {
	for curr := node; curr != nil; curr = curr.rightChild {
		if curr.rightChild == nil {
			return curr
		}
	}

	return nil
}

// Clone returns a clone of the node.
// The clone has the same key, value, and color as the node.
// The clone does not have any links to other nodes.
func (node *Node[K, V]) Clone() *Node[K, V] {
	return &Node[K, V]{
		key:   node.key,
		Value: node.Value,
		color: node.color,
	}
}

func nodeColor[K, V any](node *Node[K, V]) color {
	if node == nil {
		return black
	}

	return node.color
}
