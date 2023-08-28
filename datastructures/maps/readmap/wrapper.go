package readmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Wrapper is a wrapper around map that provides read-only access to the map.
type Wrapper[K, V any] struct {
	m maps.Map[K, V]
}

// From creates a new Wrapper from the given map.
func From[K, V any](m maps.Map[K, V]) Wrapper[K, V] {
	return Wrapper[K, V]{m}
}

// Size returns the number of entries in the map.
func (wrapper Wrapper[K, V]) Size() int {
	return wrapper.m.Size()
}

// Contains returns true if the map contains the given key.
func (wrapper Wrapper[K, V]) Contains(key K) bool {
	return wrapper.m.Contains(key)
}

// TryGet returns the value associated with the given key and true if the key is present in the map.
// Otherwise, it returns the zero value for the value type and false.
func (wrapper Wrapper[K, V]) TryGet(key K) (V, bool) {
	return wrapper.m.TryGet(key)
}

// Get returns the value associated with the given key.
func (wrapper Wrapper[K, V]) Get(key K) V {
	return wrapper.m.Get(key)
}

// Keys returns a slice of all keys in the map.
func (wrapper Wrapper[K, V]) Keys() []K {
	return wrapper.m.Keys()
}

// Clone returns a shallow copy of a Wrapper.
// Cloned Wrapper will have the same underlying map as the original Wrapper.
func (wrapper Wrapper[K, V]) Clone() Wrapper[K, V] {
	clonedMap := wrapper.m.Clone()
	return Wrapper[K, V]{clonedMap}
}

// Iterator returns an iter.Iterator over the entries in the map.
func (wrapper Wrapper[K, V]) Iterator() iter.Iterator[misc.Pair[K, V]] {
	return wrapper.m.Iterator()
}

// Stream returns a stream over the entries in the map.
func (wrapper Wrapper[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return wrapper.m.Stream()
}
