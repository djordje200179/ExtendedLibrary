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

func (node *Node[K, V]) Prev() *Node[K, V] {

}

func (node *Node[K, V]) Next() *Node[K, V] {

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
