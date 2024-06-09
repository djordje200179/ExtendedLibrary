package synclist

// Iterator is an iterator over a List.
type Iterator[T any] struct {
	curr *Node[T]
}

// Valid returns true if the iterator is
// currently pointing to a valid element.
func (it *Iterator[T]) Valid() bool { return it.curr != nil }

// Move moves to the next element.
func (it *Iterator[T]) Move() {
	if it.curr == nil {
		return
	}

	it.curr = it.curr.next
}

// Get returns the current element.
func (it *Iterator[T]) Get() T { return it.curr.Value }
