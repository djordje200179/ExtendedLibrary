package syncset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"sync"
)

// Iterator is a wrapper around an iterator that provides thread-safe access to the
// underlying collection.
type Iterator[T any] struct {
	setIt sets.Iterator[T]

	mutex *sync.RWMutex
}

// Valid returns true if the iterator is currently pointing to a valid element.
func (it Iterator[T]) Valid() bool {
	return it.setIt.Valid()
}

// Move moves the iterator to the next element.
func (it Iterator[T]) Move() {
	it.setIt.Move()
}

// Get returns the current element.
func (it Iterator[T]) Get() T {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.setIt.Get()
}

// Remove removes the current element.
func (it Iterator[T]) Remove() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.setIt.Remove()
}
