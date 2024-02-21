package mapset

import (
	"cmp"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/rbt"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
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
func (s Set[T]) Size() int {
	return s.m.Size()
}

// Add adds a value to the set.
func (s Set[T]) Add(value T) {
	if !s.m.Contains(value) {
		s.m.Set(value, empty{})
	}
}

// Remove removes a value from the set.
func (s Set[T]) Remove(value T) {
	s.m.Remove(value)
}

// Contains returns true if the set contains the value.
func (s Set[T]) Contains(value T) bool {
	return s.m.Contains(value)
}

// Clear removes all elements from the set.
func (s Set[T]) Clear() {
	s.m.Clear()
}

// Clone returns a shallow copy of the set.
func (s Set[T]) Clone() sets.Set[T] {
	return FromMap[T](s.m.Clone())
}

// Iterator returns an iter.Iterator over the set.
func (s Set[T]) Iterator() iter.Iterator[T] {
	return s.SetIterator()
}

// SetIterator returns an iterator over the set.
func (s Set[T]) SetIterator() sets.Iterator[T] {
	return Iterator[T]{s.m.MapIterator()}
}

// Stream streams the elements of the Set.
func (s Set[T]) Stream(yield func(T) bool) {
	for k, _ := range s.m.Stream2 {
		if !yield(k) {
			return
		}
	}
}

// Map returns the underlying map.
func (s Set[T]) Map() maps.Map[T, empty] {
	return s.m
}
