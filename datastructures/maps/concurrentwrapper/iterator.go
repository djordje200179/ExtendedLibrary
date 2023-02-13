package concurrentwrapper

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"sync"
)

type iterator[K comparable, V any] struct {
	maps.Iterator[K, V]
	mutex *sync.RWMutex
}

func (it iterator[K, V]) Get() K {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.Get()
}

func (it iterator[K, V]) Value() V {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.Value()
}

func (it iterator[K, V]) ValueRef() *V {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.ValueRef()
}

func (it iterator[K, V]) SetValue(value V) {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	it.Iterator.SetValue(value)
}

func (it iterator[K, V]) Remove() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.Remove()
}
