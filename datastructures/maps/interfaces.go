package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Iterator[K comparable, V any] interface {
	datastructures.Iterator[Entry[K, V]]

	Remove()
}

type Map[K comparable, V any] interface {
	datastructures.Collection[K, V]
	Remove(key K)
	Contains(key K) bool

	Empty()
	misc.Cloner[Map[K, V]]

	Iterator() Iterator[K, V]
	streams.Streamer[misc.Pair[K, V]]
}
