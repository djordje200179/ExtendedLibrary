package readcol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Wrapper is a wrapper around a cols.Collection that provides read-only access to it.
type Wrapper[T any] struct {
	collection cols.Collection[T]
}

// From creates a new Wrapper around the given cols.Collection.
func From[T any](collection cols.Collection[T]) Wrapper[T] {
	return Wrapper[T]{collection}
}

// Size returns the number of elements.
func (wrapper Wrapper[T]) Size() int {
	return wrapper.collection.Size()
}

// Get returns the element at the given index.
func (wrapper Wrapper[T]) Get(index int) T {
	return wrapper.collection.Get(index)
}

// Clone returns a copy of a Wrapper with the same underlying cols.Collection.
func (wrapper Wrapper[T]) Clone() Wrapper[T] {
	clonedCollection := wrapper.collection.Clone()
	return Wrapper[T]{clonedCollection}
}

// Iterator returns an iter.Iterator over the elements.
func (wrapper Wrapper[T]) Iterator() iter.Iterator[T] {
	return wrapper.collection.Iterator()
}

// Stream returns a streams.Stream over the elements.
func (wrapper Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.collection.Stream()
}

// FindIndex returns the index of the first element that satisfies the given predicate.
// If no element satisfies the predicate, 0 and false are returned.
func (wrapper Wrapper[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	return wrapper.collection.FindIndex(predicate)
}
