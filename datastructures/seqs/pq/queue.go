package pq

import (
	"errors"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

// Queue is a priority queue implementation based on a binary heap.
// By default, the Queue is a min-heap, but a custom comparator can be provided.
type Queue[T any] struct {
	slice []T

	cmp comparison.Comparator[T]
}

// New creates a new Queue with the given comparator.
func New[T any](cmp comparison.Comparator[T]) *Queue[T] {
	pq := &Queue[T]{
		slice: make([]T, 0),

		cmp: cmp,
	}

	return pq
}

// NewFromIterable creates a new Queue with
// the given comparator and elements from the given iter.Iterable.
func NewFromIterable[T any](comparator comparison.Comparator[T], iterable iter.Iterable[T]) *Queue[T] {
	pq := New[T](comparator)

	for val := range iter.Iterate(iterable) {
		pq.PushBack(val)
	}

	return pq
}

// Empty returns true if there are no elements.
func (pq *Queue[T]) Empty() bool {
	return len(pq.slice) == 0
}

// PushBack adds the given value to the back.
//
// The element is inserted at the end of the Queue
// and then moved up the heap until the heap property is satisfied.
func (pq *Queue[T]) PushBack(value T) {
	pq.slice = append(pq.slice, value)

	currNode := len(pq.slice) - 1
	for {
		parentNode := (currNode - 1) / 2
		if parentNode == currNode || pq.cmp(pq.slice[currNode], pq.slice[parentNode]) != comparison.FirstSmaller {
			break
		}

		pq.slice[parentNode], pq.slice[currNode] = pq.slice[currNode], pq.slice[parentNode]
		currNode = parentNode
	}
}

// TryPushBack adds the given value to the back
// and is always successful.
func (pq *Queue[T]) TryPushBack(value T) bool {
	pq.PushBack(value)
	return true
}

// ErrNoElements is an error that occurs
// when there are no elements in the Queue.
var ErrNoElements = errors.New("no elements in the sequence")

// PeekFront returns the value at the front
// without removing it.
//
// ErrNoElements panic occurs if there are no elements.
func (pq *Queue[T]) PeekFront() T {
	if pq.Empty() {
		panic(ErrNoElements)
	}

	return pq.slice[0]
}

// TryPeekFront returns the value at the front
// without removing it and true if successful.
func (pq *Queue[T]) TryPeekFront() (T, bool) {
	if pq.Empty() {
		var zero T
		return zero, false
	}

	return pq.slice[0], true
}

// PopFront removes and returns the value at the front.
//
// ErrNoElements panic occurs if there are no elements.
func (pq *Queue[T]) PopFront() T {
	if pq.Empty() {
		panic(ErrNoElements)
	}

	lastIndex := len(pq.slice) - 1

	pq.slice[0], pq.slice[lastIndex] = pq.slice[lastIndex], pq.slice[0]

	currNode := 0
	for {
		leftChild := 2*currNode + 1
		if leftChild >= lastIndex || leftChild < 0 {
			break
		}

		child := leftChild
		if rightChild := leftChild + 1; rightChild < lastIndex && pq.cmp(pq.slice[rightChild], pq.slice[leftChild]) == comparison.FirstSmaller {
			child = rightChild
		}

		if pq.cmp(pq.slice[child], pq.slice[currNode]) != comparison.FirstSmaller {
			break
		}

		pq.slice[currNode], pq.slice[child] = pq.slice[child], pq.slice[currNode]
		currNode = child
	}

	lastElem := pq.slice[lastIndex]
	pq.slice = pq.slice[:lastIndex]
	return lastElem
}

// TryPopFront tries to remove and return
// the value at the front and true if successful.
func (pq *Queue[T]) TryPopFront() (T, bool) {
	if pq.Empty() {
		var zero T
		return zero, false
	}

	return pq.PopFront(), true
}
