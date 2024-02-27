package hashmap

import "github.com/djordje200179/extendedlibrary/misc"

// Iterator is an iterator over a HashMap.
type Iterator[K comparable, V any] struct {
	m Map[K, V]

	keys  []K
	index int
}

// Valid returns true if it points to a valid entry.
func (it *Iterator[K, V]) Valid() bool {
	return it.index < len(it.keys)
}

// Move moves the iterator to the next entry.
func (it *Iterator[K, V]) Move() {
	it.index++
}

// Get returns the current entry as a key-value pair.
func (it *Iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.MakePair(it.Key(), it.Value())
}

// Key returns the key of the current entry.
func (it *Iterator[K, V]) Key() K {
	return it.keys[it.index]
}

// Value returns the value of the current entry.
func (it *Iterator[K, V]) Value() V {
	return it.m.Get(it.Key())
}

// ValueRef returns a reference to the value of the current entry.
func (it *Iterator[K, V]) ValueRef() *V {
	return it.m.GetRef(it.Key())
}

// SetValue sets the value of the current entry.
func (it *Iterator[K, V]) SetValue(value V) {
	it.m.Set(it.Key(), value)
}

// Remove removes the current entry from the map.
// The iterator will point to the next entry afterward.
func (it *Iterator[K, V]) Remove() {
	it.m.Remove(it.Key())
}
