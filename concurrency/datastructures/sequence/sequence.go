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

func (seq *SynchronizedSequence[T]) Size() int {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.sequence.Size()
}

func (seq *SynchronizedSequence[T]) Get(index int) T {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.sequence.Get(index)
}

func (seq *SynchronizedSequence[T]) GetRef(index int) *T {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.sequence.GetRef(index)
}

func (seq *SynchronizedSequence[T]) Set(index int, value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.Set(index, value)
}

func (seq *SynchronizedSequence[T]) Append(value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.Append(value)
}

func (seq *SynchronizedSequence[T]) AppendMany(values ...T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.AppendMany(values...)
}

func (seq *SynchronizedSequence[T]) Insert(index int, value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.Insert(index, value)
}

func (seq *SynchronizedSequence[T]) Remove(index int) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.Remove(index)
}

func (seq *SynchronizedSequence[T]) Clear() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.Clear()
}

func (seq *SynchronizedSequence[T]) Reverse() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.Reverse()
}

func (seq *SynchronizedSequence[T]) Sort(comparator functions.Comparator[T]) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.Sort(comparator)
}

func (seq *SynchronizedSequence[T]) Join(other sequences.Sequence[T]) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.sequence.Join(other)
}

func (seq *SynchronizedSequence[T]) Clone() sequences.Sequence[T] {
	return New[T](seq.sequence.Clone())
}

func (seq *SynchronizedSequence[T]) Iterator() datastructures.Iterator[T] {
	return seq.ModifyingIterator()
}

func (seq *SynchronizedSequence[T]) ModifyingIterator() sequences.Iterator[T] {
	return iterator[T]{seq.sequence.ModifyingIterator(), seq}
}

func (seq *SynchronizedSequence[T]) Stream() *streams.Stream[T] {
	return seq.sequence.Stream()
}

func (seq *SynchronizedSequence[T]) RefStream() *streams.Stream[*T] {
	return seq.sequence.RefStream()
}
