package concurrentwrapper

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Wrapper[K comparable, V any] struct {
	maps.Map[K, V]
	mutex sync.Mutex
}

func (m *Wrapper[K, V]) Size() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Size()
}

func (m *Wrapper[K, V]) Get(key K) V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Get(key)
}

func (m *Wrapper[K, V]) GetRef(key K) *V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.GetRef(key)
}

func (m *Wrapper[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Set(key, value)
}

func (m *Wrapper[K, V]) Remove(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Remove(key)
}

func (m *Wrapper[K, V]) Contains(key K) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Contains(key)
}

func (m *Wrapper[K, V]) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Clear()
}

func (m *Wrapper[K, V]) Clone() maps.Map[K, V] {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return &Wrapper[K, V]{Map: m.Map.Clone()}
}

func (m *Wrapper[K, V]) Iterator() iterable.Iterator[maps.Entry[K, V]] {
	return m.ModifyingIterator()
}

func (m *Wrapper[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return iterator[K, V]{m.Map.ModifyingIterator(), m}
}

func (m *Wrapper[K, V]) Stream() streams.Stream[misc.Pair[K, V]]     { return m.Map.Stream() }
func (m *Wrapper[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] { return m.Map.RefStream() }
