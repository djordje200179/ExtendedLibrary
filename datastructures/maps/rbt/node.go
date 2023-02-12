package rbt

import "github.com/djordje200179/extendedlibrary/misc/functions/comparison"

type color bool

const (
	black color = false
	red
)

type node[K comparable, V any] struct {
	key   K
	value V

	left, right, parent *node[K, V]

	c color
}

func (tree *RedBlackTree[K, V]) getNode(key K) *node[K, V] {
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
