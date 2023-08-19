package bitset

type Iterator struct {
	index int

	set *Set
}

func (it *Iterator) Valid() bool {
	return it.index < it.set.arr.Size()
}

func (it *Iterator) Move() {
	for it.Valid() && it.set.arr.Get(it.index) == false {
		it.index++
	}
}

func (it *Iterator) Get() int {
	return it.index
}

func (it *Iterator) Remove() {
	it.set.Remove(it.index)
}
