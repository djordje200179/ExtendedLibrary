package hashmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"unsafe"
)

// Map is a hash map with builtin map as a base.
type Map[K comparable, V any] map[K]V

// New creates an empty Map.
func New[K comparable, V any]() Map[K, V] {
	return NewWithCapacity[K, V](0)
}

// NewWithCapacity creates an empty Map with the specified capacity.
func NewWithCapacity[K comparable, V any](capacity int) Map[K, V] {
	return FromMap(make(map[K]V, capacity))
}

// NewFromIterable creates a Map from the specified iter.Iterable.
func NewFromIterable[K comparable, V any](iterable iter.Iterable[misc.Pair[K, V]]) Map[K, V] {
	var m Map[K, V]

	if finiteIter, ok := any(iterable).(iter.FiniteIterable[misc.Pair[K, V]]); ok {
		m = NewWithCapacity[K, V](finiteIter.Size())
	} else {
		m = New[K, V]()
	}

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		entry := it.Get()

		m[entry.First] = entry.Second
	}

	return m
}

// FromMap creates a new Map from the specified map.
func FromMap[K comparable, V any](m map[K]V) Map[K, V] {
	return m
}

// Size returns the number of entries in the map.
func (m Map[K, V]) Size() int {
	return len(m)
}

// Contains returns true if the map contains the specified key.
func (m Map[K, V]) Contains(key K) bool {
	_, ok := m[key]
	return ok
}

// TryGet returns the value associated with the specified key,
// or zero value and false if the key is not present.
func (m Map[K, V]) TryGet(key K) (V, bool) {
	value, ok := m[key]
	return value, ok
}

// Get returns the value associated with the specified key.
// Panics if the key is not present.
func (m Map[K, V]) Get(key K) V {
	value, ok := m[key]
	if !ok {
		panic(maps.ErrMissingKey[K]{Key: key})
	}

	return value
}

// GetRef returns a reference to the value associated with the specified key.
// Panics if the key is not present.
func (m Map[K, V]) GetRef(key K) *V {
	mt, mv := mapTypeAndValue(m)
	ptr, ok := internalMapGet(mt, mv, unsafe.Pointer(&key))

	if !ok {
		panic(maps.ErrMissingKey[K]{Key: key})
	}

	return (*V)(ptr)
}

// Set sets the value associated with the specified key.
func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

// Remove removes the entry with the specified key.
// Does nothing if the key is not present.
func (m Map[K, V]) Remove(key K) {
	delete(m, key)
}

// Clear removes all entries from the map.
func (m Map[K, V]) Clear() {
	clear(m)
}

// Clone returns a shallow copy of the map.
func (m Map[K, V]) Clone() maps.Map[K, V] {
	cloned := New[K, V]()
	for k, v := range m {
		cloned[k] = v
	}

	return cloned
}

// Iterator returns an iter.Iterator over the map.
func (m Map[K, V]) Iterator() iter.Iterator[misc.Pair[K, V]] {
	return m.MapIterator()
}

// MapIterator returns an iterator over the map.
func (m Map[K, V]) MapIterator() maps.Iterator[K, V] {
	// TODO: Use builtin function
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return &Iterator[K, V]{
		m:     m,
		keys:  keys,
		index: 0,
	}
}

// Stream2 streams over the entries in the Map.
func (m Map[K, V]) Stream2(yield func(K, V) bool) {
	for k, v := range m {
		if !yield(k, v) {
			break
		}
	}
}

// RefsStream2 streams over the keys and references to the values in the Map.
func (m Map[K, V]) RefsStream2(yield func(K, *V) bool) {
	for it := m.MapIterator(); it.Valid(); it.Move() {
		if !yield(it.Key(), it.ValueRef()) {
			break
		}
	}
}

// Map returns the builtin map used as a base.
func (m Map[K, V]) Map() map[K]V {
	return m
}
