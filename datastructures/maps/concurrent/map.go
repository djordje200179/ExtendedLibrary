package concurrent

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Map[K comparable, V any] struct {
	maps.Map[K, V]
	mutex sync.Mutex
}

func (m *Map[K, V]) Size() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Size()
}

func (m *Map[K, V]) Get(key K) V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Get(key)
}

func (m *Map[K, V]) GetRef(key K) *V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.GetRef(key)
}

func (m *Map[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Set(key, value)
}

func (m *Map[K, V]) Remove(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Remove(key)
}

func (m *Map[K, V]) Contains(key K) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Contains(key)
}

func (m *Map[K, V]) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Clear()
}

func (m *Map[K, V]) Clone() maps.Map[K, V] {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return &Map[K, V]{Map: m.Map.Clone()}
}

func (m *Map[K, V]) Iterator() collections.Iterator[maps.Entry[K, V]] {
	return m.ModifyingIterator()
}

func (m *Map[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return iterator[K, V]{m.Map.ModifyingIterator(), m}
}

func (m *Map[K, V]) Stream() streams.Stream[misc.Pair[K, V]]     { return m.Map.Stream() }
func (m *Map[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] { return m.Map.RefStream() }
