package binarysearchtree

import "github.com/djordje200179/extendedlibrary/misc"

type iterator[K comparable, V any] struct {
	tree *BinarySearchTree[K, V]

	curr *Node[K, V]
}

func (it *iterator[K, V]) Valid() bool {
	return it.curr != nil
}

func (it *iterator[K, V]) Move() {
	it.curr = it.curr.Next()
}

func (it *iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.Pair[K, V]{
		First:  it.Key(),
		Second: it.Value(),
	}
}

func (it *iterator[K, V]) Key() K {
	return it.curr.key
}

func (it *iterator[K, V]) Value() V {
	return *it.ValueRef()
}

func (it *iterator[K, V]) ValueRef() *V {
	return &it.curr.Value
}

func (it *iterator[K, V]) SetValue(value V) {
	*it.ValueRef() = value
}

func (it *iterator[K, V]) Remove() {
	it.tree.removeNode(it.curr)
}
