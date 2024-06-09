package sets

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc"
)

// An Iterator is used to iterate
// through values of a Set.
type Iterator[T any] interface {
	iter.Iterator[T]

	// Remove removes the current element
	Remove()
}

// Set is a data structure for storing unique values.
type Set[T any] interface {
	// Size returns the cardinality.
	Size() int

	// Add inserts the given value.
	Add(value T)
	// Remove removes the given value.
	Remove(value T)
	// Contains returns true if the
	// value is already present.
	Contains(value T) bool

	// Clear removes all the values.
	Clear()
	misc.Cloner[Set[T]]

	iter.Iterable[T]
	// SetIterator returns a specialized
	// Iterator over the elements.
	SetIterator() Iterator[T]
	// Stream streams the elements.
	Stream(yield func(T) bool)
}
