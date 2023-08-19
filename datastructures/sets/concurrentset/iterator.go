package concurrentset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"sync"
)

type Iterator[T any] struct {
	sets.Iterator[T]

	mutex *sync.RWMutex
}

func (it Iterator[T]) Get() T {
	it.mutex.RLock()
	defer it.mutex.RUnlock()

	return it.Iterator.Get()
}

func (it Iterator[T]) Remove() {
	it.mutex.Lock()
	defer it.mutex.Unlock()

	it.Iterator.Remove()
}
