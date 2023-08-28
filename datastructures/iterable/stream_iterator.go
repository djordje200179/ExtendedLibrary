package iterable

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

// StreamIterator is an iterator that iterates over a stream.
type StreamIterator[T any] struct {
	stream streams.Stream[T]

	current T

	started, ended bool
}

// NewStreamIterator creates a new stream iterator.
func NewStreamIterator[T any](stream streams.Stream[T]) *StreamIterator[T] {
	return &StreamIterator[T]{
		stream: stream,
	}
}

// Valid returns true if the iterator is valid.
func (it *StreamIterator[T]) Valid() bool {
	return !it.ended
}

// Move moves fetches the next element from the stream.
func (it *StreamIterator[T]) Move() {
	if elem := it.stream.First(); elem.Valid {
		it.current = elem.Value
	} else {
		it.ended = true
	}
}

// Get returns the current element.
// Value fetched by Move is cached and returned by Get.
// Therefore, it is safe to call Get multiple times.
func (it *StreamIterator[T]) Get() T {
	if !it.started {
		it.Move()
		it.started = true
	}

	return it.current
}
