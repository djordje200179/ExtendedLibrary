package collections

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

func StreamIterator[T any](stream streams.Stream[T]) Iterator[T] {
	return &iterator[T]{stream: stream}
}

type iterator[T any] struct {
	stream  streams.Stream[T]
	current T

	started, ended bool
}

func (it *iterator[T]) Valid() bool { return !it.ended }

func (it *iterator[T]) Move() {
	if elem := it.stream.First(); elem.HasValue() {
		it.current = elem.Get()
	} else {
		it.ended = true
	}
}

func (it *iterator[T]) Get() T {
	if !it.started {
		it.Move()
		it.started = true
	}

	return it.current
}
