package collectors

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type sliceCollector[T any] struct {
	slice []T
}

func ToSlice[T any]() streams.Collector[T, []T] {
	return &sliceCollector[T]{
		slice: make([]T, 0),
	}
}

func (c *sliceCollector[T]) Supply(value T) {
	c.slice = append(c.slice, value)
}

func (c *sliceCollector[T]) Finish() []T {
	return c.slice
}
