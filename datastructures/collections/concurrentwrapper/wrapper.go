package concurrentwrapper

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Wrapper[T any] struct {
	collections.Collection[T]

	mutex sync.RWMutex
}

func (wrapper *Wrapper[T]) Size() int {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.Collection.Size()
}

func (wrapper *Wrapper[T]) Get(index int) T {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.Collection.Get(index)
}

func (wrapper *Wrapper[T]) GetRef(index int) *T {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.Collection.GetRef(index)
}

func (wrapper *Wrapper[T]) Set(index int, value T) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	wrapper.Collection.Set(index, value)
}

func (wrapper *Wrapper[T]) Append(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Append(value)
}

func (wrapper *Wrapper[T]) AppendMany(values ...T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Append(values...)
}

func (wrapper *Wrapper[T]) Insert(index int, value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Insert(index, value)
}

func (wrapper *Wrapper[T]) Remove(index int) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Remove(index)
}

func (wrapper *Wrapper[T]) Clear() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Clear()
}

func (wrapper *Wrapper[T]) Reverse() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Reverse()
}

func (wrapper *Wrapper[T]) Swap(index1, index2 int) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Swap(index1, index2)
}

func (wrapper *Wrapper[T]) Sort(comparator functions.Comparator[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Sort(comparator)
}

func (wrapper *Wrapper[T]) Join(other collections.Collection[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.Collection.Join(other)
}

func (wrapper *Wrapper[T]) Clone() collections.Collection[T] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return &Wrapper[T]{
		Collection: wrapper.Collection.Clone(),
	}
}

func (wrapper *Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.ModifyingIterator()
}

func (wrapper *Wrapper[T]) ModifyingIterator() collections.Iterator[T] {
	return iterator[T]{
		Iterator: wrapper.Collection.ModifyingIterator(),
		mutex:    &wrapper.mutex,
	}
}

func (wrapper *Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.Collection.Stream()
}

func (wrapper *Wrapper[T]) RefStream() streams.Stream[*T] {
	return wrapper.Collection.RefStream()
}
