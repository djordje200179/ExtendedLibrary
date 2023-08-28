package array

// Iterator is an iterator over an Array.
type Iterator[T any] struct {
	array *Array[T]
	index int
}

// Valid returns true if the iterator is currently pointing to a valid element.
func (it *Iterator[T]) Valid() bool {
	return it.index < it.array.Size()
}

// Move moves the iterator to the next element.
func (it *Iterator[T]) Move() {
	it.index++
}

// GetRef returns a reference to the current element.
func (it *Iterator[T]) GetRef() *T {
	return it.array.GetRef(it.index)
}

// Get returns the current element.
func (it *Iterator[T]) Get() T {
	return it.array.Get(it.index)
}

// Set sets the current element.
func (it *Iterator[T]) Set(value T) {
	it.array.Set(it.index, value)
}

// InsertBefore inserts the specified element before the current element.
// Iterator then points to the inserted element.
func (it *Iterator[T]) InsertBefore(value T) {
	it.array.Insert(it.index, value)
}

// InsertAfter inserts the specified element after the current element.
// Iterator keeps pointing to the current element.
func (it *Iterator[T]) InsertAfter(value T) {
	it.array.Insert(it.index+1, value)
}

// Remove removes the current element.
// Iterator then points to the next element.
func (it *Iterator[T]) Remove() {
	it.array.Remove(it.index)
}

// Index returns the current index.
func (it *Iterator[T]) Index() int {
	return it.index
}
