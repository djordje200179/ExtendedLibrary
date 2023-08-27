package synccol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"sync"
)

type Iterator[T any] struct {
	cols.Iterator[T]

	mutex *sync.RWMutex
}

func (it Iterator[T]) GetRef() *T {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.GetRef()
}

func (it Iterator[T]) Get() T {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.Get()
}

func (it Iterator[T]) Set(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.Set(value)
}

func (it Iterator[T]) InsertBefore(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.InsertBefore(value)
}

func (it Iterator[T]) InsertAfter(value T) {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.InsertAfter(value)
}

func (it Iterator[T]) Remove() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.Remove()
}
