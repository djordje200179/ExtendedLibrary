package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

func Stream[K comparable, V any](m Map[K, V]) *streams.Stream[misc.Pair[K, V]] {
	entryMapper := func(entry Entry[K, V]) misc.Pair[K, V] {
		return misc.Pair[K, V]{entry.Key(), entry.Value()}
	}
	
	return streams.Map[Entry[K, V], misc.Pair[K, V]](streams.FromIterable[Entry[K, V]](m), entryMapper)
}
