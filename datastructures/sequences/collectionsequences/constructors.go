package collectionsequences

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/collections/array"
)

func NewDequeFrom[T any](sequence collections.Collection[T]) Deque[T] {
	return Deque[T]{sequence}
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{array.New[T]()}
}
