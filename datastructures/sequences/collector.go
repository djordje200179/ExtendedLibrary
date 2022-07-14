package sequences

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T any] struct {
	seq Sequence[T]
}

func Collector[T any](empty Sequence[T]) streams.Collector[T, Sequence[T]] {
	return collector[T]{empty}
}

func (collector collector[T]) Supply(value T) { collector.seq.Append(value) }

func (collector collector[T]) Finish() Sequence[T] { return collector.seq }
