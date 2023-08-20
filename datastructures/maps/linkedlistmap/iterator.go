package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
)

type Iterator[K comparable, V any] struct {
	listIt collections.Iterator[misc.Pair[K, V]]
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
