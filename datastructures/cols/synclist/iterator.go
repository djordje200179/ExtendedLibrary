package synclist

// Iterator is an iterator over the elements in a List.
type Iterator[T any] struct {
	curr *Node[T]
}

// Valid returns true if the iterator is positioned at a valid element.
func (it *Iterator[T]) Valid() bool {
	return it.curr != nil
}

// Move moves the iterator to the next element.
func (it *Iterator[T]) Move() {
	if it.curr == nil {
		return
	}

	it.curr = it.curr.next
}

// Get returns the element at the current position.
func (it *Iterator[T]) Get() T {
	return it.curr.Value
}
