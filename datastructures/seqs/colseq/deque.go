package colseq

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/cols/array"
	"github.com/djordje200179/extendedlibrary/datastructures/cols/linklist"
)

// Deque is a double-ended queue.
// It is a sequence of elements in which items can be inserted or removed from either the front or back.
// Pushing is always possible, but popping and peeking panics if the deque is empty.
type Deque[T any] struct {
	collection cols.Collection[T]
}

// NewArrayDeque creates a new Deque backed by an array.
func NewArrayDeque[T any]() Deque[T] {
	return Deque[T]{array.New[T]()}
}

// NewLinkedListDeque creates a new Deque backed by a linked list.
func NewLinkedListDeque[T any]() Deque[T] {
	return Deque[T]{linklist.New[T]()}
}

// From creates a new Deque from a collection.
func From[T any](collection cols.Collection[T]) Deque[T] {
	return Deque[T]{collection}
}

// Empty returns true if the Deque is empty.
func (deque Deque[T]) Empty() bool {
	return deque.collection.Size() == 0
}

// PushFront inserts an element at the front of the Deque.
func (deque Deque[T]) PushFront(value T) {
	deque.collection.Insert(0, value)
}

// PushBack inserts an element at the back of the Deque.
func (deque Deque[T]) PushBack(value T) {
	deque.collection.Append(value)
}

// PeekFront returns the element at the front of the Deque.
// Panics if the Deque is empty.
func (deque Deque[T]) PeekFront() T {
	if deque.Empty() {
		panic("Deque is empty")
	}

	return deque.collection.Get(0)
}

// PeekBack returns the element at the back of the Deque.
// Panics if the Deque is empty.
func (deque Deque[T]) PeekBack() T {
	if deque.Empty() {
		panic("Deque is empty")
	}

	return deque.collection.Get(-1)
}

// PopFront removes and returns the element at the front of the Deque.
// Panics if the Deque is empty.
func (deque Deque[T]) PopFront() T {
	if deque.Empty() {
		panic("Deque is empty")
	}

	defer deque.collection.Remove(0)
	return deque.PeekFront()
}

// PopBack removes and returns the element at the back of the Deque.
// Panics if the Deque is empty.
func (deque Deque[T]) PopBack() T {
	if deque.Empty() {
		panic("Deque is empty")
	}

	defer deque.collection.Remove(-1)
	return deque.PeekBack()
}
