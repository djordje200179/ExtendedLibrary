package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Map[K comparable, V any] interface {
	Size() int

	Get(key K) V
	Set(key K, value V)

	Remove(key K)
	Contains(key K) bool

	Iterator() Iterator[K, V]
	streams.Streamer[misc.Pair[K, V]]
}
