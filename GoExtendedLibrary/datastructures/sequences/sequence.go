package sequences

import (
	"github.com/djordje200179/GoExtendedLibrary/misc/functions"
	"github.com/djordje200179/GoExtendedLibrary/streams"
)

type Sequence[T any] interface {
	Size() int

	Get(index int) T
	Set(index int, value T)

	Append(values ...T)
	Insert(index int, value T)
	Remove(index int) T

	Sort(less functions.Less[T])
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
