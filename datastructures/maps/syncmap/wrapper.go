package syncmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

// Wrapper is a wrapper around a map that provides thread-safe access to the map.
// Locking is done through read-write mutex. This means that multiple goroutines can read from the map at the same time,
// but only one goroutine can write to the map at the same time.
type Wrapper[K, V any] struct {
	m maps.Map[K, V]

	mutex sync.RWMutex
}

// From creates a new Wrapper from the given map.
func From[K, V any](m maps.Map[K, V]) *Wrapper[K, V] {
	return &Wrapper[K, V]{m, sync.RWMutex{}}
}

// Size returns the number of entries in the map.
func (wrapper *Wrapper[K, V]) Size() int {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.Size()
}

// Contains returns true if the map contains the given key.
func (wrapper *Wrapper[K, V]) Contains(key K) bool {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.Contains(key)
}

// TryGet returns the value associated with the given key, and true if the key exists.
// If the key does not exist, the default value for the value type is returned, and false is returned.
func (wrapper *Wrapper[K, V]) TryGet(key K) (V, bool) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.TryGet(key)
}

// Get returns the value associated with the given key.
func (wrapper *Wrapper[K, V]) Get(key K) V {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.Get(key)
}

// GetRef returns a reference to the value associated with the given key.
// Usage of this method is discouraged, as it breaks the thread-safety of the map.
// Lock will not be held while the reference is used, so it is possible that the value of the element changes while the reference is used.
func (wrapper *Wrapper[K, V]) GetRef(key K) *V {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.GetRef(key)
}

// Set sets the value associated with the given key.
func (wrapper *Wrapper[K, V]) Set(key K, value V) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.m.Set(key, value)
}

// Update updates the value associated with the given key.
func (wrapper *Wrapper[K, V]) Update(key K, updateFunction func(oldValue V) V) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.m.Get(key)
	newValue := updateFunction(oldValue)

	wrapper.m.Set(key, newValue)
}

// UpdateRef updates the value associated with the given key.
func (wrapper *Wrapper[K, V]) UpdateRef(key K, updateFunction func(oldValue *V)) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.m.GetRef(key)
	updateFunction(oldValue)
}

// Remove removes the entry with the given key.
func (wrapper *Wrapper[K, V]) Remove(key K) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.m.Remove(key)
}

// Keys returns a slice of all keys in the map.
func (wrapper *Wrapper[K, V]) Keys() []K {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.Keys()
}

// Clear removes all entries from the map.
func (wrapper *Wrapper[K, V]) Clear() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.m.Clear()
}

// Clone returns a copy of the wrapper and the underlying map.
func (wrapper *Wrapper[K, V]) Clone() maps.Map[K, V] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	clonedMap := wrapper.m.Clone()
	return &Wrapper[K, V]{clonedMap, sync.RWMutex{}}
}

// Iterator returns an iterator over the map.
func (wrapper *Wrapper[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return wrapper.MapIterator()
}

// MapIterator returns an iterator over the map.
func (wrapper *Wrapper[K, V]) MapIterator() maps.Iterator[K, V] {
	return Iterator[K, V]{wrapper.m.MapIterator(), &wrapper.mutex}
}

// Stream returns a stream over the map.
func (wrapper *Wrapper[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return wrapper.m.Stream()
}

// RefsStream returns a stream over the map.
func (wrapper *Wrapper[K, V]) RefsStream() streams.Stream[misc.Pair[K, *V]] {
	return wrapper.m.RefsStream()
}

// Transaction executes the given function with the map as an argument.
// The map is locked for writing while the function is executed.
func (wrapper *Wrapper[K, V]) Transaction(updateFunction func(m maps.Map[K, V])) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	updateFunction(wrapper.m)
}
