package hashset

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T comparable] struct {
	set Set[T]
}

func Collector[T comparable]() streams.Collector[T, Set[T]] {
	return collector[T]{
		set: New[T](),
	}
}

func (c collector[T]) Supply(value T) {
	c.set.Add(value)
}

func (c collector[T]) Finish() Set[T] {
	return c.set
}
