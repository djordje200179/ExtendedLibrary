package sequences

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Iterator[T any] interface {
	collections.Iterator[T]
	GetRef() *T
	Set(value T)

	InsertBefore(value T)
	InsertAfter(value T)

	Remove()
}

type Sequence[T any] interface {
	Size() int

	Get(index int) T
	GetRef(index int) *T
	Set(index int, value T)

	Append(value T)
	AppendMany(values ...T)
	Insert(index int, value T)
	Remove(index int)

	Clear()
	Reverse()
	Sort(comparator functions.Comparator[T])
	Join(other Sequence[T])
	misc.Cloner[Sequence[T]]

	collections.Iterable[T]
	ModifyingIterator() Iterator[T]
	streams.Streamer[T]
	RefStream() streams.Stream[*T]
}
