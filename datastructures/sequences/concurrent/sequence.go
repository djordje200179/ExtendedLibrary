package concurrent

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync"
)

type Sequence[T any] struct {
	sequences.Sequence[T]
	mutex sync.Mutex
}

func (seq *Sequence[T]) Size() int {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.Sequence.Size()
}

func (seq *Sequence[T]) Get(index int) T {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.Sequence.Get(index)
}

func (seq *Sequence[T]) GetRef(index int) *T {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return seq.Sequence.GetRef(index)
}

func (seq *Sequence[T]) Set(index int, value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Set(index, value)
}

func (seq *Sequence[T]) Append(value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Append(value)
}

func (seq *Sequence[T]) AppendMany(values ...T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.AppendMany(values...)
}

func (seq *Sequence[T]) Insert(index int, value T) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Insert(index, value)
}

func (seq *Sequence[T]) Remove(index int) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Remove(index)
}

func (seq *Sequence[T]) Clear() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Clear()
}

func (seq *Sequence[T]) Reverse() {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Reverse()
}

func (seq *Sequence[T]) Sort(comparator functions.Comparator[T]) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Sort(comparator)
}

func (seq *Sequence[T]) Join(other sequences.Sequence[T]) {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	seq.Sequence.Join(other)
}

func (seq *Sequence[T]) Clone() sequences.Sequence[T] {
	seq.mutex.Lock()
	defer seq.mutex.Unlock()

	return &Sequence[T]{Sequence: seq.Sequence.Clone()}
}

func (seq *Sequence[T]) Iterator() collections.Iterator[T] {
	return seq.ModifyingIterator()
}

func (seq *Sequence[T]) ModifyingIterator() sequences.Iterator[T] {
	return iterator[T]{seq.Sequence.ModifyingIterator(), seq}
}

func (seq *Sequence[T]) Stream() streams.Stream[T]     { return seq.Sequence.Stream() }
func (seq *Sequence[T]) RefStream() streams.Stream[*T] { return seq.Sequence.RefStream() }
