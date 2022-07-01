package _map

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
)

type iterator[K comparable, V any] struct {
	maps.Iterator[K, V]
	m *SynchronizedMap[K, V]
}

func (it iterator[K, V]) Get() maps.Entry[K, V] {
	it.m.mutex.Lock()
	defer it.m.mutex.Unlock()

	return it.Iterator.Get()
}

func (it iterator[K, V]) Remove() {
	it.m.mutex.Lock()
	defer it.m.mutex.Unlock()

	it.Iterator.Remove()
}
