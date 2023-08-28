package sets

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Iterator is an iterator over a set.
type Iterator[T any] interface {
	iterable.Iterator[T]

	Remove() // Remove removes the current element
}

// Set is an interface that represents a set of elements
type Set[T any] interface {
	Size() int // Size returns the size of the set

	Add(value T)           // Add adds the given value to the set
	Remove(value T)        // Remove removes the given value from the set
	Contains(value T) bool // Contains returns true if the set contains the given value

	Clear() // Clear clears the set
	misc.Cloner[Set[T]]

	iterable.Iterable[T]
	SetIterator() Iterator[T] // SetIterator returns an iterator over the set
	streams.Streamer[T]
}
