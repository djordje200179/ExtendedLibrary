package colseq

import (
	"errors"
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/cols/array"
	"github.com/djordje200179/extendedlibrary/datastructures/cols/linklist"
)

// Deque is a seqs.Deque implemented
// using a cols.Collection.
type Deque[T any] struct {
	coll cols.Collection[T]
}

// NewArrayDeque creates a new Deque
// backed by an array.Array.
func NewArrayDeque[T any]() Deque[T] { return Deque[T]{array.New[T]()} }

// NewLinkedListDeque creates a new Deque
// backed by a linklist.List.
func NewLinkedListDeque[T any]() Deque[T] { return Deque[T]{linklist.New[T]()} }

// From creates a new Deque from
// the given cols.Collection.
func From[T any](coll cols.Collection[T]) Deque[T] { return Deque[T]{coll} }

// Empty returns true if there are no elements.
func (deque Deque[T]) Empty() bool { return deque.coll.Size() == 0 }

// PushFront adds the given value to the front.
func (deque Deque[T]) PushFront(value T) { deque.coll.Insert(0, value) }

// TryPushFront adds the given value to the front
// and is always successful.
func (deque Deque[T]) TryPushFront(value T) bool {
	deque.PushFront(value)
	return true
}

// PushBack adds the given value to the back.
func (deque Deque[T]) PushBack(value T) { deque.coll.Append(value) }

// TryPushBack adds the given value to the back
// and is always successful.
func (deque Deque[T]) TryPushBack(value T) bool {
	deque.PushBack(value)
	return true
}

// ErrNoElements is an error that occurs
// when there are no elements in the Deque.
var ErrNoElements = errors.New("no elements in the sequence")

// PeekFront returns the value at the front
// without removing it.
// ErrNoElements panic occurs if there are no elements.
func (deque Deque[T]) PeekFront() T {
	if deque.Empty() {
		panic(ErrNoElements)
	}

	return deque.coll.Get(0)
}

// TryPeekFront returns the value at the front
// without removing it and true if successful.
func (deque Deque[T]) TryPeekFront() (T, bool) {
	if deque.Empty() {
		var zero T
		return zero, false
	}

	return deque.PeekFront(), true
}

// PeekBack returns the value at the back
// without removing it.
// ErrNoElements panic occurs if there are no elements.
func (deque Deque[T]) PeekBack() T {
	if deque.Empty() {
		panic(ErrNoElements)
	}

	return deque.coll.Get(-1)
}

// TryPeekBack returns the value at the back
// without removing it and true if successful.
func (deque Deque[T]) TryPeekBack() (T, bool) {
	if deque.Empty() {
		var zero T
		return zero, false
	}

	return deque.PeekBack(), true
}

// PopFront removes and returns the value at the front.
// ErrNoElements panic occurs if there are no elements.
func (deque Deque[T]) PopFront() T {
	if deque.Empty() {
		panic(ErrNoElements)
	}

	defer deque.coll.Remove(0)
	return deque.PeekFront()
}

// TryPopFront removes and returns the value at the front
// and true if successful.
func (deque Deque[T]) TryPopFront() (T, bool) {
	if deque.Empty() {
		var zero T
		return zero, false
	}

	value := deque.PopFront()
	return value, true
}

// PopBack removes and returns the value at the back.
// ErrNoElements panic occurs if there are no elements.
func (deque Deque[T]) PopBack() T {
	if deque.Empty() {
		panic(ErrNoElements)
	}

	defer deque.coll.Remove(-1)
	return deque.PeekBack()
}
