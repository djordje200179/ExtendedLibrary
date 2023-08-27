package bitset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections/bitarray"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Set struct {
	arr *bitarray.Array

	elements int
}

func New(size int) *Set {
	return &Set{bitarray.NewWithSize(size), 0}
}

func FromArray(arr *bitarray.Array) *Set {
	elements := 0
	return &Set{arr, elements}
}

func Collector(size int) streams.Collector[int, *Set] {
	return sets.Collector[int, *Set]{New(size)}
}

func (set *Set) Size() int {
	return set.elements
}

func (set *Set) Add(value int) {
	if !set.Contains(value) {
		set.arr.Set(value, true)
		set.elements++
	}
}

func (set *Set) Remove(value int) {
	if set.Contains(value) {
		set.arr.Set(value, false)
		set.elements--
	}
}

func (set *Set) Contains(value int) bool {
	return set.arr.Get(value)
}

func (set *Set) Clear() {
	set.arr.Clear()
	set.elements = 0
}

func (set *Set) Clone() sets.Set[int] {
	clonedArray := set.arr.Clone()
	return &Set{clonedArray, set.elements}
}

func (set *Set) Iterator() iterable.Iterator[int] {
	return set.SetIterator()
}

func (set *Set) SetIterator() sets.Iterator[int] {
	return &Iterator{0, set}
}

func (set *Set) Stream() streams.Stream[int] {
	return iterable.IteratorStream(set.Iterator())
}

func (set *Set) Array() *bitarray.Array {
	return set.arr
}
