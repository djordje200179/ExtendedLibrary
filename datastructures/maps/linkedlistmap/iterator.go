package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
)

type iterator[K comparable, V any] struct {
	collections.Iterator[misc.Pair[K, V]]
}

func (it iterator[K, V]) Get() K {
	return it.Iterator.Get().First
}

func (it iterator[K, V]) Value() K {
	return *it.ValueRef()
}

func (it iterator[K, V]) ValueRef() *K {
	return &it.Iterator.GetRef().First
}

func (it iterator[K, V]) SetValue(value K) {
	*it.ValueRef() = value
}
