package collectors

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type slicify[T any] struct {
	slice *[]T
}

func Slicify[T any]() streams.Collector[T, []T] {
	return slicify[T]{
		slice: new([]T),
	}
}

func (slicify slicify[T]) Supply(value T) {
	*slicify.slice = append(*slicify.slice, value)
}

func (slicify slicify[T]) Finish() []T {
	return *slicify.slice
}
