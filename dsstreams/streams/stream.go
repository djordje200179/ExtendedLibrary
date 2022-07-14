package streams

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/dsstreams/suppliers"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

func FromIterable[T any](iterable collections.Iterable[T]) streams.Stream[T] {
	return streams.New(suppliers.FromIterable(iterable))
}

func FromSequenceRefs[T any](sequence sequences.Sequence[T]) streams.Stream[*T] {
	return streams.New(suppliers.FromSequenceRefs(sequence))
}

func FromMap[K comparable, V any](m maps.Map[K, V]) streams.Stream[misc.Pair[K, V]] {
	return streams.New(suppliers.FromMap(m))
}
