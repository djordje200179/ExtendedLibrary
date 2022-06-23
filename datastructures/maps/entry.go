package maps

import "github.com/djordje200179/extendedlibrary/misc"

type Entry[K comparable, V any] struct {
	m   Map[K, V]
	key K
}

func NewEntry[K comparable, V any](m Map[K, V], key K) Entry[K, V] {
	return Entry[K, V]{
		m:   m,
		key: key,
	}
}

func (e Entry[K, V]) Get() misc.Pair[K, V] {
	return misc.Pair[K, V]{e.key, e.Value()}
}

func (e Entry[K, V]) Key() K {
	return e.key
}

func (e Entry[K, V]) Value() V {
	return e.m.Get(e.key)
}

func (e Entry[K, V]) SetValue(value V) {
	e.m.Set(e.key, value)
}
