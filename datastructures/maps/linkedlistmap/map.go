package linkedlistmap

import (
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

func (m *Map[K, V]) GetNode(key K) *linkedlist.Node[misc.Pair[K, V]] {
	collectionIterator := m.List().ModifyingIterator()
	listIterator := collectionIterator.(*linkedlist.Iterator[misc.Pair[K, V]])

	for it := listIterator; it.Valid(); it.Move() {
		if it.Get().First == key {
			return it.Node()
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
	node := m.GetNode(key)
	if node == nil {
		return value
	}

	return node.Value.Second
}

func (m *Map[K, V]) TryGet(key K) (V, bool) {
	node := m.GetNode(key)
	if node == nil {
		var zero V
		return zero, false
	}

	return node.Value.Second, true
}

func (m *Map[K, V]) GetRef(key K) *V {
	node := m.GetNode(key)
	if node == nil {
		maps.PanicOnMissingKey(key)
	}

	return &node.Value.Second
}

func (m *Map[K, V]) Set(key K, value V) {
	node := m.GetNode(key)
	if node == nil {
		m.List().Append(misc.Pair[K, V]{key, value})
	} else {
		node.Value.Second = value
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
	node := m.GetNode(key)
	if node == nil {
		return
	}

	m.List().RemoveNode(node)
}

func (m *Map[K, V]) Contains(key K) bool {
	return m.GetNode(key) != nil
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
