package synclist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"sync/atomic"
)

// List is a lock-free singly linked list.
//
// It is safe to use it concurrently from multiple goroutines.
//
// The zero value is ready to use.
// Do not copy a non-zero List.
type List[T any] struct {
	head atomic.Pointer[Node[T]]
}

// New creates an empty List.
func New[T any]() *List[T] {
	return new(List[T])
}

// NewFromIterable creates a List from the specified iter.Iterable
// in reverse order.
// List.Reverse method can be used to restore the original order if needed.
func NewFromIterable[T any](iterable iter.Iterable[T]) *List[T] {
	list := New[T]()

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		list.Prepend(it.Get())
	}

	return list
}

// Prepend atomically inserts the specified element at the beginning.
func (list *List[T]) Prepend(value T) {
	node := &Node[T]{Value: value}

	node.next = list.head.Swap(node)
}

// Clear atomically removes all elements.
func (list *List[T]) Clear() {
	list.head.Store(nil)
}

// Reverse reverses the order of the elements.
//
// This method is not thread-safe.
// It is meant to be used after all
// concurrent operations on the List have finished.
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

// Iterator returns an Iterator over the elements.
//
// Iteration starts from the first element.
func (list *List[T]) Iterator() iter.Iterator[T] {
	return &Iterator[T]{curr: list.head.Load()}
}

// Stream streams all elements.
func (list *List[T]) Stream(yield func(T) bool) {
	for curr := list.head.Load(); curr != nil; curr = curr.next {
		if !yield(curr.Value) {
			break
		}
	}
}

// Stream2 streams all elements with their indices.
func (list *List[T]) Stream2(yield func(int, T) bool) {
	for curr, i := list.head.Load(), 0; curr != nil; curr, i = curr.next, i+1 {
		if !yield(i, curr.Value) {
			break
		}
	}
}

// Head returns the first Node.
func (list *List[T]) Head() *Node[T] {
	return list.head.Load()
}
