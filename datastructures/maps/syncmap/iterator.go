package syncmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"sync"
)

type Iterator[K, V any] struct {
	maps.Iterator[K, V]

	mutex *sync.RWMutex
}

func (it Iterator[K, V]) Get() misc.Pair[K, V] {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return misc.Pair[K, V]{
		First:  it.Iterator.Key(),
		Second: it.Iterator.Value(),
	}
}

func (it Iterator[K, V]) Key() K {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.Key()
}

func (it Iterator[K, V]) Value() V {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.Value()
}

func (it Iterator[K, V]) ValueRef() *V {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.ValueRef()
}

func (it Iterator[K, V]) SetValue(value V) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.SetValue(value)
}

func (it Iterator[K, V]) Remove() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.Remove()
}
