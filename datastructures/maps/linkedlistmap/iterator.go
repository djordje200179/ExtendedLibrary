package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
)

type Iterator[K comparable, V any] struct {
	collections.Iterator[misc.Pair[K, V]]
}

func (it Iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.Pair[K, V]{it.Key(), it.Value()}
}

func (it Iterator[K, V]) Key() K {
	return it.Iterator.GetRef().First
}

func (it Iterator[K, V]) Value() V {
	return *it.ValueRef()
}

func (it Iterator[K, V]) ValueRef() *V {
	return &it.Iterator.GetRef().Second
}

func (it Iterator[K, V]) SetValue(value V) {
	*it.ValueRef() = value
}
