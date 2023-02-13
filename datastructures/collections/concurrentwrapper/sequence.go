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

func (seq *Wrapper[T]) Size() int {
	seq.mutex.RLock()
	defer seq.mutex.RUnlock()

	return seq.Collection.Size()
}

func (seq *Wrapper[T]) Get(index int) T {
	seq.mutex.RLock()
	defer seq.mutex.RUnlock()

	return seq.Collection.Get(index)
}

func (seq *Wrapper[T]) GetRef(index int) *T {
	seq.mutex.RLock()
	defer seq.mutex.RUnlock()

	return seq.Collection.GetRef(index)
}

func (seq *Wrapper[T]) Set(index int, value T) {
	seq.mutex.RLock()
	defer seq.mutex.RUnlock()

	seq.Collection.Set(index, value)
}

func (seq *Wrapper[T]) Append(value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Append(value)
}

func (seq *Wrapper[T]) AppendMany(values ...T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Append(values...)
}

func (seq *Wrapper[T]) Insert(index int, value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Insert(index, value)
}

func (seq *Wrapper[T]) Remove(index int) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Remove(index)
}

func (seq *Wrapper[T]) Clear() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Clear()
}

func (seq *Wrapper[T]) Reverse() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Reverse()
}

func (seq *Wrapper[T]) Swap(index1, index2 int) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Swap(index1, index2)
}

func (seq *Wrapper[T]) Sort(comparator functions.Comparator[T]) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Sort(comparator)
}

func (seq *Wrapper[T]) Join(other collections.Collection[T]) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Collection.Join(other)
}

func (seq *Wrapper[T]) Clone() collections.Collection[T] {
	seq.mutex.RLock()
	defer seq.mutex.RUnlock()

	return &Wrapper[T]{
		Collection: seq.Collection.Clone(),
	}
}

func (seq *Wrapper[T]) Iterator() iterable.Iterator[T] {
	return seq.ModifyingIterator()
}

func (seq *Wrapper[T]) ModifyingIterator() collections.Iterator[T] {
	return iterator[T]{
		Iterator: seq.Collection.ModifyingIterator(),
		mutex:    &seq.mutex,
	}
}

func (seq *Wrapper[T]) Stream() streams.Stream[T] {
	return seq.Collection.Stream()
}

func (seq *Wrapper[T]) RefStream() streams.Stream[*T] {
	return seq.Collection.RefStream()
}
