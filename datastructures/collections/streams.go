package collections

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

type supplier[T any] struct {
	Iterator[T]
}

func ValuesStream[T any](sequence Collection[T]) streams.Stream[T] {
	supplier := supplier[T]{sequence.ModifyingIterator()}
	return streams.FromFiniteGenerator(supplier.NextValue)
}

func RefsStream[T any](sequence Collection[T]) streams.Stream[*T] {
	supplier := supplier[T]{sequence.ModifyingIterator()}
	return streams.FromFiniteGenerator(supplier.NextRef)
}

func (supplier supplier[T]) NextValue() optional.Optional[T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.Get())
	} else {
		return optional.Empty[T]()
	}
}

func (supplier supplier[T]) NextRef() optional.Optional[*T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.GetRef())
	} else {
		return optional.Empty[*T]()
	}
}
