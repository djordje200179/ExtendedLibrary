package bitset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols/bitarray"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
)

// Set is a set implementation based on bit array.
type Set struct {
	arr *bitarray.Array

	elements int
}

// New creates an empty set with the given size.
func New(size int) *Set {
	return &Set{bitarray.NewWithSize(size), 0}
}

// NewFromIterable creates a set with the given size from the given iter.Iterable.
func NewFromIterable(size int, iterable iter.Iterable[int]) *Set {
	set := New(size)

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		set.Add(it.Get())
	}

	return set
}

// FromArray creates a set from the given array.
func FromArray(arr *bitarray.Array) *Set {
	return &Set{arr, arr.Count()}
}

// Size returns the number of elements in the Set.
func (s *Set) Size() int {
	return s.elements
}

// Add adds the given value to the Set.
func (s *Set) Add(value int) {
	if !s.Contains(value) {
		s.arr.Set(value, true)
		s.elements++
	}
}

// Remove removes the given value from the Set.
// If the value is not in the Set, this method does nothing.
func (s *Set) Remove(value int) {
	if s.Contains(value) {
		s.arr.Set(value, false)
		s.elements--
	}
}

// Contains returns true if the Set contains the given value.
func (s *Set) Contains(value int) bool {
	return s.arr.Get(value)
}

// Clear removes all elements from the Set.
func (s *Set) Clear() {
	s.arr.Clear()
	s.elements = 0
}

// Clone returns a shallow copy of the Set.
func (s *Set) Clone() sets.Set[int] {
	clonedArray := s.arr.Clone()
	return &Set{clonedArray, s.elements}
}

// Iterator returns an iter.Iterator over the elements in the Set.
func (s *Set) Iterator() iter.Iterator[int] {
	return s.SetIterator()
}

// SetIterator returns an iterator over the elements in the Set.
func (s *Set) SetIterator() sets.Iterator[int] {
	return &Iterator{0, s}
}

// Stream streams the elements of the Set.
func (s *Set) Stream(yield func(int) bool) {
	for i := range s.arr.Size() {
		if !s.arr.Get(i) {
			continue
		}

		if !yield(i) {
			break
		}
	}
}

// Array returns the underlying bit array.
func (s *Set) Array() *bitarray.Array {
	return s.arr
}
