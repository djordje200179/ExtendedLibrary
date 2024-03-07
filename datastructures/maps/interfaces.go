package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc"
)

// Iterator is an iterator over a Map.
type Iterator[K, V any] interface {
	iter.Iterator[misc.Pair[K, V]]

	// Key returns the key of the current entry.
	Key() K
	// Value returns the value of the current entry.
	Value() V
	// ValueRef returns a pointer to the value of the current entry.
	ValueRef() *V
	// SetValue sets the value of the current entry.
	SetValue(value V)

	// Remove removes the current entry.
	Remove()
}

// Map is an interface that represents a map of keys to values.
type Map[K, V any] interface {
	// Size returns the size of the map
	Size() int

	// Contains returns true if the map contains the given key
	Contains(key K) bool
	// TryGet returns the value associated with the given key, or zero value and false if the key is not present
	TryGet(key K) (V, bool)
	// Get returns the value associated with the given key
	Get(key K) V
	// GetRef returns a pointer to the value associated with the given key
	GetRef(key K) *V
	// Set sets the value associated with the given key to the given value
	Set(key K, value V)
	// Remove removes the entry associated with the given key
	Remove(key K)

	// Clear clears the map
	Clear()

	misc.Cloner[Map[K, V]]

	iter.Iterable[misc.Pair[K, V]]
	// MapIterator returns an iterator over the Map
	MapIterator() Iterator[K, V]
	// Stream2 streams the entries of the Map
	Stream2(yield func(K, V) bool)
	// Keys streams the keys of the Map
	Keys(yield func(K) bool)
	// Values streams the values of the Map
	Values(yield func(V) bool)
}
