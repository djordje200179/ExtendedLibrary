package readcol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
)

// Wrapper is a wrapper around a
// cols.Collection that provides read-only access to it.
type Wrapper[T any] struct {
	collection cols.Collection[T]
}

// From creates a new Wrapper around the given cols.Collection.
func From[T any](collection cols.Collection[T]) Wrapper[T] { return Wrapper[T]{collection} }

// Size returns the number of elements.
func (w Wrapper[T]) Size() int { return w.collection.Size() }

// Get returns the element at the specified index.
func (w Wrapper[T]) Get(index int) T { return w.collection.Get(index) }

// Clone returns a new Wrapper with
// a clone of the underlying cols.Collection.
func (w Wrapper[T]) Clone() Wrapper[T] { return Wrapper[T]{w.collection.Clone()} }

// Iterator returns an iter.Iterator over the elements.
func (w Wrapper[T]) Iterator() iter.Iterator[T] {
	return w.collection.Iterator()
}

// Stream streams all elements.
func (w Wrapper[T]) Stream(yield func(T) bool) { w.collection.Stream(yield) }

// Stream2 streams all elements with their indices.
func (w Wrapper[T]) Stream2(yield func(int, T) bool) { w.collection.Stream2(yield) }

// FindIndex returns the index of the first element
// that satisfies the specified predicate.
// If no such element is found, 0 and false are returned.
func (w Wrapper[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	return w.collection.FindIndex(predicate)
}
