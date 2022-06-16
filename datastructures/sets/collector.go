package sets

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T comparable] struct {
	set Set[T]
}

func Collector[T comparable](empty Set[T]) streams.Collector[T, Set[T]] {
	return collector[T]{
		set: empty,
	}
}

func (c collector[T]) Supply(value T) {
	c.set.Add(value)
}

func (c collector[T]) Finish() Set[T] {
	return c.set
}
