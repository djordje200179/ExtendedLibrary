package syncmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"sync"
)

// Iterator is a wrapper around a maps.Iterator
// that provides thread-safe access.
type Iterator[K, V any] struct {
	mapIt maps.Iterator[K, V]

	mutex *sync.RWMutex
}

// Valid returns if the iterator is
// currently pointing to a valid entry.
func (it Iterator[K, V]) Valid() bool {
	return it.mapIt.Valid()
}

// Move moves to the next entry.
func (it Iterator[K, V]) Move() {
	it.mapIt.Move()
}

// Get returns the current entry as a key-value misc.Pair
func (it Iterator[K, V]) Get() misc.Pair[K, V] {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return misc.MakePair(it.mapIt.Key(), it.mapIt.Value())
}

// Key returns the key of the current entry.
func (it Iterator[K, V]) Key() K {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.mapIt.Key()
}

// Value returns the value of the current entry.
func (it Iterator[K, V]) Value() V {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.mapIt.Value()
}

// ValueRef returns a reference to the value of the current entry.
//
// Usage of this method is discouraged, as it breaks the thread-safety.
// Lock will not be held while the reference is used, so it is possible
// that the value of the entry changes while the reference is used.
func (it Iterator[K, V]) ValueRef() *V {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.mapIt.ValueRef()
}

// SetValue sets the value of the current entry.
func (it Iterator[K, V]) SetValue(value V) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.mapIt.SetValue(value)
}

// Remove removes the current entry.
func (it Iterator[K, V]) Remove() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.mapIt.Remove()
}
