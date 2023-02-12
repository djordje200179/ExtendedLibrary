package collectionsequences

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/collections/array"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
)

func NewDequeFrom[T any](sequence collections.Collection[T]) Deque[T] {
	return Deque[T]{sequence}
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{array.New[T]()}
}

func NewQueueFrom[T any](sequence collections.Collection[T]) sequences.Queue[T] {
	return NewDequeFrom(sequence)
}

func NewQueue[T any]() sequences.Queue[T] {
	return NewDeque[T]()
}

func NewStackFrom[T any](sequence collections.Collection[T]) sequences.Stack[T] {
	return NewDequeFrom(sequence)
}

func NewStack[T any]() sequences.Stack[T] {
	return NewDeque[T]()
}
