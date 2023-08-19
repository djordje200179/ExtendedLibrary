package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/collections/linkedlist"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Map[K comparable, V any] linkedlist.List[misc.Pair[K, V]]

func New[K comparable, V any]() *Map[K, V] {
	list := linkedlist.New[misc.Pair[K, V]]()
	return FromList(list)
}

func FromList[K comparable, V any](list *linkedlist.List[misc.Pair[K, V]]) *Map[K, V] {
	return (*Map[K, V])(list)
}

func Collector[K comparable, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V]{New[K, V]()}
}

func (m *Map[K, V]) Size() int {
	return m.List().Size()
}

func (m *Map[K, V]) find(key K) collections.Iterator[misc.Pair[K, V]] {
	for it := m.List().ModifyingIterator(); it.Valid(); it.Move() {
		if it.Get().First == key {
			return it
		}
	}

	return nil
}

func (m *Map[K, V]) Get(key K) V {
	return *m.GetRef(key)
}

func (m *Map[K, V]) GetOrDefault(key K) V {
	return m.GetOrElse(key, functions.Zero[V]())
}

func (m *Map[K, V]) GetOrElse(key K, value V) V {
	it := m.find(key)
	if it == nil {
		return value
	}

	return it.Get().Second
}

func (m *Map[K, V]) TryGet(key K) (V, bool) {
	it := m.find(key)
	if it == nil {
		var zero V
		return zero, false
	}

	return it.Get().Second, true
}

func (m *Map[K, V]) GetRef(key K) *V {
	it := m.find(key)
	if it == nil {
		maps.PanicOnMissingKey(key)
	}

	return &it.GetRef().Second
}

func (m *Map[K, V]) Set(key K, value V) {
	it := m.find(key)
	if it == nil {
		m.List().Append(misc.Pair[K, V]{key, value})
	} else {
		it.GetRef().Second = value
	}
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())

	i := 0
	for it := m.List().Iterator(); it.Valid(); it.Move() {
		keys[i] = it.Get().First
		i++
	}

	return keys
}

func (m *Map[K, V]) Remove(key K) {
	it := m.find(key)

	if it != nil {
		it.Remove()
	}
}

func (m *Map[K, V]) Contains(key K) bool {
	return m.find(key) != nil
}

func (m *Map[K, V]) Clear() {
	m.List().Clear()
}

func (m *Map[K, V]) Clone() maps.Map[K, V] {
	clonedList := m.List().Clone().(*linkedlist.List[misc.Pair[K, V]])
	return (*Map[K, V])(clonedList)
}

func (m *Map[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return m.ModifyingIterator()
}

func (m *Map[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return Iterator[K, V]{m.List().ModifyingIterator()}
}

func (m *Map[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return iterable.IteratorStream(m.Iterator())
}

func (m *Map[K, V]) RefsStream() streams.Stream[misc.Pair[K, *V]] {
	return maps.RefsStream[K, V](m)
}

func (m *Map[K, V]) List() *linkedlist.List[misc.Pair[K, V]] {
	return (*linkedlist.List[misc.Pair[K, V]])(m)
}
