package redblacktree

import "github.com/djordje200179/extendedlibrary/misc"

type Iterator[K, V any] struct {
	tree *Tree[K, V]

	curr *Node[K, V]
}

func (it *Iterator[K, V]) Valid() bool {
	return it.curr != nil
}

func (it *Iterator[K, V]) Move() {
	it.curr = it.curr.Next()
}

func (it *Iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.Pair[K, V]{it.Key(), it.Value()}
}

func (it *Iterator[K, V]) Key() K {
	return it.curr.key
}

func (it *Iterator[K, V]) Value() V {
	return it.curr.Value
}

func (it *Iterator[K, V]) ValueRef() *V {
	return &it.curr.Value
}

func (it *Iterator[K, V]) SetValue(value V) {
	it.curr.Value = value
}

func (it *Iterator[K, V]) Remove() {
	//TODO implement me
	panic("implement me")
}

func (it *Iterator[K, V]) Node() *Node[K, V] {
	return it.curr
}
