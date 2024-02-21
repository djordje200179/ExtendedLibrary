package synclist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
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

// NewFromIterable creates a List from the specified iter.Iterable in
// reverse order. List.Reverse can be used to restore the original order.
func NewFromIterable[T any](iterable iter.Iterable[T]) *List[T] {
	list := New[T]()

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		list.Prepend(it.Get())
	}

	return list
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
func (list *List[T]) Iterator() iter.Iterator[T] {
	return &Iterator[T]{curr: list.head.Load()}
}

// Stream streams the elements of the List.
func (list *List[T]) Stream(yield func(T) bool) {
	for curr := list.head.Load(); curr != nil; curr = curr.next {
		if !yield(curr.Value) {
			break
		}
	}
}

// Head returns the first element in the List.
func (list *List[T]) Head() *Node[T] {
	return list.head.Load()
}
