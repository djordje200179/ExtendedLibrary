package bst

import "github.com/djordje200179/extendedlibrary/datastructures/maps"

type iterator[K comparable, V any] struct {
	tree *BinarySearchTree[K, V]

	curr *node[K, V]
}

func (it *iterator[K, V]) Valid() bool {
	return it.curr != nil
}

func (it *iterator[K, V]) Move() {
	it.curr = it.curr.next()
}

func (it *iterator[K, V]) Get() maps.Entry[K, V] {
	return entry[K, V]{it.curr}
}

func (it *iterator[K, V]) Remove() {
	it.tree.removeNode(it.curr)
}
