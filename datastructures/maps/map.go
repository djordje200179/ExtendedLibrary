package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Map[K comparable, V any] interface {
	datastructures.Collection[K, V]

	Remove(key K)
	Contains(key K) bool

	Empty()

	Iterator() Iterator[K, V]
	streams.Streamer[misc.Pair[K, V]]
}
