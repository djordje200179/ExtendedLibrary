package colseq

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/cols/array"
	"github.com/djordje200179/extendedlibrary/datastructures/cols/linklist"
	"github.com/djordje200179/extendedlibrary/datastructures/seqs"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Deque[T any] struct {
	collection cols.Collection[T]
}

func NewArrayDeque[T any]() Deque[T] {
	return Deque[T]{array.New[T]()}
}

func NewLinkedListDeque[T any]() Deque[T] {
	return Deque[T]{linklist.New[T]()}
}

func From[T any](sequence cols.Collection[T]) Deque[T] {
	return Deque[T]{sequence}
}

func ArrayCollector[T any]() streams.Collector[T, seqs.Queue[T]] {
	return seqs.Collector[T, seqs.Queue[T]]{NewArrayDeque[T]()}
}

func LinkedListCollector[T any]() streams.Collector[T, seqs.Queue[T]] {
	return seqs.Collector[T, seqs.Queue[T]]{NewLinkedListDeque[T]()}
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
	if deque.Empty() {
		panic("Deque is empty")
	}

	return deque.collection.Get(0)
}

func (deque Deque[T]) PeekBack() T {
	if deque.Empty() {
		panic("Deque is empty")
	}

	return deque.collection.Get(-1)
}

func (deque Deque[T]) PopFront() T {
	if deque.Empty() {
		panic("Deque is empty")
	}

	defer deque.collection.Remove(0)
	return deque.PeekFront()
}
func (deque Deque[T]) PopBack() T {
	if deque.Empty() {
		panic("Deque is empty")
	}

	defer deque.collection.Remove(-1)
	return deque.PeekBack()
}
