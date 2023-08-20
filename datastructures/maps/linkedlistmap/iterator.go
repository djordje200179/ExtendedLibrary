package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections/linkedlist"
	"github.com/djordje200179/extendedlibrary/misc"
)

type Iterator[K comparable, V any] struct {
	listIt *linkedlist.Iterator[misc.Pair[K, V]]
}

func (it Iterator[K, V]) Valid() bool {
	return it.listIt.Valid()
}

func (it Iterator[K, V]) Move() {
	it.listIt.Move()
}

func (it Iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.MakePair(it.Key(), it.Value())
}

func (it Iterator[K, V]) Key() K {
	return it.listIt.GetRef().First
}

func (it Iterator[K, V]) Value() V {
	return it.listIt.GetRef().Second
}

func (it Iterator[K, V]) ValueRef() *V {
	return &it.listIt.GetRef().Second
}

func (it Iterator[K, V]) SetValue(value V) {
	it.listIt.GetRef().Second = value
}

func (it Iterator[K, V]) Remove() {
	it.listIt.Remove()
}

func (it Iterator[K, V]) Node() *linkedlist.Node[misc.Pair[K, V]] {
	return it.listIt.Node()
}
