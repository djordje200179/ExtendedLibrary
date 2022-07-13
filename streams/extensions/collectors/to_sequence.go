package collectors

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
	"github.com/djordje200179/extendedlibrary/streams"
)

type sequenceCollector[T any] struct {
	seq sequences.Sequence[T]
}

func ToSequence[T any](empty sequences.Sequence[T]) streams.Collector[T, sequences.Sequence[T]] {
	return sequenceCollector[T]{empty}
}

func (collector sequenceCollector[T]) Supply(value T) { collector.seq.Append(value) }

func (collector sequenceCollector[T]) Finish() sequences.Sequence[T] { return collector.seq }

func ToArray[T any]() streams.Collector[T, sequences.Sequence[T]] {
	return ToSequence[T](array.New[T]())
}

func ToLinkedList[T any]() streams.Collector[T, sequences.Sequence[T]] {
	return ToSequence[T](linkedlist.New[T]())
}
