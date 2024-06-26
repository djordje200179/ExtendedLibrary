package readmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
)

// Wrapper is a wrapper around a
// maps.Map that provides read-only access to it.
type Wrapper[K, V any] struct {
	m maps.Map[K, V]
}

// From creates a new Wrapper from the given maps.Map.
func From[K, V any](m maps.Map[K, V]) Wrapper[K, V] {
	return Wrapper[K, V]{m}
}

// Size returns the number of entries.
func (w Wrapper[K, V]) Size() int {
	return w.m.Size()
}

// Contains returns true if the given key is present.
func (w Wrapper[K, V]) Contains(key K) bool {
	return w.m.Contains(key)
}

// TryGet returns the value associated with the
// given key if it is present.
// If the key is not present, it returns the zero value
// for the value type and false.
func (w Wrapper[K, V]) TryGet(key K) (V, bool) {
	return w.m.TryGet(key)
}

// Get returns the value associated with the given key.
func (w Wrapper[K, V]) Get(key K) V {
	return w.m.Get(key)
}

// Clone returns a new Wrapper with
// a clone of the underlying maps.Map.
func (w Wrapper[K, V]) Clone() Wrapper[K, V] {
	clonedMap := w.m.Clone()
	return Wrapper[K, V]{clonedMap}
}

// Iterator returns an iter.Iterator over the entries.
func (w Wrapper[K, V]) Iterator() iter.Iterator[misc.Pair[K, V]] {
	return w.m.Iterator()
}

// Stream2 streams all entries.
func (w Wrapper[K, V]) Stream2(yield func(K, V) bool) {
	w.m.Stream2(yield)
}

// Keys streams the keys.
func (w Wrapper[K, V]) Keys(yield func(K) bool) {
	w.m.Keys(yield)
}

// Values streams the values.
func (w Wrapper[K, V]) Values(yield func(V) bool) {
	w.m.Values(yield)
}
