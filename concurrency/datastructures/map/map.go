package _map

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type SynchronizedMap[K comparable, V any] struct {
	m     maps.Map[K, V]
	mutex sync.Mutex
}

func New[K comparable, V any](m maps.Map[K, V]) *SynchronizedMap[K, V] {
	syncMap := new(SynchronizedMap[K, V])

	syncMap.m = m
	syncMap.mutex = sync.Mutex{}

	return syncMap
}

func (m *SynchronizedMap[K, V]) Size() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.m.Size()
}

func (m *SynchronizedMap[K, V]) Get(key K) V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.m.Get(key)
}

func (m *SynchronizedMap[K, V]) GetRef(key K) *V {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.m.GetRef(key)
}

func (m *SynchronizedMap[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.m.Set(key, value)
}

func (m *SynchronizedMap[K, V]) Remove(key K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.m.Remove(key)
}

func (m *SynchronizedMap[K, V]) Contains(key K) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.m.Contains(key)
}

func (m *SynchronizedMap[K, V]) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.m.Clear()
}

func (m *SynchronizedMap[K, V]) Clone() maps.Map[K, V] {
	return New[K, V](m.m.Clone())
}

func (m *SynchronizedMap[K, V]) Iterator() datastructures.Iterator[maps.Entry[K, V]] {
	return m.ModifyingIterator()
}

func (m *SynchronizedMap[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return iterator[K, V]{m.m.ModifyingIterator(), m}
}

func (m *SynchronizedMap[K, V]) Stream() *streams.Stream[misc.Pair[K, V]] {
	return m.m.Stream()
}
