package readset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Wrapper is a wrapper around a Set that provides read-only access to it.
type Wrapper[T any] struct {
	set sets.Set[T]
}

// From creates a new Wrapper around the given Set.
func From[T any](set sets.Set[T]) Wrapper[T] {
	return Wrapper[T]{set}
}

// Size returns the number of elements in the Set.
func (wrapper Wrapper[T]) Size() int {
	return wrapper.set.Size()
}

// Contains returns true if the Set contains the given value.
func (wrapper Wrapper[T]) Contains(value T) bool {
	return wrapper.set.Contains(value)
}

// Clone returns a shallow copy of the Set.
func (wrapper Wrapper[T]) Clone() Wrapper[T] {
	clonedSet := wrapper.set.Clone()
	return Wrapper[T]{clonedSet}
}

// Iterator returns an iterator over the Set.
func (wrapper Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.set.Iterator()
}

// Stream returns a stream over the Set.
func (wrapper Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.set.Stream()
}
