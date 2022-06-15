package array

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T any] struct {
	array *Array[T]
}

func Collector[T any]() streams.Collector[T, *Array[T]] {
	return collector[T]{
		array: New[T](0),
	}
}

func (c collector[T]) Supply(value T) {
	c.array.Append(value)
}

func (c collector[T]) Finish() *Array[T] {
	return c.array
}
