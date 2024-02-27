package linkmap

import "github.com/djordje200179/extendedlibrary/misc"

// Iterator is a struct for iterating over the map.
type Iterator[K, V any] struct {
	wrapper *Wrapper[K, V]

	curr *Node[K, V]
}

// Valid returns true if the iterator points to a valid element.
func (it *Iterator[K, V]) Valid() bool {
	return it.curr != nil
}

// Move moves the iterator to the next element.
func (it *Iterator[K, V]) Move() {
	if it.curr == nil {
		return
	}

	it.curr = it.curr.next
}

// Get returns the current entry as a Pair.
func (it *Iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.MakePair(it.Key(), it.Value())
}

// Key returns the current key.
func (it *Iterator[K, V]) Key() K {
	return it.curr.key
}

// Value returns the value of the current entry.
func (it *Iterator[K, V]) Value() V {
	return it.curr.Value
}

// ValueRef returns a reference to the value of the current entry.
func (it *Iterator[K, V]) ValueRef() *V {
	return &it.curr.Value
}

// SetValue sets the value of the current entry.
func (it *Iterator[K, V]) SetValue(value V) {
	it.curr.Value = value
}

// Remove removes the current entry from the map.
// The iterator will be moved to the next element.
func (it *Iterator[K, V]) Remove() {
	next := it.curr.next
	it.wrapper.m.Remove(it.curr.key)
	it.curr = next
}

// Node returns the current node.
func (it *Iterator[K, V]) Node() *Node[K, V] {
	return it.curr
}
