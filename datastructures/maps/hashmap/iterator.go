package hashmap

import "github.com/djordje200179/extendedlibrary/misc"

type Iterator[K comparable, V any] struct {
	m Map[K, V]

	keys  []K
	index int
}

func (it *Iterator[K, V]) Valid() bool {
	return it.index < len(it.keys)
}

func (it *Iterator[K, V]) Move() {
	it.index++
}

func (it *Iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.Pair[K, V]{
		First:  it.Key(),
		Second: it.Value(),
	}
}

func (it *Iterator[K, V]) Key() K {
	return it.keys[it.index]
}

func (it *Iterator[K, V]) Value() V {
	return it.m.Get(it.Key())
}

func (it *Iterator[K, V]) ValueRef() *V {
	return it.m.GetRef(it.Key())
}

func (it *Iterator[K, V]) SetValue(value V) {
	it.m.Set(it.Key(), value)
}

func (it *Iterator[K, V]) Remove() {
	it.m.Remove(it.Key())
}
