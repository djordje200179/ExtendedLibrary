package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
)

type iterator[K comparable, V any] struct {
	collections.Iterator[misc.Pair[K, V]]
	m *LinkedListMap[K, V]
}

func (it iterator[K, V]) Get() maps.Entry[K, V] { return entry[K, V]{it.Iterator.GetRef()} }
