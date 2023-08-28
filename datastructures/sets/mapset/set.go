package mapset

import (
	"cmp"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/rbt"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type empty struct{}

// Set is a set implementation based on maps.
// Keys are elements of the set, and values are empty structs.
type Set[T any] struct {
	m maps.Map[T, empty]
}

// NewHashSet creates a new hash set for comparable types.
func NewHashSet[T comparable]() Set[T] {
	return FromMap[T](hashmap.New[T, empty]())
}

// NewTreeSet creates a new tree set for ordered types.
func NewTreeSet[T cmp.Ordered]() Set[T] {
	return FromMap[T](rbt.New[T, empty]())
}

// FromMap creates a new set from a map.
func FromMap[T any](m maps.Map[T, empty]) Set[T] {
	return Set[T]{m}
}

// Size returns the number of elements in the set.
func (set Set[T]) Size() int {
	return set.m.Size()
}

// Add adds a value to the set.
func (set Set[T]) Add(value T) {
	if !set.m.Contains(value) {
		set.m.Set(value, empty{})
	}
}

// Remove removes a value from the set.
func (set Set[T]) Remove(value T) {
	set.m.Remove(value)
}

// Contains returns true if the set contains the value.
func (set Set[T]) Contains(value T) bool {
	return set.m.Contains(value)
}

// Clear removes all elements from the set.
func (set Set[T]) Clear() {
	set.m.Clear()
}

// Clone returns a shallow copy of the set.
func (set Set[T]) Clone() sets.Set[T] {
	return FromMap[T](set.m.Clone())
}

// Iterator returns an iter.Iterator over the set.
func (set Set[T]) Iterator() iter.Iterator[T] {
	return set.SetIterator()
}

// SetIterator returns an iterator over the set.
func (set Set[T]) SetIterator() sets.Iterator[T] {
	return Iterator[T]{set.m.MapIterator()}
}

// Stream returns a streams.Stream of the set elements.
func (set Set[T]) Stream() streams.Stream[T] {
	return streams.Map(set.m.Stream(), func(pair misc.Pair[T, empty]) T {
		return pair.First
	})
}

// Map returns the underlying map.
func (set Set[T]) Map() maps.Map[T, empty] {
	return set.m
}
