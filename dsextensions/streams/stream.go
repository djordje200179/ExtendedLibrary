package streams

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/extensions/suppliers"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/streams"
)

func FromIterable[T any](iterable datastructures.Iterable[T]) streams.Stream[T] {
	return streams.New(suppliers.FromIterable(iterable))
}

func FromSequence[T any](sequence sequences.Sequence[T]) streams.Stream[*T] {
	return streams.New(suppliers.FromSequence(sequence))
}
