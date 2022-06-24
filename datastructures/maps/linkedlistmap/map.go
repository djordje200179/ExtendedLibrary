package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Map[K comparable, V any] linkedlist.LinkedList[misc.Pair[K, V]]

func New[K comparable, V any]() *Map[K, V] {
	list := linkedlist.New[misc.Pair[K, V]]()
	return (*Map[K, V])(list)
}

func Collector[K comparable, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V](New[K, V]())
}

func (m *Map[K, V]) list() *linkedlist.LinkedList[misc.Pair[K, V]] {
	return (*linkedlist.LinkedList[misc.Pair[K, V]])(m)
}

func (m *Map[K, V]) find(key K) sequences.Iterator[misc.Pair[K, V]] {
	for it := m.list().ModifyingIterator(); it.Valid(); it.Move() {
		if it.Get().First == key {
			return it
		}
	}

	return nil
}

func (m *Map[K, V]) Size() int { return m.list().Size() }

func (m *Map[K, V]) GetRef(key K) *V {
	if it := m.find(key); it != nil {
		return &it.GetRef().Second
	} else {
		return nil
	}
}

func (m *Map[K, V]) Get(key K) V {
	if ptr := m.GetRef(key); ptr != nil {
		return *ptr
	} else {
		var empty V
		return empty
	}
}

func (m *Map[K, V]) Set(key K, value V) {
	if ptr := m.GetRef(key); ptr != nil {
		*ptr = value
	} else {
		m.list().Append(misc.Pair[K, V]{key, value})
	}
}

func (m *Map[K, V]) Remove(key K) {
	if it := m.find(key); it != nil {
		it.Remove()
	}
}

func (m *Map[K, V]) Contains(key K) bool { return m.find(key) != nil }

func (m *Map[K, V]) Clear() { m.list().Clear() }

func (m *Map[K, V]) Clone() maps.Map[K, V] {
	clonedList := m.list().Clone().(*linkedlist.LinkedList[misc.Pair[K, V]])
	return (*Map[K, V])(clonedList)
}

func (m *Map[K, V]) Iterator() datastructures.Iterator[maps.Entry[K, V]] {
	return m.ModifyingIterator()
}

func (m *Map[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return iterator[K, V]{
		m:        m,
		Iterator: m.list().ModifyingIterator(),
	}
}

func (m *Map[K, V]) Stream() *streams.Stream[misc.Pair[K, V]] { return m.list().Stream() }
