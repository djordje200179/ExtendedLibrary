package sequence

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type SynchronizedSequence[T any] struct {
	sequence sequences.Sequence[T]
	mutex    sync.Mutex
}

func New[T any](sequence sequences.Sequence[T]) *SynchronizedSequence[T] {
	syncSeq := new(SynchronizedSequence[T])

	syncSeq.sequence = sequence
	syncSeq.mutex = sync.Mutex{}

	return syncSeq
}

func (syncSeq *SynchronizedSequence[T]) Size() int {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	return syncSeq.sequence.Size()
}

func (syncSeq *SynchronizedSequence[T]) Get(index int) T {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	return syncSeq.sequence.Get(index)
}

func (syncSeq *SynchronizedSequence[T]) Set(index int, value T) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Set(index, value)
}

func (syncSeq *SynchronizedSequence[T]) Append(value T) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Append(value)
}

func (syncSeq *SynchronizedSequence[T]) AppendMany(values ...T) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.AppendMany(values...)
}

func (syncSeq *SynchronizedSequence[T]) Insert(index int, value T) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Insert(index, value)
}

func (syncSeq *SynchronizedSequence[T]) Remove(index int) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Remove(index)
}

func (syncSeq *SynchronizedSequence[T]) Clear() {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Clear()
}

func (syncSeq *SynchronizedSequence[T]) Reverse() {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Reverse()
}

func (syncSeq *SynchronizedSequence[T]) Sort(comparator functions.Comparator[T]) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Sort(comparator)
}

func (syncSeq *SynchronizedSequence[T]) Join(other sequences.Sequence[T]) {
	syncSeq.mutex.Lock()
	defer syncSeq.mutex.Unlock()

	syncSeq.sequence.Join(other)
}

func (syncSeq *SynchronizedSequence[T]) Clone() sequences.Sequence[T] {
	return New[T](syncSeq.sequence.Clone())
}

func (syncSeq *SynchronizedSequence[T]) Iterator() datastructures.Iterator[T] {
	return syncSeq.ModifyingIterator()
}

func (syncSeq *SynchronizedSequence[T]) ModifyingIterator() sequences.Iterator[T] {
	return iterator[T]{
		Iterator: syncSeq.sequence.ModifyingIterator(),
		seq:      syncSeq,
	}
}

func (syncSeq *SynchronizedSequence[T]) Stream() *streams.Stream[T] {
	return syncSeq.sequence.Stream()
}
