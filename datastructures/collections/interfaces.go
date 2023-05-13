package collections

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Iterator[T any] interface {
	iterable.Iterator[T]
	GetRef() *T
	Set(value T)

	InsertBefore(value T)
	InsertAfter(value T)

	Remove()

	Index() int
}

type Collection[T any] interface {
	Size() int

	Get(index int) T
	GetRef(index int) *T
	Set(index int, value T)

	Append(value T)
	Insert(index int, value T)
	Remove(index int)

	Clear()
	Reverse()
	Swap(index1, index2 int)
	Sort(comparator comparison.Comparator[T])
	Join(other Collection[T])

	misc.Cloner[Collection[T]]

	iterable.Iterable[T]
	ModifyingIterator() Iterator[T]
	streams.Streamer[T]
	RefsStream() streams.Stream[*T]

	FindIndex(predictor predication.Predictor[T]) (int, bool)
	//FindValue(value T) (int, bool)
}
