package concurrentcollection

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Wrapper[T any] struct {
	collection collections.Collection[T]

	mutex sync.RWMutex
}

func From[T any](collection collections.Collection[T]) *Wrapper[T] {
	return &Wrapper[T]{collection, sync.RWMutex{}}
}

func (wrapper *Wrapper[T]) Size() int {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.Size()
}

func (wrapper *Wrapper[T]) Get(index int) T {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.Get(index)
}

func (wrapper *Wrapper[T]) GetRef(index int) *T {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.GetRef(index)
}

func (wrapper *Wrapper[T]) Set(index int, value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Set(index, value)
}

func (wrapper *Wrapper[T]) Update(index int, updateFunction func(value T) T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.collection.Get(index)
	newValue := updateFunction(oldValue)
	wrapper.collection.Set(index, newValue)
}

func (wrapper *Wrapper[T]) UpdateRef(index int, updateFunction func(value *T)) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	oldValue := wrapper.collection.GetRef(index)
	updateFunction(oldValue)
}

func (wrapper *Wrapper[T]) Append(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Append(value)
}

func (wrapper *Wrapper[T]) Insert(index int, value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Insert(index, value)
}

func (wrapper *Wrapper[T]) Remove(index int) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Remove(index)
}

func (wrapper *Wrapper[T]) Clear() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Clear()
}

func (wrapper *Wrapper[T]) Reverse() {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Reverse()
}

func (wrapper *Wrapper[T]) Swap(index1, index2 int) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Swap(index1, index2)
}

func (wrapper *Wrapper[T]) Sort(comparator functions.Comparator[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Sort(comparator)
}

func (wrapper *Wrapper[T]) Join(other collections.Collection[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Join(other)
}

func (wrapper *Wrapper[T]) Clone() collections.Collection[T] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	clonedCollection := wrapper.collection.Clone()
	return &Wrapper[T]{clonedCollection, sync.RWMutex{}}
}

func (wrapper *Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.ModifyingIterator()
}

func (wrapper *Wrapper[T]) ModifyingIterator() collections.Iterator[T] {
	return iterator[T]{wrapper.collection.ModifyingIterator(), &wrapper.mutex}
}

func (wrapper *Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.collection.Stream()
}

func (wrapper *Wrapper[T]) RefStream() streams.Stream[*T] {
	return wrapper.collection.RefStream()
}

func (wrapper *Wrapper[T]) Mutex() *sync.RWMutex {
	return &wrapper.mutex
}
