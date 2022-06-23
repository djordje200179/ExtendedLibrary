package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc"
)

type iterator[K comparable, V any] struct {
	sequences.Iterator[misc.Pair[K, V]]
	m *Map[K, V]
}

func (it iterator[K, V]) Get() maps.Entry[K, V] {
	return maps.NewEntry[K, V](it.m, it.Iterator.Get().First)
}
