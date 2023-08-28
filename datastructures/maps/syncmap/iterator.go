package syncmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"sync"
)

// Iterator is a wrapper around an map iterator that provides thread-safe access to the
// underlying map.
type Iterator[K, V any] struct {
	mapIt maps.Iterator[K, V]

	mutex *sync.RWMutex
}

// Valid returns true if the iterator is currently pointing to a valid entry.
func (it Iterator[K, V]) Valid() bool {
	return it.mapIt.Valid()
}

// Move moves the iterator to the next entry.
func (it Iterator[K, V]) Move() {
	it.mapIt.Move()
}

// Get returns the current entry as a key-value pair.
func (it Iterator[K, V]) Get() misc.Pair[K, V] {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return misc.Pair[K, V]{
		First:  it.mapIt.Key(),
		Second: it.mapIt.Value(),
	}
}

// Key returns the current entry's key.
func (it Iterator[K, V]) Key() K {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.mapIt.Key()
}

// Value returns the current entry's value.
func (it Iterator[K, V]) Value() V {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.mapIt.Value()
}

// ValueRef returns a reference to the current entry's value.
func (it Iterator[K, V]) ValueRef() *V {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.mapIt.ValueRef()
}

// SetValue sets the current entry's value.
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
