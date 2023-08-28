package syncset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/streams"
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
func (wrapper *Wrapper[T]) Size() int {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.set.Size()
}

// Add adds the given value to the set.
func (wrapper *Wrapper[T]) Add(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.set.Add(value)
}

// Remove removes the given value from the set.
func (wrapper *Wrapper[T]) Remove(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.set.Remove(value)
}

// Contains returns true if the set contains the given value.
func (wrapper *Wrapper[T]) Contains(value T) bool {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.set.Contains(value)
}

// Clear removes all elements from the set.
func (wrapper *Wrapper[T]) Clear() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.set.Clear()
}

// Clone returns a shallow copy of the set.
func (wrapper *Wrapper[T]) Clone() sets.Set[T] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	clonedSet := wrapper.set.Clone()
	return From[T](clonedSet)
}

// Iterator returns an iter.Iterator over the elements in the set.
func (wrapper *Wrapper[T]) Iterator() iter.Iterator[T] {
	return wrapper.SetIterator()
}

// SetIterator returns an iterator over the elements in the set.
func (wrapper *Wrapper[T]) SetIterator() sets.Iterator[T] {
	return Iterator[T]{wrapper.set.SetIterator(), &wrapper.mutex}
}

// Stream returns a streams.Stream over the elements in the set.
func (wrapper *Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.set.Stream()
}

// Transaction executes the given update function with the set as an argument.
// The set is locked for writing during the execution of the update function.
func (wrapper *Wrapper[T]) Transaction(updateFunction func(set sets.Set[T])) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	updateFunction(wrapper.set)
}
