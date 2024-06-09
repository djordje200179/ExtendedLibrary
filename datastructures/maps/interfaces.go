package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc"
)

// Iterator is a special iter.Iterator
// that can modify the Map it iterates over.
type Iterator[K, V any] interface {
	iter.Iterator[misc.Pair[K, V]]

	// Key returns the key of the current entry.
	Key() K
	// Value returns the value of the current entry.
	Value() V
	// ValueRef returns a reference to the value of the current entry.
	ValueRef() *V
	// SetValue sets the value of the current entry.
	SetValue(value V)

	// Remove removes the current entry.
	Remove()
}

// Map is a special iter.Iterable
// that represents a map of keys to values.
type Map[K, V any] interface {
	// Size returns the number of entries.
	Size() int

	// Contains returns true if the given key is present.
	Contains(key K) bool
	// TryGet returns the value associated with the
	// given key if it is present.
	// If the key is not present, it returns the zero value
	// for the value type and false.
	TryGet(key K) (V, bool)
	// Get returns the value associated with the given key.
	//
	// Panic occurs if the key is not present.
	Get(key K) V
	// GetRef returns a reference to the value associated with the given key.
	//
	// Panic occurs if the key is not present.
	GetRef(key K) *V
	// Set sets the value associated with the given key.
	// If the key is not present, it adds the entry.
	Set(key K, value V)
	// Remove removes the entry associated with the given key.
	//
	// If the key is not present, it does nothing.
	Remove(key K)

	// Clear removes all entries.
	Clear()

	misc.Cloner[Map[K, V]]

	iter.Iterable[misc.Pair[K, V]]
	// MapIterator returns an Iterator over the entries.
	// It can be used to modify the entries while iterating.
	//
	// The order of iteration is not guaranteed.
	MapIterator() Iterator[K, V]
	// Stream2 streams the entries.
	Stream2(yield func(K, V) bool)
	// Keys streams the keys.
	Keys(yield func(K) bool)
	// Values streams the values.
	Values(yield func(V) bool)
}
