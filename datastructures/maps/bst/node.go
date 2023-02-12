package bst

import "github.com/djordje200179/extendedlibrary/misc/functions/comparison"

type node[K comparable, V any] struct {
	key   K
	value V

	left, right, parent *node[K, V]
}

func (originalNode *node[K, V]) Clone() *node[K, V] {
	cloned := new(node[K, V])
	cloned.key = originalNode.key
	cloned.value = originalNode.value

	return cloned
}

func (node *node[K, V]) prev() *node[K, V] {
	if node.left != nil {
		return node.left.max()
	} else {
		for prev, curr := node.parent, node; prev != nil && curr == prev.left; curr, prev = prev, prev.parent {
			if curr != prev.left {
				return prev
			}
		}

		return nil
	}
}

func (node *node[K, V]) next() *node[K, V] {
	if node.right != nil {
		return node.right.min()
	} else {
		for prev, curr := node.parent, node; prev != nil; curr, prev = prev, prev.parent {
			if curr != prev.right {
				return prev
			}
		}

		return nil
	}
}

func (node *node[K, V]) locationInParent() **node[K, V] {
	if node.parent == nil {
		return nil
	}

	if node.parent.left == node {
		return &node.parent.left
	} else {
		return &node.parent.right
	}
}

func (node *node[K, V]) min() *node[K, V] {
	for curr := node; curr != nil; curr = curr.left {
		if curr.left == nil {
			return curr
		}
	}

	return nil
}

func (node *node[K, V]) max() *node[K, V] {
	for curr := node; curr != nil; curr = curr.right {
		if curr.right == nil {
			return curr
		}
	}

	return nil
}

func (tree *BinarySearchTree[K, V]) removeNode(node *node[K, V]) {
	locationInParent := node.locationInParent()
	if locationInParent == nil {
		locationInParent = &tree.root
	}

	if node.left == nil && node.right == nil {
		*locationInParent = nil
	} else if node.left == nil {
		*locationInParent = node.right
	} else if node.right == nil {
		*locationInParent = node.left
	} else {
		next := node.next()

		node.key, next.key = next.key, node.key
		node.value, next.value = next.value, node.value

		tree.removeNode(next)
	}

	tree.nodes--
}

func (tree *BinarySearchTree[K, V]) getNode(key K) *node[K, V] {
	for curr := tree.root; curr != nil; {
		if key == curr.key {
			return curr
		}

		switch tree.comparator(key, curr.key) {
		case comparison.FirstSmaller:
			curr = curr.left
		case comparison.FirstBigger:
			curr = curr.right
		case comparison.Equal:
			return curr
		}
	}

	return nil
}
