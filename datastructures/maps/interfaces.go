package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Iterator is an iterator over a Map.
type Iterator[K, V any] interface {
	iterable.Iterator[misc.Pair[K, V]]

	Key() K           // Key returns the key of the current entry.
	Value() V         // Value returns the value of the current entry.
	ValueRef() *V     // ValueRef returns a pointer to the value of the current entry.
	SetValue(value V) // SetValue sets the value of the current entry.

	Remove() // Remove removes the current entry.
}

// Map is an interface that represents a map of keys to values.
type Map[K, V any] interface {
	Size() int // Size returns the size of the map

	Contains(key K) bool    // Contains returns true if the map contains the given key
	TryGet(key K) (V, bool) // TryGet returns the value associated with the given key, or zero value and false if the key is not present
	Get(key K) V            // Get returns the value associated with the given key
	GetRef(key K) *V        // GetRef returns a pointer to the value associated with the given key
	Set(key K, value V)     // Set sets the value associated with the given key to the given value
	Remove(key K)           // Remove removes the entry associated with the given key

	Keys() []K // Keys returns a slice of all keys in the map

	Clear() // Clear clears the map

	misc.Cloner[Map[K, V]]

	iterable.Iterable[misc.Pair[K, V]]
	MapIterator() Iterator[K, V] // MapIterator returns an iterator over the map
	streams.Streamer[misc.Pair[K, V]]
	RefsStream() streams.Stream[misc.Pair[K, *V]] // RefsStream returns a stream keys and pointers to the values of the map
}
