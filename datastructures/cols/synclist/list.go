package synclist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync/atomic"
)

// List is a lock-free singly linked list.
// It is safe to use it concurrently from multiple goroutines.
// The zero value is ready to use. Do not copy a non-zero List.
type List[T any] struct {
	head atomic.Pointer[Node[T]]
}

// New creates an empty List.
func New[T any]() *List[T] {
	return new(List[T])
}

// Prepend adds the specified value to the beginning of the List.
func (list *List[T]) Prepend(value T) {
	node := &Node[T]{Value: value}

	node.next = list.head.Swap(node)
}

// Clear removes all elements from the List.
func (list *List[T]) Clear() {
	list.head.Store(nil)
}

// Reverse reverses the order of the elements in the List.
// This method is not thread-safe. It is meant to be used after
// all concurrent operations on the List have finished.
func (list *List[T]) Reverse() {
	var prev *Node[T]

	curr := list.head.Load()
	for curr != nil {
		next := curr.next

		curr.next = prev

		prev = curr
		curr = next
	}

	list.head.Store(prev)
}

// Iterator returns an iterator over the elements in the List.
func (list *List[T]) Iterator() iterable.Iterator[T] {
	return &Iterator[T]{curr: list.head.Load()}
}

// Stream returns a stream over the elements in the List.
func (list *List[T]) Stream() streams.Stream[T] {
	return iterable.IteratorStream(list.Iterator())
}

// Head returns the first element in the List.
func (list *List[T]) Head() *Node[T] {
	return list.head.Load()
}
