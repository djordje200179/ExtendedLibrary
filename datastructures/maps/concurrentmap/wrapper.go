package concurrentmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Wrapper[K, V any] struct {
	m maps.Map[K, V]

	mutex sync.RWMutex
}

func From[K, V any](m maps.Map[K, V]) maps.Map[K, V] {
	return &Wrapper[K, V]{m, sync.RWMutex{}}
}

func (wrapper *Wrapper[K, V]) Size() int {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.Size()
}

func (wrapper *Wrapper[K, V]) Get(key K) V {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.Get(key)
}

func (wrapper *Wrapper[K, V]) GetOrDefault(key K) V {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.GetOrDefault(key)
}

func (wrapper *Wrapper[K, V]) GetOrElse(key K, value V) V {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.GetOrElse(key, value)
}

func (wrapper *Wrapper[K, V]) TryGet(key K) (V, bool) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.TryGet(key)
}

func (wrapper *Wrapper[K, V]) GetRef(key K) *V {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.GetRef(key)
}

func (wrapper *Wrapper[K, V]) Set(key K, value V) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.m.Set(key, value)
}

func (wrapper *Wrapper[K, V]) Update(key K, updateFunction func(oldValue V) V) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.m.Get(key)
	newValue := updateFunction(oldValue)

	wrapper.m.Set(key, newValue)
}

func (wrapper *Wrapper[K, V]) UpdateRef(key K, updateFunction func(oldValue *V)) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.m.GetRef(key)
	updateFunction(oldValue)
}

func (wrapper *Wrapper[K, V]) Keys() []K {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.Keys()
}

func (wrapper *Wrapper[K, V]) Remove(key K) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.m.Remove(key)
}

func (wrapper *Wrapper[K, V]) Contains(key K) bool {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.m.Contains(key)
}

func (wrapper *Wrapper[K, V]) Clear() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.m.Clear()
}

func (wrapper *Wrapper[K, V]) Clone() maps.Map[K, V] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	clonedMap := wrapper.m.Clone()
	return &Wrapper[K, V]{clonedMap, sync.RWMutex{}}
}

func (wrapper *Wrapper[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return wrapper.ModifyingIterator()
}

func (wrapper *Wrapper[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return Iterator[K, V]{wrapper.m.ModifyingIterator(), &wrapper.mutex}
}

func (wrapper *Wrapper[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return wrapper.m.Stream()
}

func (wrapper *Wrapper[K, V]) RefsStream() streams.Stream[misc.Pair[K, *V]] {
	return wrapper.m.RefsStream()
}

func (wrapper *Wrapper[K, V]) Transaction(updateFunction func(m maps.Map[K, V])) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	updateFunction(wrapper.m)
}
