package linkedlistmap

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/collections/linkedlist"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type LinkedListMap[K comparable, V any] linkedlist.List[misc.Pair[K, V]]

func New[K comparable, V any]() *LinkedListMap[K, V] {
	list := linkedlist.New[misc.Pair[K, V]]()
	return (*LinkedListMap[K, V])(list)
}

func Collector[K comparable, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V]{New[K, V]()}
}

func (m *LinkedListMap[K, V]) Size() int {
	return m.List().Size()
}

func (m *LinkedListMap[K, V]) find(key K) collections.Iterator[misc.Pair[K, V]] {
	for it := m.List().ModifyingIterator(); it.Valid(); it.Move() {
		if it.Get().First == key {
			return it
		}
	}

	return nil
}

func (m *LinkedListMap[K, V]) GetRef(key K) *V {
	it := m.find(key)
	if it == nil {
		panic(fmt.Sprintf("Key %v not found", key))
	}

	return &it.GetRef().Second
}

func (m *LinkedListMap[K, V]) Get(key K) V {
	return *m.GetRef(key)
}

func (m *LinkedListMap[K, V]) Set(key K, value V) {
	it := m.find(key)
	if it == nil {
		m.List().Append(misc.Pair[K, V]{key, value})
	} else {
		it.GetRef().Second = value
	}
}

func (m *LinkedListMap[K, V]) Keys() []K {
	keys := make([]K, m.Size())

	i := 0
	for it := m.List().Iterator(); it.Valid(); it.Move() {
		keys[i] = it.Get().First
		i++
	}

	return keys
}

func (m *LinkedListMap[K, V]) Remove(key K) {
	it := m.find(key)

	if it != nil {
		it.Remove()
	}
}

func (m *LinkedListMap[K, V]) Contains(key K) bool {
	return m.find(key) != nil
}

func (m *LinkedListMap[K, V]) Clear() {
	m.List().Clear()
}

func (m *LinkedListMap[K, V]) Swap(key1, key2 K) {
	it1 := m.find(key1)
	if it1 == nil {
		panic(fmt.Sprintf("Key %v not found", key1))
	}

	it2 := m.find(key2)
	if it2 == nil {
		panic(fmt.Sprintf("Key %v not found", key2))
	}

	it1.GetRef().First, it2.GetRef().First = it2.GetRef().First, it1.GetRef().First
}

func (m *LinkedListMap[K, V]) Clone() maps.Map[K, V] {
	clonedList := m.List().Clone().(*linkedlist.List[misc.Pair[K, V]])
	return (*LinkedListMap[K, V])(clonedList)
}

func (m *LinkedListMap[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return m.ModifyingIterator()
}

func (m *LinkedListMap[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return iterator[K, V]{m.List().ModifyingIterator()}
}

func (m *LinkedListMap[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	supplier := iterable.IteratorSupplier[misc.Pair[K, V]]{m.Iterator()}
	return streams.Stream[misc.Pair[K, V]]{supplier}
}

func (m *LinkedListMap[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] {
	supplier := maps.RefsSupplier[K, V]{m.ModifyingIterator()}
	return streams.Stream[misc.Pair[K, *V]]{supplier}
}

func (m *LinkedListMap[K, V]) List() *linkedlist.List[misc.Pair[K, V]] {
	return (*linkedlist.List[misc.Pair[K, V]])(m)
}
