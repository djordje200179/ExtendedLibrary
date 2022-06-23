package bst

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
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
	//TODO implement me
	panic("implement me")
}

func (tree *BinarySearchTree[K, V]) Remove(key K) {
	//TODO implement me
	panic("implement me")
}

func (tree *BinarySearchTree[K, V]) Contains(key K) bool {
	return tree.GetNode(key) != nil
}

func (tree *BinarySearchTree[K, V]) Clear() {
	tree.root = nil
	tree.nodes = 0
}

func (tree *BinarySearchTree[K, V]) Clone() maps.Map[K, V] {
	//TODO implement me
	panic("implement me")
}

func (tree *BinarySearchTree[K, V]) Iterator() datastructures.Iterator[maps.Entry[K, V]] {
	//TODO implement me
	panic("implement me")
}

func (tree *BinarySearchTree[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	//TODO implement me
	panic("implement me")
}

func (tree *BinarySearchTree[K, V]) Stream() *streams.Stream[misc.Pair[K, V]] {
	//TODO implement me
	panic("implement me")
}
