package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Map[K comparable, V any] struct {
	list *linkedlist.LinkedList[misc.Pair[K, V]]
}

func New[K comparable, V any]() Map[K, V] {
	return Map[K, V]{linkedlist.New[misc.Pair[K, V]]()}
}

func (m Map[K, V]) find(key K) sequences.Iterator[misc.Pair[K, V]] {
	for it := m.list.Iterator(); it.IsValid(); it.Move() {
		if it.Get().First == key {
			return it
		}
	}

	return nil
}

func (m Map[K, V]) Size() int {
	return m.list.Size()
}

func (m Map[K, V]) Get(key K) V {
	return m.find(key).Get().Second
}

func (m Map[K, V]) Set(key K, value V) {
	it := m.find(key)

	if it != nil {
		data := it.Get()
		data.Second = value
		it.Set(data)
	} else {
		m.list.Append(misc.Pair[K, V]{key, value})
	}
}

func (m Map[K, V]) Remove(key K) {
	it := m.find(key)
	if it != nil {
		it.Remove()
	}
}

func (m Map[K, V]) Contains(key K) bool {
	return m.find(key) != nil
}

func (m Map[K, V]) Empty() {
	m.list.Empty()
}

func (m Map[K, V]) Clone() Map[K, V] {

}

func (m Map[K, V]) Iterator() maps.Iterator[K, V] {
	return iterator[K, V]{
		m:        m,
		Iterator: m.list.Iterator(),
	}
}

func (m Map[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return m.list.Stream()
}
