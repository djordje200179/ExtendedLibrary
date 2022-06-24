package bst

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type BinarySearchTree[K comparable, V any] struct {
	root  *Node[K, V]
	nodes int

	comparator functions.Comparator[K]
}

func New[K comparable, V any](comparator functions.Comparator[K]) *BinarySearchTree[K, V] {
	tree := new(BinarySearchTree[K, V])
	tree.root = nil
	tree.nodes = 0
	tree.comparator = comparator

	return tree
}

func (tree *BinarySearchTree[K, V]) Size() int {
	return tree.nodes
}

func (tree *BinarySearchTree[K, V]) Get(key K) V {
	if node := tree.GetNode(key); node != nil {
		return node.Value
	} else {
		var empty V
		return empty
	}
}

func (tree *BinarySearchTree[K, V]) Set(key K, value V) {
	if tree.root == nil {
		tree.root = &Node[K, V]{
			key:   key,
			Value: value,
		}
		tree.nodes++

		return
	}

	prev := (*Node[K, V])(nil)
	for curr := tree.root; curr != nil; {
		if key == curr.key {
			curr.Value = value
			return
		}

		prev = curr
		switch tree.comparator(key, curr.key) {
		case comparison.FirstSmaller:
			curr = curr.left
		case comparison.FirstBigger:
			curr = curr.right
		case comparison.Equal:
			curr.Value = value
			return
		}
	}

	node := &Node[K, V]{
		key:    key,
		Value:  value,
		parent: prev,
	}

	if tree.comparator(key, prev.key) == comparison.FirstSmaller {
		prev.left = node
	} else {
		prev.right = node
	}

	tree.nodes++
}

func (tree *BinarySearchTree[K, V]) Remove(key K) {
	if node := tree.GetNode(key); node != nil {
		tree.removeNode(node)
	}
}

func (tree *BinarySearchTree[K, V]) Contains(key K) bool {
	return tree.GetNode(key) != nil
}

func (tree *BinarySearchTree[K, V]) Clear() {
	tree.root = nil
	tree.nodes = 0
}

func (tree *BinarySearchTree[K, V]) Clone() maps.Map[K, V] {
	cloned := New[K, V](tree.comparator)
	cloned.nodes = tree.nodes

	if tree.root == nil {
		return cloned
	}

	//TODO: Implement
	panic("Not implemented")

	return cloned
}

func (tree *BinarySearchTree[K, V]) Iterator() datastructures.Iterator[maps.Entry[K, V]] {
	return tree.ModifyingIterator()
}

func (tree *BinarySearchTree[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return &iterator[K, V]{
		tree: tree,
		curr: tree.root.min(),
	}
}

func (tree *BinarySearchTree[K, V]) Stream() *streams.Stream[misc.Pair[K, V]] {
	it := tree.Iterator()

	return streams.Supply(func() misc.Pair[K, V] {
		defer it.Move()

		entry := it.Get()
		return misc.Pair[K, V]{entry.Key(), entry.Value()}
	})
}
