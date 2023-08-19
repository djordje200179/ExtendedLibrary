package redblacktree

type color bool

const (
	red   color = false
	black color = true
)

type Node[K, V any] struct {
	key   K
	Value V

	color color

	leftChild  *Node[K, V]
	rightChild *Node[K, V]
	parent     *Node[K, V]
}

func (node *Node[K, V]) Key() K {
	return node.key
}

func (node *Node[K, V]) LeftChild() *Node[K, V] {
	return node.leftChild
}

func (node *Node[K, V]) RightChild() *Node[K, V] {
	return node.rightChild
}

func (node *Node[K, V]) Parent() *Node[K, V] {
	return node.parent
}

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

func (node *Node[K, V]) Min() *Node[K, V] {
	for curr := node; curr != nil; curr = curr.leftChild {
		if curr.leftChild == nil {
			return curr
		}
	}

	return nil
}

func (node *Node[K, V]) Max() *Node[K, V] {
	for curr := node; curr != nil; curr = curr.rightChild {
		if curr.rightChild == nil {
			return curr
		}
	}

	return nil
}

func nodeColor[K, V any](node *Node[K, V]) color {
	if node == nil {
		return black
	}

	return node.color
}
