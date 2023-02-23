package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
)

type iterator[K comparable, V any] struct {
	collections.Iterator[misc.Pair[K, V]]
}

func (it iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.Pair[K, V]{it.Key(), it.Value()}
}

func (it iterator[K, V]) Key() K {
	return it.Iterator.GetRef().First
}

func (it iterator[K, V]) Value() V {
	return *it.ValueRef()
}

func (it iterator[K, V]) ValueRef() *V {
	return &it.Iterator.GetRef().Second
}

func (it iterator[K, V]) SetValue(value V) {
	*it.ValueRef() = value
}
