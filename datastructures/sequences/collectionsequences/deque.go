package collectionsequences

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
)

type Deque[T any] struct {
	sequence collections.Collection[T]
}

func (deque Deque[T]) Empty() bool {
	return deque.sequence.Size() == 0
}

func (deque Deque[T]) PushFront(value T) {
	deque.sequence.Insert(0, value)

}
func (deque Deque[T]) PushBack(value T) {
	deque.sequence.Append(value)
}

func (deque Deque[T]) PeekFront() T {
	return deque.sequence.Get(0)
}

func (deque Deque[T]) PeekBack() T {
	return deque.sequence.Get(-1)
}

func (deque Deque[T]) PopFront() T {
	defer deque.sequence.Remove(0)
	return deque.PeekFront()
}
func (deque Deque[T]) PopBack() T {
	defer deque.sequence.Remove(-1)
	return deque.PeekBack()
}
