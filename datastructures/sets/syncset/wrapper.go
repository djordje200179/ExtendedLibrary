package syncset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Wrapper[T any] struct {
	set sets.Set[T]

	mutex sync.RWMutex
}

func From[T any](set sets.Set[T]) *Wrapper[T] {
	return &Wrapper[T]{set, sync.RWMutex{}}
}

func (wrapper *Wrapper[T]) Size() int {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.set.Size()
}

func (wrapper *Wrapper[T]) Add(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.set.Add(value)
}

func (wrapper *Wrapper[T]) Remove(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.set.Remove(value)
}

func (wrapper *Wrapper[T]) Contains(value T) bool {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.set.Contains(value)
}

func (wrapper *Wrapper[T]) Clear() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.set.Clear()
}

func (wrapper *Wrapper[T]) Clone() sets.Set[T] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	clonedSet := wrapper.set.Clone()
	return From[T](clonedSet)
}

func (wrapper *Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.SetIterator()
}

func (wrapper *Wrapper[T]) SetIterator() sets.Iterator[T] {
	return Iterator[T]{wrapper.set.SetIterator(), &wrapper.mutex}
}

func (wrapper *Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.set.Stream()
}

func (wrapper *Wrapper[T]) Transaction(updateFunction func(set sets.Set[T])) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	updateFunction(wrapper.set)
}
