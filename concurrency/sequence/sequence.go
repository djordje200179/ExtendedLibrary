package sequence

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Sequence[T any] struct {
	sequence sequences.Sequence[T]
	mutex    sync.Mutex
}

func New[T any](sequence sequences.Sequence[T]) *Sequence[T] {
	syncSeq := new(Sequence[T])

	syncSeq.sequence = sequence
	syncSeq.mutex = sync.Mutex{}

	return syncSeq
}

func (syncSeq *Sequence[T]) Size() int {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	return syncSeq.sequence.Size()
}

func (syncSeq *Sequence[T]) Get(index int) T {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	return syncSeq.sequence.Get(index)
}

func (syncSeq *Sequence[T]) Set(index int, value T) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Set(index, value)
}

func (syncSeq *Sequence[T]) Append(value T) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Append(value)
}

func (syncSeq *Sequence[T]) AppendMany(values ...T) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.AppendMany(values...)
}

func (syncSeq *Sequence[T]) Insert(index int, value T) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Insert(index, value)
}

func (syncSeq *Sequence[T]) Remove(index int) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Remove(index)
}

func (syncSeq *Sequence[T]) Clear() {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Clear()
}

func (syncSeq *Sequence[T]) Sort(comparator functions.Comparator[T]) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Sort(comparator)
}

func (syncSeq *Sequence[T]) Join(other sequences.Sequence[T]) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Join(other)
}

func (syncSeq *Sequence[T]) Clone() *Sequence[T] {
	return New[T](syncSeq.sequence.Clone())
}

func (syncSeq *Sequence[T]) Iterator() datastructures.Iterator[T] {
	return syncSeq.sequence.Iterator()
}

func (syncSeq *Sequence[T]) ModifyingIterator() sequences.Iterator[T] {
	return syncSeq.sequence.ModifyingIterator()
}

func (syncSeq *Sequence[T]) Stream() *streams.Stream[T] {
	return syncSeq.sequence.Stream()
}
