package datastructures

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
)

type mapIterator[K comparable, V any] struct {
	maps.Iterator[K, V]
	m *SynchronizedMap[K, V]
}

func (it mapIterator[K, V]) Get() maps.Entry[K, V] {
	it.m.mutex.Lock()
	defer it.m.mutex.Unlock()

	return it.Iterator.Get()
}

func (it mapIterator[K, V]) Remove() {
	it.m.mutex.Lock()
	defer it.m.mutex.Unlock()

	it.Iterator.Remove()
}
