package sequence

import "github.com/djordje200179/extendedlibrary/datastructures/sequences"

type iterator[T any] struct {
	sequences.Iterator[T]
	seq *SynchronizedSequence[T]
}

func (it iterator[T]) GetRef() *T {
	it.seq.mutex.Lock()
	defer it.seq.mutex.Unlock()

	return it.Iterator.GetRef()
}

func (it iterator[T]) Get() T      { return *it.GetRef() }
func (it iterator[T]) Set(value T) { *it.GetRef() = value }

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
