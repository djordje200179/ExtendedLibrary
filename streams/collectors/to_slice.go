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

func (collector *sliceCollector[T]) Supply(value T) {
	collector.slice = append(collector.slice, value)
}

func (collector *sliceCollector[T]) Finish() []T {
	return collector.slice
}
