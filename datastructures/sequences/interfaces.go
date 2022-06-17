package sequences

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Iterator[T any] interface {
	datastructures.Iterator[T]

	Set(value T)

	InsertBefore(value T)
	InsertAfter(value T)

	Remove()
}

type Sequence[T any] interface {
	datastructures.Collection[int, T]
	Append(values ...T)
	Insert(index int, value T)
	Remove(index int)

	Clear()
	Sort(comparator functions.Comparator[T])
	Join(other Sequence[T])
	misc.Cloner[Sequence[T]]

	datastructures.Iterable[T]
	ModifyingIterator() Iterator[T]
	streams.Streamer[T]
}
