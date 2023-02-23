package binarysearchtree

type Node[K comparable, V any] struct {
	key   K
	Value V

	left, right, parent *Node[K, V]
}

func (node *Node[K, V]) Key() K {
	return node.key
}

func (node *Node[K, V]) Left() *Node[K, V] {
	return node.left
}

func (node *Node[K, V]) Right() *Node[K, V] {
	return node.right
}

func (node *Node[K, V]) Parent() *Node[K, V] {
	return node.parent
}

func (node *Node[K, V]) Clone() *Node[K, V] {
	cloned := &Node[K, V]{
		key:   node.key,
		Value: node.Value,
	}

	return cloned
}

func (node *Node[K, V]) Prev() *Node[K, V] {
	if node.left != nil {
		return node.left.Max()
	}

	for prev, curr := node.parent, node; prev != nil && curr == prev.left; curr, prev = prev, prev.parent {
		if curr != prev.left {
			return prev
		}
	}

	return nil
}

func (node *Node[K, V]) Next() *Node[K, V] {
	if node.right != nil {
		return node.right.Min()
	}

	for prev, curr := node.parent, node; prev != nil; curr, prev = prev, prev.parent {
		if curr != prev.right {
			return prev
		}
	}

	return nil
}

func (node *Node[K, V]) locationInParent() **Node[K, V] {
	if node.parent == nil {
		return nil
	}

	if node.parent.left == node {
		return &node.parent.left
	} else {
		return &node.parent.right
	}
}

func (node *Node[K, V]) Min() *Node[K, V] {
	for curr := node; curr != nil; curr = curr.left {
		if curr.left == nil {
			return curr
		}
	}

	return nil
}

func (node *Node[K, V]) Max() *Node[K, V] {
	for curr := node; curr != nil; curr = curr.right {
		if curr.right == nil {
			return curr
		}
	}

	return nil
}
