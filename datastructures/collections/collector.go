package collections

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T any] struct {
	seq Collection[T]
}

func Collector[T any](empty Collection[T]) streams.Collector[T, Collection[T]] {
	return collector[T]{empty}
}

func (collector collector[T]) Supply(value T) { collector.seq.Append(value) }

func (collector collector[T]) Finish() Collection[T] { return collector.seq }
