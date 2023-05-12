package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Iterator[K comparable, V any] interface {
	iterable.Iterator[misc.Pair[K, V]]

	Key() K
	Value() V
	ValueRef() *V
	SetValue(value V)

	Remove()
}

type Map[K comparable, V any] interface {
	Size() int

	Get(key K) V
	GetOrDefault(key K) V
	GetOrElse(key K, value V) V
	TryGet(key K) (V, bool)
	GetRef(key K) *V
	Set(key K, value V)

	Keys() []K

	Remove(key K)
	Contains(key K) bool

	Clear()
	Swap(key1, key2 K)

	misc.Cloner[Map[K, V]]

	iterable.Iterable[misc.Pair[K, V]]
	ModifyingIterator() Iterator[K, V]
	streams.Streamer[misc.Pair[K, V]]
	RefsStream() streams.Stream[misc.Pair[K, *V]]
}
