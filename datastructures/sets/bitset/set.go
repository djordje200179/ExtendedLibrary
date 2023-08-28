package bitset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols/bitarray"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/streams"
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

// NewFromIterable creates a set with the given size from the given iterable.
func NewFromIterable(size int, iter iterable.Iterable[int]) *Set {
	set := New(size)

	for it := iter.Iterator(); it.Valid(); it.Move() {
		set.Add(it.Get())
	}

	return set
}

// FromArray creates a set from the given array.
func FromArray(arr *bitarray.Array) *Set {
	return &Set{arr, arr.Count()}
}

// Size returns the number of elements in the Set.
func (set *Set) Size() int {
	return set.elements
}

// Add adds the given value to the Set.
func (set *Set) Add(value int) {
	if !set.Contains(value) {
		set.arr.Set(value, true)
		set.elements++
	}
}

// Remove removes the given value from the Set.
// If the value is not in the Set, this method does nothing.
func (set *Set) Remove(value int) {
	if set.Contains(value) {
		set.arr.Set(value, false)
		set.elements--
	}
}

// Contains returns true if the Set contains the given value.
func (set *Set) Contains(value int) bool {
	return set.arr.Get(value)
}

// Clear removes all elements from the Set.
func (set *Set) Clear() {
	set.arr.Clear()
	set.elements = 0
}

// Clone returns a shallow copy of the Set.
func (set *Set) Clone() sets.Set[int] {
	clonedArray := set.arr.Clone()
	return &Set{clonedArray, set.elements}
}

// Iterator returns an iterator over the elements in the Set.
func (set *Set) Iterator() iterable.Iterator[int] {
	return set.SetIterator()
}

// SetIterator returns an iterator over the elements in the Set.
func (set *Set) SetIterator() sets.Iterator[int] {
	return &Iterator{0, set}
}

// Stream returns a stream over the elements in the Set.
func (set *Set) Stream() streams.Stream[int] {
	return iterable.IteratorStream(set.Iterator())
}

// Array returns the underlying bit array.
func (set *Set) Array() *bitarray.Array {
	return set.arr
}
