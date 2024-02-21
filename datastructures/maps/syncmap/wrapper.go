package syncmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
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
func (w *Wrapper[K, V]) Size() int {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.m.Size()
}

// Contains returns true if the map contains the given key.
func (w *Wrapper[K, V]) Contains(key K) bool {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.m.Contains(key)
}

// TryGet returns the value associated with the given key, and true if the key exists.
// If the key does not exist, the default value for the value type is returned, and false is returned.
func (w *Wrapper[K, V]) TryGet(key K) (V, bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.m.TryGet(key)
}

// Get returns the value associated with the given key.
func (w *Wrapper[K, V]) Get(key K) V {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.m.Get(key)
}

// GetRef returns a reference to the value associated with the given key.
// Usage of this method is discouraged, as it breaks the thread-safety of the map.
// Lock will not be held while the reference is used, so it is possible that the value of the element changes while the reference is used.
func (w *Wrapper[K, V]) GetRef(key K) *V {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.m.GetRef(key)
}

// Set sets the value associated with the given key.
func (w *Wrapper[K, V]) Set(key K, value V) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.m.Set(key, value)
}

// Update updates the value associated with the given key.
func (w *Wrapper[K, V]) Update(key K, updateFunction func(oldValue V) V) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	oldValue := w.m.Get(key)
	newValue := updateFunction(oldValue)

	w.m.Set(key, newValue)
}

// UpdateRef updates the value associated with the given key.
func (w *Wrapper[K, V]) UpdateRef(key K, updateFunction func(oldValue *V)) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	oldValue := w.m.GetRef(key)
	updateFunction(oldValue)
}

// Remove removes the entry with the given key.
func (w *Wrapper[K, V]) Remove(key K) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.m.Remove(key)
}

// Clear removes all entries from the map.
func (w *Wrapper[K, V]) Clear() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.m.Clear()
}

// Clone returns a copy of the wrapper and the underlying map.
func (w *Wrapper[K, V]) Clone() maps.Map[K, V] {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	clonedMap := w.m.Clone()
	return &Wrapper[K, V]{clonedMap, sync.RWMutex{}}
}

// Iterator returns an iter.Iterator over the map.
func (w *Wrapper[K, V]) Iterator() iter.Iterator[misc.Pair[K, V]] {
	return w.MapIterator()
}

// MapIterator returns an iterator over the map.
func (w *Wrapper[K, V]) MapIterator() maps.Iterator[K, V] {
	return Iterator[K, V]{w.m.MapIterator(), &w.mutex}
}

// Stream2 streams over the entries in the Map.
func (w *Wrapper[K, V]) Stream2(yield func(K, V) bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	w.m.Stream2(yield)
}

// RefsStream2 streams over the keys and references to the values in the Map.
func (w *Wrapper[K, V]) RefsStream2(yield func(K, *V) bool) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.m.RefsStream2(yield)
}

// Transaction executes the given function with the map as an argument.
// The map is locked for writing while the function is executed.
func (w *Wrapper[K, V]) Transaction(updateFunction func(m maps.Map[K, V])) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	updateFunction(w.m)
}
