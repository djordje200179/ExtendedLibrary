package datastructures

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type SynchronizedMap[K comparable, V any] struct {
	maps.Map[K, V]
	mutex sync.Mutex
}

func FromMap[K comparable, V any](m maps.Map[K, V]) *SynchronizedMap[K, V] {
	syncMap := new(SynchronizedMap[K, V])

	syncMap.Map = m
	syncMap.mutex = sync.Mutex{}

	return syncMap
}

func (m *SynchronizedMap[K, V]) Size() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Size()
}

func (m *SynchronizedMap[K, V]) Get(key K) V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Get(key)
}

func (m *SynchronizedMap[K, V]) GetRef(key K) *V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.GetRef(key)
}

func (m *SynchronizedMap[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Set(key, value)
}

func (m *SynchronizedMap[K, V]) Remove(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Remove(key)
}

func (m *SynchronizedMap[K, V]) Contains(key K) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Map.Contains(key)
}

func (m *SynchronizedMap[K, V]) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Map.Clear()
}

func (m *SynchronizedMap[K, V]) Clone() maps.Map[K, V] {
	return FromMap[K, V](m.Map.Clone())
}

func (m *SynchronizedMap[K, V]) Iterator() datastructures.Iterator[maps.Entry[K, V]] {
	return m.ModifyingIterator()
}

func (m *SynchronizedMap[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return mapIterator[K, V]{m.Map.ModifyingIterator(), m}
}

func (m *SynchronizedMap[K, V]) Stream() *streams.Stream[misc.Pair[K, V]] {
	return m.Map.Stream()
}
