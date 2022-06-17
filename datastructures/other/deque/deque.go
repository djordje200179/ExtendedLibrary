package deque

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

type Deque[T any] struct {
	seq sequences.Sequence[T]
}

func New[T any]() Deque[T] {
	return Deque[T]{linkedlist.New[T]()}
}

func (deque Deque[T]) PushFront(value T) {
	deque.seq.Insert(0, value)
}

func (deque Deque[T]) PushBack(value T) {
	deque.seq.Append(value)
}

func (deque Deque[T]) PopFront() T {
	defer deque.seq.Remove(0)
	return deque.PeekFront()
}

func (deque Deque[T]) PopBack() T {
	defer deque.seq.Remove(-1)
	return deque.PeekBack()
}

func (deque Deque[T]) PeekFront() T {
	return deque.seq.Get(0)
}

func (deque Deque[T]) PeekBack() T {
	return deque.seq.Get(-1)
}

func (deque Deque[T]) Empty() bool {
	return deque.seq.Size() == 0
}
