package sequences

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Sequence[T any] interface {
	datastructures.Collection[int, T]

	Append(values ...T)
	Insert(index int, value T)
	Remove(index int) T

	Empty()

	Sort(comparator comparison.Comparator[T])
	Join(other Sequence[T])

	Iterator() Iterator[T]
	streams.Streamer[T]
}

func CreateStream[T any](sequence Sequence[T]) streams.Stream[T] {
	it := sequence.Iterator()

	return streams.Supply(func() T {
		defer it.Move()
		return it.Get()
	})
}
