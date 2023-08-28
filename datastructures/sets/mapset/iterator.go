package mapset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
)

// Iterator is an iterator over a set.
type Iterator[T any] struct {
	mapIt maps.Iterator[T, empty]
}

// Valid returns true if the iterator is pointing to a valid element.
func (it Iterator[T]) Valid() bool {
	return it.mapIt.Valid()
}

// Move moves the iterator to the next element.
func (it Iterator[T]) Move() {
	it.mapIt.Move()
}

// Get returns the current element.
func (it Iterator[T]) Get() T {
	return it.mapIt.Key()
}

// Remove removes the current element from the set.
func (it Iterator[T]) Remove() {
	it.mapIt.Remove()
}
