package collectionsequence

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/collections/array"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Deque[T any] struct {
	collection collections.Collection[T]
}

func NewDequeFrom[T any](sequence collections.Collection[T]) Deque[T] {
	return Deque[T]{sequence}
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{array.New[T]()}
}

func Collector[T any]() streams.Collector[T, sequences.Queue[T]] {
	return sequences.Collector[T, sequences.Queue[T]]{
		BackPusher: NewDeque[T](),
	}
}

func (deque Deque[T]) Empty() bool {
	return deque.collection.Size() == 0
}

func (deque Deque[T]) PushFront(value T) {
	deque.collection.Insert(0, value)

}
func (deque Deque[T]) PushBack(value T) {
	deque.collection.Append(value)
}

func (deque Deque[T]) PeekFront() T {
	return deque.collection.Get(0)
}

func (deque Deque[T]) PeekBack() T {
	return deque.collection.Get(-1)
}

func (deque Deque[T]) PopFront() T {
	defer deque.collection.Remove(0)
	return deque.PeekFront()
}
func (deque Deque[T]) PopBack() T {
	defer deque.collection.Remove(-1)
	return deque.PeekBack()
}
