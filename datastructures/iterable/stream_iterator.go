package iterable

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type StreamIterator[T any] struct {
	streams.Stream[T]

	current T

	started, ended bool
}

func (it *StreamIterator[T]) Valid() bool {
	return !it.ended
}

func (it *StreamIterator[T]) Move() {
	if elem := it.Stream.First(); elem.Valid {
		it.current = elem.Value
	} else {
		it.ended = true
	}
}

func (it *StreamIterator[T]) Get() T {
	if !it.started {
		it.Move()
		it.started = true
	}

	return it.current
}
