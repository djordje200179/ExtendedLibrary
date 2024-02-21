package syncset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"sync"
)

// Wrapper is a wrapper around a set that provides thread-safe access to the set.
// Locking is done through read-write mutex. This means that multiple goroutines can read from the set at the same time,
// but only one goroutine can write to the set at the same time.
type Wrapper[T any] struct {
	set sets.Set[T]

	mutex sync.RWMutex
}

// From creates a new Wrapper from the given set.
func From[T any](set sets.Set[T]) *Wrapper[T] {
	return &Wrapper[T]{set, sync.RWMutex{}}
}

// Size returns the number of elements in the set.
func (w *Wrapper[T]) Size() int {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.set.Size()
}

// Add adds the given value to the set.
func (w *Wrapper[T]) Add(value T) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.set.Add(value)
}

// Remove removes the given value from the set.
func (w *Wrapper[T]) Remove(value T) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.set.Remove(value)
}

// Contains returns true if the set contains the given value.
func (w *Wrapper[T]) Contains(value T) bool {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.set.Contains(value)
}

// Clear removes all elements from the set.
func (w *Wrapper[T]) Clear() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.set.Clear()
}

// Clone returns a shallow copy of the set.
func (w *Wrapper[T]) Clone() sets.Set[T] {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	clonedSet := w.set.Clone()
	return From[T](clonedSet)
}

// Iterator returns an iter.Iterator over the elements in the set.
func (w *Wrapper[T]) Iterator() iter.Iterator[T] {
	return w.SetIterator()
}

// SetIterator returns an iterator over the elements in the set.
func (w *Wrapper[T]) SetIterator() sets.Iterator[T] {
	return Iterator[T]{w.set.SetIterator(), &w.mutex}
}

// Stream streams the elements of the Set.
func (w *Wrapper[T]) Stream(yield func(T) bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	w.set.Stream(yield)
}

// Transaction executes the given update function with the set as an argument.
// The set is locked for writing during the execution of the update function.
func (w *Wrapper[T]) Transaction(updateFunction func(set sets.Set[T])) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	updateFunction(w.set)
}
