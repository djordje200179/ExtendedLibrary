package sequences

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T any] struct {
	seq Sequence[T]
}

func Collector[T any](empty Sequence[T]) streams.Collector[T, Sequence[T]] {
	return collector[T]{
		seq: empty,
	}
}

func (c collector[T]) Supply(value T) {
	c.seq.Append(value)
}

func (c collector[T]) Finish() Sequence[T] {
	return c.seq
}
