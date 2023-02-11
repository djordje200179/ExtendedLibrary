package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Entry[K comparable, V any] interface {
	Key() K

	Value() V
	ValueRef() *V
	SetValue(value V)
}

type Iterator[K comparable, V any] interface {
	iterable.Iterator[Entry[K, V]]

	Remove()
}

type Map[K comparable, V any] interface {
	Size() int

	Get(key K) V
	GetRef(key K) *V
	Set(key K, value V)

	Remove(key K)
	Contains(key K) bool

	Clear()
	misc.Cloner[Map[K, V]]

	iterable.Iterable[Entry[K, V]]
	ModifyingIterator() Iterator[K, V]
	streams.Streamer[misc.Pair[K, V]]
	RefStream() streams.Stream[misc.Pair[K, *V]]
}
