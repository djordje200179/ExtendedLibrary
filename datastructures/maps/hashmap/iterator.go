package hashmap

import "github.com/djordje200179/extendedlibrary/misc"

type iterator[K comparable, V any] struct {
	m Map[K, V]

	keys  []K
	index int
}

func (it *iterator[K, V]) Valid() bool {
	return it.index < len(it.keys)
}

func (it *iterator[K, V]) Move() {
	it.index++
}

func (it *iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.Pair[K, V]{
		First:  it.Key(),
		Second: it.Value(),
	}
}

func (it *iterator[K, V]) Key() K {
	return it.keys[it.index]
}

func (it *iterator[K, V]) Value() V {
	return it.m.Get(it.Key())
}

func (it *iterator[K, V]) ValueRef() *V {
	return it.m.GetRef(it.Key())
}

func (it *iterator[K, V]) SetValue(value V) {
	it.m.Set(it.Key(), value)
}

func (it *iterator[K, V]) Remove() {
	it.m.Remove(it.Key())
}
