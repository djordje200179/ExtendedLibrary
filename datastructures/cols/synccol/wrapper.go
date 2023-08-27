package synccol

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Wrapper[T any] struct {
	collection cols.Collection[T]

	mutex sync.RWMutex
}

func From[T any](collection cols.Collection[T]) *Wrapper[T] {
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

func (wrapper *Wrapper[T]) Prepend(value T) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Prepend(value)
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

func (wrapper *Wrapper[T]) Sort(comparator comparison.Comparator[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Sort(comparator)
}

func (wrapper *Wrapper[T]) Join(other cols.Collection[T]) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	wrapper.collection.Join(other)
}

func (wrapper *Wrapper[T]) Clone() cols.Collection[T] {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	clonedCollection := wrapper.collection.Clone()
	return &Wrapper[T]{clonedCollection, sync.RWMutex{}}
}

func (wrapper *Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.CollectionIterator()
}

func (wrapper *Wrapper[T]) CollectionIterator() cols.Iterator[T] {
	return Iterator[T]{wrapper.collection.CollectionIterator(), &wrapper.mutex}
}

func (wrapper *Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.collection.Stream()
}

func (wrapper *Wrapper[T]) RefsStream() streams.Stream[*T] {
	return wrapper.collection.RefsStream()
}

func (wrapper *Wrapper[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.FindIndex(predicate)
}

func (wrapper *Wrapper[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	wrapper.mutex.RLock()
	defer wrapper.mutex.RUnlock()

	return wrapper.collection.FindRef(predicate)
}

func (wrapper *Wrapper[T]) Transaction(updateFunction func(collection cols.Collection[T])) {
	wrapper.mutex.Lock()
	defer wrapper.mutex.Unlock()

	updateFunction(wrapper.collection)
}
