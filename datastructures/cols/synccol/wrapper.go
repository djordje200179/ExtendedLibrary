package synccol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
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
func (w *Wrapper[T]) Size() int {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.collection.Size()
}

// Get returns the element at the given index.
func (w *Wrapper[T]) Get(index int) T {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.collection.Get(index)
}

// GetRef returns a reference to the element at the given index.
// Usage of this method is discouraged, as it breaks the thread-safety.
// Lock will not be held while the reference is used, so it is possible
// that the value of the element changes while the reference is used.
func (w *Wrapper[T]) GetRef(index int) *T {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.collection.GetRef(index)
}

// Set sets the value of the element at the given index.
func (w *Wrapper[T]) Set(index int, value T) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Set(index, value)
}

// Update calculates and sets the new value of the element at the given index.
func (w *Wrapper[T]) Update(index int, updateFunction func(value T) T) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	oldValue := w.collection.Get(index)
	newValue := updateFunction(oldValue)
	w.collection.Set(index, newValue)
}

// UpdateRef in-place updates the value of the element at the given index.
func (w *Wrapper[T]) UpdateRef(index int, updateFunction func(value *T)) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	oldValue := w.collection.GetRef(index)
	updateFunction(oldValue)
}

// Prepend prepends the given value to the collection.
func (w *Wrapper[T]) Prepend(value T) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Prepend(value)
}

// Append appends the given value to the collection.
func (w *Wrapper[T]) Append(value T) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Append(value)
}

// Insert inserts the given value at the given index.
func (w *Wrapper[T]) Insert(index int, value T) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Insert(index, value)
}

// Remove removes the element at the given index.
func (w *Wrapper[T]) Remove(index int) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Remove(index)
}

// Clear clears the collection.
func (w *Wrapper[T]) Clear() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Clear()
}

// Reverse reverses the collection.
func (w *Wrapper[T]) Reverse() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Reverse()
}

// Sort sorts the elements using the given comparator.
func (w *Wrapper[T]) Sort(comparator comparison.Comparator[T]) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Sort(comparator)
}

// Join joins the collection with the given collection.
func (w *Wrapper[T]) Join(other cols.Collection[T]) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.Join(other)
}

// Clone returns a copy of a Wrapper with the new
// underlying collection that is also cloned.
func (w *Wrapper[T]) Clone() cols.Collection[T] {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	clonedCollection := w.collection.Clone()
	return &Wrapper[T]{clonedCollection, sync.RWMutex{}}
}

// Iterator returns an iter.Iterator over the elements.
func (w *Wrapper[T]) Iterator() iter.Iterator[T] {
	return w.CollectionIterator()
}

// CollectionIterator returns a cols.Iterator over the elements.
func (w *Wrapper[T]) CollectionIterator() cols.Iterator[T] {
	return Iterator[T]{w.collection.CollectionIterator(), &w.mutex}
}

// Stream streams the elements of the collection.
func (w *Wrapper[T]) Stream(yield func(T) bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	w.collection.Stream(yield)
}

// RefsStream streams the references to the elements of the collection.
func (w *Wrapper[T]) RefsStream(yield func(*T) bool) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.collection.RefsStream(yield)
}

// FindIndex returns the index of the first element that satisfies the given predicate.
// If no element satisfies the predicate, 0 and false are returned.
func (w *Wrapper[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.collection.FindIndex(predicate)
}

// FindRef returns a reference to the first element that satisfies the given predicate.
// If no element satisfies the predicate, nil and false are returned.
// Lock will not be held while the reference is used, so it is possible
// that the value of the element changes while the reference is used.
func (w *Wrapper[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.collection.FindRef(predicate)
}

// Transaction executes the given function with the cols.Collection as an argument.
// Wrapped cols.Collection will be locked for writing while the function is executed.
func (w *Wrapper[T]) Transaction(updateFunction func(collection cols.Collection[T])) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	updateFunction(w.collection)
}
