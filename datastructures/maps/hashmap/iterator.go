package hashmap

import "github.com/djordje200179/extendedlibrary/datastructures/maps"

type iterator[K comparable, V any] struct {
	m     Map[K, V]
	keys  []K
	index int
}

func (it *iterator[K, V]) IsValid() bool {
	return it.index < len(it.keys)
}

func (it *iterator[K, V]) Move() {
	it.index++
}

func (it *iterator[K, V]) Get() maps.Entry[K, V] {
	return entry[K, V]{
		m:   it.m,
		key: it.keys[it.index],
	}
}
