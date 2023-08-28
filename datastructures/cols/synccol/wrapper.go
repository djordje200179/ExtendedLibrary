package synccol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

// Wrapper is a wrapper around a collection that provides thread-safe access to the collection.
// Locking is done through read-write mutex. This means that multiple goroutines can read from the collection at the same time,
// but only one goroutine can write to the collection at the same time.
type Wrapper[T any] struct {
	collection cols.Collection[T]

	mutex sync.RWMutex
}

// From creates a new Wrapper from the given collection.
func From[T any](collection cols.Collection[T]) *Wrapper[T] {
	return &Wrapper[T]{collection, sync.RWMutex{}}
}

// Size returns the number of elements in the collection.
func (wrapper *Wrapper[T]) Size() int {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.Size()
}

// Get returns the element at the given index.
func (wrapper *Wrapper[T]) Get(index int) T {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.Get(index)
}

// GetRef returns a reference to the element at the given index.
// Usage of this method is discouraged, as it breaks the thread-safety of the collection.
// Lock will not be held while the reference is used, so it is possible that the value of the element changes while the reference is used.
func (wrapper *Wrapper[T]) GetRef(index int) *T {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.GetRef(index)
}

// Set sets the value of the element at the given index.
func (wrapper *Wrapper[T]) Set(index int, value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Set(index, value)
}

// Update updates the value of the element at the given index.
func (wrapper *Wrapper[T]) Update(index int, updateFunction func(value T) T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.collection.Get(index)
	newValue := updateFunction(oldValue)
	wrapper.collection.Set(index, newValue)
}

// UpdateRef updates the value of the element at the given index.
func (wrapper *Wrapper[T]) UpdateRef(index int, updateFunction func(value *T)) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.collection.GetRef(index)
	updateFunction(oldValue)
}

// Prepend prepends the given value to the collection.
func (wrapper *Wrapper[T]) Prepend(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Prepend(value)
}

// Append appends the given value to the collection.
func (wrapper *Wrapper[T]) Append(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Append(value)
}

// Insert inserts the given value at the given index.
func (wrapper *Wrapper[T]) Insert(index int, value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Insert(index, value)
}

// Remove removes the element at the given index.
func (wrapper *Wrapper[T]) Remove(index int) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Remove(index)
}

// Clear removes all elements from the collection.
func (wrapper *Wrapper[T]) Clear() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Clear()
}

// Reverse reverses the order of elements in the collection.
func (wrapper *Wrapper[T]) Reverse() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Reverse()
}

// Sort sorts the elements in the collection using the given comparator.
func (wrapper *Wrapper[T]) Sort(comparator comparison.Comparator[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Sort(comparator)
}

// Join appends all elements from the given collection to the end of the collection.
func (wrapper *Wrapper[T]) Join(other cols.Collection[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Join(other)
}

// Clone returns a shallow copy of a Wrapper.
// Cloned Wrapper will have the same underlying collection as the original Wrapper.
func (wrapper *Wrapper[T]) Clone() cols.Collection[T] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	clonedCollection := wrapper.collection.Clone()
	return &Wrapper[T]{clonedCollection, sync.RWMutex{}}
}

// Iterator returns an iterator over the elements in the collection.
func (wrapper *Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.CollectionIterator()
}

// CollectionIterator returns an iterator over the elements in the collection.
func (wrapper *Wrapper[T]) CollectionIterator() cols.Iterator[T] {
	return Iterator[T]{wrapper.collection.CollectionIterator(), &wrapper.mutex}
}

// Stream returns a stream over the elements in the collection.
func (wrapper *Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.collection.Stream()
}

// RefsStream returns a stream over the references to the elements in the collection.
func (wrapper *Wrapper[T]) RefsStream() streams.Stream[*T] {
	return wrapper.collection.RefsStream()
}

// FindIndex returns the index of the first element that satisfies the given predicate.
func (wrapper *Wrapper[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.FindIndex(predicate)
}

// FindRef returns a reference to the first element that satisfies the given predicate.
// Usage of this method is discouraged, as it breaks the thread-safety of the collection.
// Lock will not be held while the reference is used, so it is possible that the value of the element changes while the reference is used.
func (wrapper *Wrapper[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.FindRef(predicate)
}

// Transaction executes the given function with the collection as an argument.
// The collection will be locked for writing while the function is executed.
func (wrapper *Wrapper[T]) Transaction(updateFunction func(collection cols.Collection[T])) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	updateFunction(wrapper.collection)
}
