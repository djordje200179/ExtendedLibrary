package datastructures

import "github.com/djordje200179/extendedlibrary/datastructures/sequences"

type sequenceIterator[T any] struct {
	sequences.Iterator[T]
	seq *SynchronizedSequence[T]
}

func (it sequenceIterator[T]) GetRef() *T {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	return it.Iterator.GetRef()
}

func (it sequenceIterator[T]) Get() T {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	return it.Iterator.Get()
}

func (it sequenceIterator[T]) Set(value T) {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	it.Iterator.Set(value)
}

func (it sequenceIterator[T]) InsertBefore(value T) {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	it.Iterator.InsertBefore(value)
}

func (it sequenceIterator[T]) InsertAfter(value T) {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	it.Iterator.InsertAfter(value)
}

func (it sequenceIterator[T]) Remove() {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	it.Iterator.Remove()
}
