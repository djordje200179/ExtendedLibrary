package concurrentwrapper

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"sync"
)

type iterator[T any] struct {
	collections.Iterator[T]

	mutex *sync.RWMutex
}

func (it iterator[T]) GetRef() *T {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.GetRef()
}

func (it iterator[T]) Get() T {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.Get()
}

func (it iterator[T]) Set(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.Set(value)
}

func (it iterator[T]) InsertBefore(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.InsertBefore(value)
}

func (it iterator[T]) InsertAfter(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.InsertAfter(value)
}

func (it iterator[T]) Remove() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.Remove()
}
