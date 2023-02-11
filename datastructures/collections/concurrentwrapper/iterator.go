package concurrentwrapper

import "github.com/djordje200179/extendedlibrary/datastructures/collections"

type iterator[T any] struct {
	collections.Iterator[T]
	seq *Wrapper[T]
}

func (it iterator[T]) GetRef() *T {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	return it.Iterator.GetRef()
}

func (it iterator[T]) Get() T {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	return it.Iterator.Get()
}

func (it iterator[T]) Set(value T) {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	it.Iterator.Set(value)
}

func (it iterator[T]) InsertBefore(value T) {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	it.Iterator.InsertBefore(value)
}

func (it iterator[T]) InsertAfter(value T) {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	it.Iterator.InsertAfter(value)
}

func (it iterator[T]) Remove() {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	it.Iterator.Remove()
}
