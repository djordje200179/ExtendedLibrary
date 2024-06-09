package bitset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols/bitarray"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
)

// Set stores information about the presence
// of values in the 0..N-1 range.
type Set struct {
	arr *bitarray.Array

	elements int
}

// New creates an empty Set with the given size.
func New(size int) *Set {
	return &Set{bitarray.NewWithSize(size), 0}
}

// NewFromIterable creates a new Set with the given size
// and elements from the given iter.Iterable.
func NewFromIterable(size int, iterable iter.Iterable[int]) *Set {
	set := New(size)

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		set.Add(it.Get())
	}

	return set
}

// FromArray creates a new Set from the given bitarray.Array.
func FromArray(arr *bitarray.Array) *Set {
	return &Set{arr, arr.Count()}
}

// Size returns the cardinality.
func (s *Set) Size() int {
	return s.elements
}

// Add inserts the given value.
// If the value is already present, this method does nothing.
//
// Panic cols.IndexOutOfBoundsError occurs
// if the value is out of bounds.
func (s *Set) Add(value int) {
	if !s.Contains(value) {
		s.arr.Set(value, true)
		s.elements++
	}
}

// Remove removes the given value.
// If the value is not present, this method does nothing.
//
// Panic cols.IndexOutOfBoundsError occurs
// if the value is out of bounds.
func (s *Set) Remove(value int) {
	if s.Contains(value) {
		s.arr.Set(value, false)
		s.elements--
	}
}

// Contains returns true if the value is already present.
//
// Panic cols.IndexOutOfBoundsError occurs
// if the value is out of bounds.
func (s *Set) Contains(value int) bool {
	return s.arr.Get(value)
}

// Clear removes all the values.
func (s *Set) Clear() {
	s.arr.Clear()
	s.elements = 0
}

// Clone returns a new Set with the same values.
func (s *Set) Clone() sets.Set[int] {
	clonedArray := s.arr.Clone()
	return &Set{clonedArray, s.elements}
}

// Iterator returns a read-only iter.Iterator over the elements.
func (s *Set) Iterator() iter.Iterator[int] {
	return s.SetIterator()
}

// SetIterator returns a specialized Iterator over the elements.
func (s *Set) SetIterator() sets.Iterator[int] {
	return &Iterator{0, s}
}

// Stream streams the elements.
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

// Array returns the underlying bitarray.Array.
func (s *Set) Array() *bitarray.Array {
	return s.arr
}
