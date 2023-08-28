package pq

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

// Queue is a priority queue implementation based on a binary heap.
// By default, the queue is a min-heap, but a custom comparator can be provided.
type Queue[T any] struct {
	slice []T

	comparator comparison.Comparator[T]
}

// New creates a new priority queue with the given comparator.
func New[T any](comparator comparison.Comparator[T]) *Queue[T] {
	pq := &Queue[T]{
		slice: make([]T, 0),

		comparator: comparator,
	}

	return pq
}

// NewFromIterable creates a new Queue with the given comparator and elements from the given iter.Iterable
func NewFromIterable[T any](comparator comparison.Comparator[T], iterable iter.Iterable[T]) *Queue[T] {
	pq := New[T](comparator)

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		pq.PushBack(it.Get())
	}

	return pq
}

// Empty returns true if the queue is empty.
func (pq *Queue[T]) Empty() bool {
	return len(pq.slice) == 0
}

// PushBack pushes a new element into the queue.
// The element is inserted at the end of the queue and then moved up the heap until the heap property is satisfied.
func (pq *Queue[T]) PushBack(value T) {
	pq.slice = append(pq.slice, value)

	currNode := len(pq.slice) - 1
	for {
		parentNode := (currNode - 1) / 2
		if parentNode == currNode || pq.comparator(pq.slice[currNode], pq.slice[parentNode]) != comparison.FirstSmaller {
			break
		}

		pq.slice[parentNode], pq.slice[currNode] = pq.slice[currNode], pq.slice[parentNode]
		currNode = parentNode
	}
}

// PeekFront returns the element at the front of the queue.
// Panics if the queue is empty.
func (pq *Queue[T]) PeekFront() T {
	if pq.Empty() {
		panic("Priority queue is empty")
	}

	return pq.slice[0]
}

// PopFront removes and returns the element at the front of the queue.
// Panics if the queue is empty.
func (pq *Queue[T]) PopFront() T {
	if pq.Empty() {
		panic("Priority queue is empty")
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
		if rightChild := leftChild + 1; rightChild < lastIndex && pq.comparator(pq.slice[rightChild], pq.slice[leftChild]) == comparison.FirstSmaller {
			child = rightChild
		}

		if pq.comparator(pq.slice[child], pq.slice[currNode]) != comparison.FirstSmaller {
			break
		}

		pq.slice[currNode], pq.slice[child] = pq.slice[child], pq.slice[currNode]
		currNode = child
	}

	lastElem := pq.slice[lastIndex]
	pq.slice = pq.slice[:lastIndex]
	return lastElem
}
