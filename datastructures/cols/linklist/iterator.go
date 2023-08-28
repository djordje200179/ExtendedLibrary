package linklist

// Iterator is an iterator over a List.
type Iterator[T any] struct {
	list *List[T]

	curr  *Node[T]
	index int
}

// Valid returns true if the iterator is currently pointing to a valid element.
func (it *Iterator[T]) Valid() bool {
	return it.curr != nil
}

// Move moves the iterator to the next element.
func (it *Iterator[T]) Move() {
	if it.curr == nil {
		return
	}

	it.curr = it.curr.next
	it.index++
}

// GetRef returns a reference to the current element.
func (it *Iterator[T]) GetRef() *T {
	return &it.curr.Value
}

// Get returns the current element.
func (it *Iterator[T]) Get() T {
	return it.curr.Value
}

// Set sets the current element.
func (it *Iterator[T]) Set(value T) {
	it.curr.Value = value
}

// InsertBefore inserts the specified element before the current element.
// Iterator then points to the inserted element.
func (it *Iterator[T]) InsertBefore(value T) {
	it.curr.InsertBefore(value)
	it.curr = it.curr.prev
}

// InsertAfter inserts the specified element after the current element.
// Iterator keeps pointing to the current element.
func (it *Iterator[T]) InsertAfter(value T) {
	it.curr.InsertAfter(value)
}

// Remove removes the current element.
// Iterator then points to the next element.
func (it *Iterator[T]) Remove() {
	next := it.curr.next
	it.list.RemoveNode(it.curr)
	it.curr = next
}

// Node returns the current node.
func (it *Iterator[T]) Node() *Node[T] {
	return it.curr
}

// Index returns the current index.
func (it *Iterator[T]) Index() int {
	return it.index
}
