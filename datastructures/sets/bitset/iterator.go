package bitset

// Iterator is an iterator over a Set.
type Iterator struct {
	index int

	set *Set
}

// Valid returns true if the iterator is positioned at a valid element.
func (it *Iterator) Valid() bool {
	return it.index < it.set.arr.Size()
}

// Move moves the iterator to the next element.
func (it *Iterator) Move() {
	for it.Valid() && it.set.arr.Get(it.index) == false {
		it.index++
	}
}

// Get returns the current element.
func (it *Iterator) Get() int {
	return it.index
}

// Remove removes the current element from the Set.
func (it *Iterator) Remove() {
	it.set.Remove(it.index)
}
