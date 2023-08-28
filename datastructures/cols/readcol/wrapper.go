package readcol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Wrapper is a wrapper around a collection that provides read-only access to the collection.
type Wrapper[T any] struct {
	collection cols.Collection[T]
}

// From creates a new Wrapper from the given collection.
func From[T any](collection cols.Collection[T]) Wrapper[T] {
	return Wrapper[T]{collection}
}

// Size returns the number of elements in the collection.
func (wrapper Wrapper[T]) Size() int {
	return wrapper.collection.Size()
}

// Get returns the element at the given index.
func (wrapper Wrapper[T]) Get(index int) T {
	return wrapper.collection.Get(index)
}

// Clone returns a shallow copy of a Wrapper.
// Cloned Wrapper will have the same underlying collection as the original Wrapper.
func (wrapper Wrapper[T]) Clone() Wrapper[T] {
	clonedCollection := wrapper.collection.Clone()
	return Wrapper[T]{clonedCollection}
}

// Iterator returns an iterator over the elements in the collection.
func (wrapper Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.collection.Iterator()
}

// Stream returns a stream over the elements in the collection.
func (wrapper Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.collection.Stream()
}

// FindIndex returns the index of the first element that satisfies the given predicate.
func (wrapper Wrapper[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	return wrapper.collection.FindIndex(predicate)
}
