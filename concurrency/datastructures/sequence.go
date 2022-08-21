package datastructures

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type SynchronizedSequence[T any] struct {
	sequences.Sequence[T]
	mutex sync.Mutex
}

func FromSequence[T any](sequence sequences.Sequence[T]) *SynchronizedSequence[T] {
	syncSeq := new(SynchronizedSequence[T])

	syncSeq.Sequence = sequence
	syncSeq.mutex = sync.Mutex{}

	return syncSeq
}

func (seq *SynchronizedSequence[T]) Size() int {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.Sequence.Size()
}

func (seq *SynchronizedSequence[T]) Get(index int) T {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.Sequence.Get(index)
}

func (seq *SynchronizedSequence[T]) GetRef(index int) *T {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.Sequence.GetRef(index)
}

func (seq *SynchronizedSequence[T]) Set(index int, value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Set(index, value)
}

func (seq *SynchronizedSequence[T]) Append(value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Append(value)
}

func (seq *SynchronizedSequence[T]) AppendMany(values ...T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.AppendMany(values...)
}

func (seq *SynchronizedSequence[T]) Insert(index int, value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Insert(index, value)
}

func (seq *SynchronizedSequence[T]) Remove(index int) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Remove(index)
}

func (seq *SynchronizedSequence[T]) Clear() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Clear()
}

func (seq *SynchronizedSequence[T]) Reverse() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Reverse()
}

func (seq *SynchronizedSequence[T]) Sort(comparator functions.Comparator[T]) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Sort(comparator)
}

func (seq *SynchronizedSequence[T]) Join(other sequences.Sequence[T]) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Join(other)
}

func (seq *SynchronizedSequence[T]) Clone() sequences.Sequence[T] {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return FromSequence[T](seq.Sequence.Clone())
}

func (seq *SynchronizedSequence[T]) Iterator() collections.Iterator[T] {
	return seq.ModifyingIterator()
}

func (seq *SynchronizedSequence[T]) ModifyingIterator() sequences.Iterator[T] {
	return sequenceIterator[T]{seq.Sequence.ModifyingIterator(), seq}
}

func (seq *SynchronizedSequence[T]) Stream() streams.Stream[T]     { return seq.Sequence.Stream() }
func (seq *SynchronizedSequence[T]) RefStream() streams.Stream[*T] { return seq.Sequence.RefStream() }
