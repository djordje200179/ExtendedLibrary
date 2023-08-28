package hashmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
	"unsafe"
)

// Map is a hash map with builtin map as a base.
type Map[K comparable, V any] map[K]V

// New creates an empty hash map.
func New[K comparable, V any]() Map[K, V] {
	return NewWithCapacity[K, V](0)
}

// NewWithCapacity creates an empty hash map with the specified capacity.
func NewWithCapacity[K comparable, V any](capacity int) Map[K, V] {
	return FromMap(make(map[K]V, capacity))
}

// FromMap creates a hash map from the specified map.
func FromMap[K comparable, V any](m map[K]V) Map[K, V] {
	return m
}

// Collector returns a collector that collects key-value pairs into an empty hash map.
func Collector[K comparable, V any]() streams.Collector[misc.Pair[K, V], Map[K, V]] {
	return maps.Collector[K, V, Map[K, V]]{New[K, V]()}
}

// Size returns the number of entries in the map.
func (hashmap Map[K, V]) Size() int {
	return len(hashmap)
}

// Contains returns true if the map contains the specified key.
func (hashmap Map[K, V]) Contains(key K) bool {
	_, ok := hashmap[key]
	return ok
}

// TryGet returns the value associated with the specified key,
// or zero value and false if the key is not present.
func (hashmap Map[K, V]) TryGet(key K) (V, bool) {
	value, ok := hashmap[key]
	return value, ok
}

// Get returns the value associated with the specified key.
// Panics if the key is not present.
func (hashmap Map[K, V]) Get(key K) V {
	value, ok := hashmap[key]
	if !ok {
		maps.PanicOnMissingKey(key)
	}

	return value
}

// GetRef returns a reference to the value associated with the specified key.
// Panics if the key is not present.
func (hashmap Map[K, V]) GetRef(key K) *V {
	mt, mv := mapTypeAndValue(hashmap)
	ptr, ok := internalMapGet(mt, mv, unsafe.Pointer(&key))

	if !ok {
		maps.PanicOnMissingKey(key)
	}

	return (*V)(ptr)
}

// Set sets the value associated with the specified key.
func (hashmap Map[K, V]) Set(key K, value V) {
	hashmap[key] = value
}

// Remove removes the entry with the specified key.
// Does nothing if the key is not present.
func (hashmap Map[K, V]) Remove(key K) {
	delete(hashmap, key)
}

// Keys returns a slice of all keys in the map.
func (hashmap Map[K, V]) Keys() []K {
	keys := make([]K, len(hashmap))

	i := 0
	for key := range hashmap {
		keys[i] = key
		i++
	}

	return keys
}

// Clear removes all entries from the map.
func (hashmap Map[K, V]) Clear() {
	clear(hashmap)
}

// Clone returns a shallow copy of the map.
func (hashmap Map[K, V]) Clone() maps.Map[K, V] {
	cloned := New[K, V]()
	for k, v := range hashmap {
		cloned[k] = v
	}

	return cloned
}

// Iterator returns an iterator over the map.
func (hashmap Map[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return hashmap.MapIterator()
}

// MapIterator returns an iterator over the map.
func (hashmap Map[K, V]) MapIterator() maps.Iterator[K, V] {
	return &Iterator[K, V]{
		m:     hashmap,
		keys:  hashmap.Keys(),
		index: 0,
	}
}

// Stream returns a stream over the map.
func (hashmap Map[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	supplier := suppliers.Map(hashmap)
	return streams.New(supplier)
}

// RefsStream returns a stream over references to the map values.
func (hashmap Map[K, V]) RefsStream() streams.Stream[misc.Pair[K, *V]] {
	return maps.RefsStream[K, V](hashmap)
}

// Map returns the builtin map used as a base.
func (hashmap Map[K, V]) Map() map[K]V {
	return hashmap
}
