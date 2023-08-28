package synccol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

// Wrapper is a wrapper around a cols.Collection that provides thread-safe access.
// Locking is done through read-write mutex. This means that multiple goroutines
// can read at the same time, but only one goroutine can write at the same time.
type Wrapper[T any] struct {
	collection cols.Collection[T]

	mutex sync.RWMutex
}

// From creates a new Wrapper around the given cols.Collection.
func From[T any](collection cols.Collection[T]) *Wrapper[T] {
	return &Wrapper[T]{collection, sync.RWMutex{}}
}

// Size returns the number of elements.
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
// Usage of this method is discouraged, as it breaks the thread-safety.
// Lock will not be held while the reference is used, so it is possible
// that the value of the element changes while the reference is used.
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

// Update calculates and sets the new value of the element at the given index.
func (wrapper *Wrapper[T]) Update(index int, updateFunction func(value T) T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.collection.Get(index)
	newValue := updateFunction(oldValue)
	wrapper.collection.Set(index, newValue)
}

// UpdateRef in-place updates the value of the element at the given index.
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

// Clear clears the collection.
func (wrapper *Wrapper[T]) Clear() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Clear()
}

// Reverse reverses the collection.
func (wrapper *Wrapper[T]) Reverse() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Reverse()
}

// Sort sorts the elements using the given comparator.
func (wrapper *Wrapper[T]) Sort(comparator comparison.Comparator[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Sort(comparator)
}

// Join joins the collection with the given collection.
func (wrapper *Wrapper[T]) Join(other cols.Collection[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Join(other)
}

// Clone returns a copy of a Wrapper with the new
// underlying collection that is also cloned.
func (wrapper *Wrapper[T]) Clone() cols.Collection[T] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	clonedCollection := wrapper.collection.Clone()
	return &Wrapper[T]{clonedCollection, sync.RWMutex{}}
}

// Iterator returns an iter.Iterator over the elements.
func (wrapper *Wrapper[T]) Iterator() iter.Iterator[T] {
	return wrapper.CollectionIterator()
}

// CollectionIterator returns a cols.Iterator over the elements.
func (wrapper *Wrapper[T]) CollectionIterator() cols.Iterator[T] {
	return Iterator[T]{wrapper.collection.CollectionIterator(), &wrapper.mutex}
}

// Stream returns a streams.Stream over the elements.
func (wrapper *Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.collection.Stream()
}

// RefsStream returns a streams.Stream over the references to the elements.
func (wrapper *Wrapper[T]) RefsStream() streams.Stream[*T] {
	return wrapper.collection.RefsStream()
}

// FindIndex returns the index of the first element that satisfies the given predicate.
// If no element satisfies the predicate, 0 and false are returned.
func (wrapper *Wrapper[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.FindIndex(predicate)
}

// FindRef returns a reference to the first element that satisfies the given predicate.
// If no element satisfies the predicate, nil and false are returned.
// Lock will not be held while the reference is used, so it is possible
// that the value of the element changes while the reference is used.
func (wrapper *Wrapper[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.FindRef(predicate)
}

// Transaction executes the given function with the cols.Collection as an argument.
// Wrapped cols.Collection will be locked for writing while the function is executed.
func (wrapper *Wrapper[T]) Transaction(updateFunction func(collection cols.Collection[T])) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	updateFunction(wrapper.collection)
}
