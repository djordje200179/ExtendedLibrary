package readset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
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
func (w Wrapper[T]) Size() int {
	return w.set.Size()
}

// Contains returns true if the Set contains the given value.
func (w Wrapper[T]) Contains(value T) bool {
	return w.set.Contains(value)
}

// Clone returns a shallow copy of the Set.
func (w Wrapper[T]) Clone() Wrapper[T] {
	clonedSet := w.set.Clone()
	return Wrapper[T]{clonedSet}
}

// Iterator returns an iter.Iterator over the Set.
func (w Wrapper[T]) Iterator() iter.Iterator[T] {
	return w.set.Iterator()
}

// Stream streams the elements of the Set.
func (w Wrapper[T]) Stream(yield func(T) bool) {
	w.set.Stream(yield)
}
