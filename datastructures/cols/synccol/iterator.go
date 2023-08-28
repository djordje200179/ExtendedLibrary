package synccol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"sync"
)

// Iterator is a wrapper around a cols.Iterator that provides thread-safe
// access to the underlying collection.
type Iterator[T any] struct {
	colIt cols.Iterator[T]

	mutex *sync.RWMutex
}

// Valid returns true if the iterator is currently pointing to a valid element.
func (it Iterator[T]) Valid() bool {
	return it.colIt.Valid()
}

// Move moves the iterator to the next element.
func (it Iterator[T]) Move() {
	it.colIt.Move()
}

// GetRef returns a reference to the current element.
func (it Iterator[T]) GetRef() *T {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.colIt.GetRef()
}

// Get returns the current element.
func (it Iterator[T]) Get() T {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.colIt.Get()
}

// Set sets the current element.
func (it Iterator[T]) Set(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.colIt.Set(value)
}

// InsertBefore inserts the specified element before the current element.
func (it Iterator[T]) InsertBefore(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.colIt.InsertBefore(value)
}

// InsertAfter inserts the specified element after the current element.
func (it Iterator[T]) InsertAfter(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.colIt.InsertAfter(value)
}

// Remove removes the current element.
func (it Iterator[T]) Remove() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.colIt.Remove()
}

// Index returns the current index.
func (it Iterator[T]) Index() int {
	return it.colIt.Index()
}
