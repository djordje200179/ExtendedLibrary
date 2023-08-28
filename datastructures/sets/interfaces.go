package sets

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Iterator is an iterator over a set.
type Iterator[T any] interface {
	iter.Iterator[T]

	// Remove removes the current element
	Remove()
}

// Set is an interface that represents a set of elements
type Set[T any] interface {
	// Size returns the size of the set
	Size() int

	// Add adds the given value to the set
	Add(value T)
	// Remove removes the given value from the set
	Remove(value T)
	// Contains returns true if the set contains the given value
	Contains(value T) bool

	// Clear clears the set
	Clear()
	misc.Cloner[Set[T]]

	iter.Iterable[T]
	// SetIterator returns an iterator over the set
	SetIterator() Iterator[T]
	streams.Streamer[T]
}
