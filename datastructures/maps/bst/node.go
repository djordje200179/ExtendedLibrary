package bst

import "github.com/djordje200179/extendedlibrary/misc/comparison"

type Node[K comparable, V any] struct {
	key   K
	Value V

	left, right, parent *Node[K, V]
}

func (node *Node[K, V]) Key() K { return node.key }

func (node *Node[K, V]) Left() *Node[K, V]   { return node.left }
func (node *Node[K, V]) Right() *Node[K, V]  { return node.right }
func (node *Node[K, V]) Parent() *Node[K, V] { return node.parent }

func (tree *BinarySearchTree[K, V]) Root() *Node[K, V] { return tree.root }

func (node *Node[K, V]) Clone() *Node[K, V] {
	cloned := new(Node[K, V])
	cloned.key = node.key
	cloned.Value = node.Value

	return cloned
}

func (node *Node[K, V]) Prev() *Node[K, V] {
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

func (node *Node[K, V]) Next() *Node[K, V] {
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

func (node *Node[K, V]) min() *Node[K, V] {
	for curr := node; curr != nil; curr = curr.left {
		if curr.left == nil {
			return curr
		}
	}

	return nil
}

func (node *Node[K, V]) max() *Node[K, V] {
	for curr := node; curr != nil; curr = curr.right {
		if curr.right == nil {
			return curr
		}
	}

	return nil
}

func (tree *BinarySearchTree[K, V]) removeNode(node *Node[K, V]) {
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
		next := node.Next()

		node.key, next.key = next.key, node.key
		node.Value, next.Value = next.Value, node.Value

		tree.removeNode(next)
	}

	tree.nodes--
}

func (tree *BinarySearchTree[K, V]) GetNode(key K) *Node[K, V] {
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
