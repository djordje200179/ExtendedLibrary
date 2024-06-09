package syncmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"sync"
)

// Wrapper is a wrapper around a maps.Map
// that provides thread-safe access.
//
// Locking is done through read-write mutex.
// This means that multiple goroutines can read
// at the same time, but only one goroutine can write at that time.
type Wrapper[K, V any] struct {
	m maps.Map[K, V]

	mutex sync.RWMutex
}

// From creates a new Wrapper around the given maps.Map.
func From[K, V any](m maps.Map[K, V]) *Wrapper[K, V] {
	return &Wrapper[K, V]{m, sync.RWMutex{}}
}

// Size returns the number of entries.
func (w *Wrapper[K, V]) Size() int {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.m.Size()
}

// Contains returns true if the given key is present.
func (w *Wrapper[K, V]) Contains(key K) bool {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.m.Contains(key)
}

// TryGet returns the value associated with the
// given key if it is present.
// If the key is not present, it returns the zero value
// for the value type and false.
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
//
// Usage of this method is discouraged, as it breaks the thread-safety.
// Lock will not be held while the reference is used, so it is possible
// that the value of the element changes while the reference is used.
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

// Update calculates and sets the new value
// of the entry at the given key using
// the given update function.
func (w *Wrapper[K, V]) Update(key K, updateFunction func(oldValue V) V) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	oldValue := w.m.Get(key)
	newValue := updateFunction(oldValue)

	w.m.Set(key, newValue)
}

// UpdateRef updates the value at the given key
// in-place using the given update function.
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

// Clear removes all entries.
func (w *Wrapper[K, V]) Clear() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.m.Clear()
}

// Clone returns a new Wrapper with
// a clone of the underlying maps.Map.
func (w *Wrapper[K, V]) Clone() maps.Map[K, V] {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	clonedMap := w.m.Clone()
	return &Wrapper[K, V]{clonedMap, sync.RWMutex{}}
}

// Iterator returns a read-only iter.Iterator over the entries.
func (w *Wrapper[K, V]) Iterator() iter.Iterator[misc.Pair[K, V]] {
	return w.MapIterator()
}

// MapIterator returns an Iterator over the entries.
// It can be used to modify the entries while iterating.
func (w *Wrapper[K, V]) MapIterator() maps.Iterator[K, V] {
	return Iterator[K, V]{w.m.MapIterator(), &w.mutex}
}

// Stream2 streams all entries.
//
// Updates to the underlying maps.Map
// while streaming are not allowed.
func (w *Wrapper[K, V]) Stream2(yield func(K, V) bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	w.m.Stream2(yield)
}

// Keys streams the keys.
func (w *Wrapper[K, V]) Keys(yield func(K) bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	w.m.Keys(yield)
}

// Values streams the values.
func (w *Wrapper[K, V]) Values(yield func(V) bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	w.m.Values(yield)
}

// Transaction executes the given function
// on the underlying maps.Map
// while holding the write lock.
func (w *Wrapper[K, V]) Transaction(updateFunction func(m maps.Map[K, V])) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	updateFunction(w.m)
}
